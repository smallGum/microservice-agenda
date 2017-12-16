package entity

import (
	"fmt"
	"time"
)

// ---------------------------------------------------
// data structures definition
// ---------------------------------------------------

// Meeting one meeting entity
type Meeting struct {
	Title         string    `xorm:"notnull pk 'title'"`
	Participators []string  `xorm:"participators"`
	StartTime     time.Time `xorm:"starttime"`
	EndTime       time.Time `xorm:"endtime"`
	Sponsor       string    `xorm:"sponsor"`
}

type Participation struct {
	UserName          string   `xorm:"notnull pk 'username'"`
	Meetings          []string `xorm:"meetings"`
	SponsoredMeetings []string `xorm:"sponsored_meetings"`
}

// -----------------------------------------------------
// Meeting structure methods definition
// -----------------------------------------------------

// NewMeeting create a new meeting and add to database
func NewMeeting(title, start, end, cuName string, parts []string) bool {
	if !(validateTitle(title)) {
		return false
	}
	if !(validateParticipators(parts)) {
		return false
	}
	startTime, ok1 := getTime(start)
	endTime, ok2 := getTime(end)
	if (!ok1) || (!ok2) {
		return false
	}
	if !(validateTime(startTime, endTime)) {
		return false
	}
	if !(validateNoConflicts(parts, startTime, endTime)) {
		return false
	}

	m := Meeting{
		Title:         title,
		Participators: parts,
		StartTime:     startTime,
		EndTime:       endTime,
		Sponsor:       cuName,
	}

	_, err := agendaDB.Insert(&m)
	if err != nil {
		return false
	}

	var participation Participation
	got, _ := agendaDB.Where("username=?", cuName).Get(&participation)
	if !got {
		tmpParticipation := Participation{
			UserName:          cuName,
			Meetings:          []string{},
			SponsoredMeetings: []string{},
		}
		_, err = agendaDB.Insert(&tmpParticipation)
		if err != nil {
			return false
		}
	}
	participation.SponsoredMeetings = append(participation.SponsoredMeetings, title)
	agendaDB.Where("username=?", cuName).Update(&participation)
	for _, part := range parts {
		var p Participation
		got, _ = agendaDB.Where("username=?", part).Get(&p)
		if !got {
			tmpP := Participation{
				UserName:          part,
				Meetings:          []string{},
				SponsoredMeetings: []string{},
			}
			_, err = agendaDB.Insert(&tmpP)
			if err != nil {
				return false
			}
		}
		p.Meetings = append(p.Meetings, title)
		agendaDB.Where("username=?", part).Update(&p)
	}

	recordOperation(cuName, "create meeting "+title)

	return true
}

func QuitMeeting(title, cuName string) bool {
	var p Participation
	agendaDB.Where("username = ?", cuName).Get(&p)
	for index, meeting := range p.Meetings {
		if meeting == title {
			p.Meetings = append(p.Meetings[:index], p.Meetings[index+1:]...)
			agendaDB.Where("username = ?", cuName).Cols("meetings").Update(&p)

			var m Meeting
			agendaDB.Where("title = ?", title).Get(&m)
			for i, part := range m.Participators {
				if part == cuName {
					m.Participators = append(m.Participators[:i], m.Participators[i+1:]...)
					break
				}
			}
			if len(m.Participators) == 0 {
				var sponsors []Participation
				agendaDB.Find(&sponsors)
				for _, spon := range sponsors {
					isFind := false
					for j, meet := range spon.SponsoredMeetings {
						if meet == title {
							spon.SponsoredMeetings = append(spon.SponsoredMeetings[:j], spon.SponsoredMeetings[j+1:]...)
							agendaDB.Where("username = ?", spon.UserName).Update(&spon)
							isFind = true
							break
						}
					}
					if isFind {
						break
					}
				}
				agendaDB.Where("title = ?", title).Delete(&m)
			} else {
				agendaDB.Where("title = ?", title).Update(&m)
			}

			recordOperation(cuName, "quit meeting "+title)

			return true
		}
	}

	return false
}

func CancelMeeting(title, cuName string) bool {
	var p Participation
	agendaDB.Where("username = ?", cuName).Get(&p)
	for index, meeting := range p.SponsoredMeetings {
		if meeting == title {
			p.SponsoredMeetings = append(p.SponsoredMeetings[:index], p.SponsoredMeetings[index+1:]...)
			agendaDB.Where("username = ?", cuName).Update(&p)
			var m Meeting
			agendaDB.Where("title = ?", title).Get(&m)
			for _, usr := range m.Participators {
				QuitMeeting(title, usr)
			}

			recordOperation(cuName, "cancel meeting "+title)

			return true
		}
	}

	return false
}

func ClearAllMeetings(cuName string) {
	var p Participation
	agendaDB.Where("username = ?", cuName).Get(&p)
	for _, meeting := range p.SponsoredMeetings {
		CancelMeeting(meeting, cuName)
	}

	recordOperation(cuName, "clear all meetings sponsored by "+cuName)
}

// GetMeetings show meetings between time interval [start, end]
func GetMeetings(cuName, start, end string) {
	ms := queryMeeting(cuName, start, end)

	fmt.Println(cuName + "'s meetings between " + start + " and " + end + ": ")
	if len(ms) == 0 {
		fmt.Println("none.")
		return
	}

	for _, v := range ms {
		fmt.Println()
		fmt.Println("-------------------------------")

		fmt.Println("title: " + v.Title)
		fmt.Printf("participators: %v\n", v.Participators)
		fmt.Println("start time: " + v.StartTime.Format("2006-01-02"))
		fmt.Println("end time: " + v.EndTime.Format("2006-01-02"))
		fmt.Println("sponsor: " + v.Sponsor)

		fmt.Println("-------------------------------")
		fmt.Println()
	}

	recordOperation(cuName, "check all meetings between "+start+" and "+end)
}

// check if title has existed
func validateTitle(title string) bool {
	var m Meeting
	got, err := agendaDB.Where("title=?", title).Get(&m)
	if got == false || err != nil {
		return true
	}
	return false
}

// check if all the participators have registered
func validateParticipators(parts []string) bool {
	for _, part := range parts {
		var user User
		got, err := agendaDB.Where("username=?", part).Get(&user)
		if got != true || err != nil {
			return false
		}
	}

	return true
}

// check if start time is less than end time
func validateTime(start, end time.Time) bool {
	if start.After(end) || start.Equal(end) {
		return false
	}
	return true
}

// check if there are confilts
func validateNoConflicts(parts []string, start, end time.Time) bool {
	for _, part := range parts {
		var u Participation
		agendaDB.Where("username=?", part).Get(&u)
		for _, ms := range u.Meetings {
			var m Meeting
			agendaDB.Where("title=?", ms).Get(&m)
			if !(end.Before(m.StartTime) || end.Equal(m.StartTime) ||
				start.After(m.EndTime) || start.Equal(m.EndTime)) {
				return false
			}
		}
	}
	return true
}

// -----------------------------------------------------
// helpful function
// -----------------------------------------------------

// convert string to time.Time
func getTime(t string) (time.Time, bool) {
	tmpTime, err := time.Parse("2006-01-02", t)
	if err != nil {
		return time.Time{}, false
	}

	return tmpTime, true
}

// ------------------------------------------------------
// query meetings methods
// ------------------------------------------------------

// QueryMeeting query one's meetings between a specified time interval
func queryMeeting(user, start, end string) []Meeting {
	var rst []Meeting
	startTime, ok1 := getTime(start)
	endTime, ok2 := getTime(end)

	if (!ok1) || (!ok2) {
		fmt.Println("time format error: require format \"YYYY-MM-DD\"")
		return []Meeting{}
	}
	if !(validateTime(startTime, endTime)) {
		fmt.Println("start time must before end time")
		return []Meeting{}
	}

	var u Participation
	agendaDB.Where("username=?", user).Get(&u)
	for _, t := range u.Meetings {
		var m Meeting
		agendaDB.Where("title=?", t).Get(&m)
		if !(m.StartTime.After(endTime) || m.EndTime.Before(startTime)) {
			rst = append(rst, m)
		}
	}
	for _, t := range u.SponsoredMeetings {
		var m Meeting
		agendaDB.Where("title=?", t).Get(&m)
		if !(m.StartTime.After(endTime) || m.EndTime.Before(startTime)) {
			rst = append(rst, m)
		}
	}

	return rst
}
