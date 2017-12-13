package entities

import (
	"errors"
	"time"
)

// Meeting one meeting entity
type Meeting struct {
	Title         string    `xorm:"notnull pk 'title'"`
	Participators []string  `xorm:"participators"`
	StartTime     time.Time `xorm:"starttime"`
	EndTime       time.Time `xorm:"endtime"`
	Sponsor       string    `xorm:"sponsor"`
}

// ---------------------
// meeting actions
// ---------------------

// NewMeeting create a new meeting and add to database
func NewMeeting(title, start, end, spon string, parts []string) (Meeting, error) {
	newM := Meeting{}
	if !(validateTitle(title)) {
		return newM, errors.New("title has existed")
	}
	if !validateParticipators(parts) {
		return newM, errors.New("not registered participators")
	}
	startTime, ok1 := getTime(start)
	endTime, ok2 := getTime(end)
	if (!ok1) || (!ok2) {
		return newM, errors.New("time format error: require format \"YYYY-MM-DD\"")
	}
	if !(validateTime(startTime, endTime)) {
		return newM, errors.New("start time must before end time")
	}
	if !(validateNoConflicts(parts, startTime, endTime)) {
		return newM, errors.New("time conflict: someone has meetings during this time interval")
	}

	newM.Title = title
	newM.Participators = parts
	newM.StartTime = startTime
	newM.EndTime = endTime
	newM.Sponsor = spon
	_, err := agendaDB.Insert(&newM)
	if err != nil {
		return newM, err
	}

	var user User
	agendaDB.Where("username=?", spon).Get(&user)
	user.Meetings = append(user.Meetings, title)
	agendaDB.Where("username=?", spon).Update(&user)
	for _, part := range parts {
		var u User
		agendaDB.Where("username=?", part).Get(&u)
		u.Meetings = append(u.Meetings, title)
		agendaDB.Where("username=?", part).Update(&u)
	}

	return newM, nil
}

// QuitMeeting quit the meeting the user participated in
func QuitMeeting(user, title string) error {
	var u User
	var m Meeting

	agendaDB.Where("username=?", user).Get(&u)
	got, err := agendaDB.Where("title=?", title).Get(&m)
	if !got {
		return errors.New("title error: no such meeting title")
	}
	if err != nil {
		return err
	}

	isDel := false
	for index, t := range u.Meetings {
		if t == title {
			u.Meetings = append(u.Meetings[:index], u.Meetings[index+1:]...)
			isDel = true
			break
		}
	}
	if !isDel {
		return errors.New("user error: the user has not participated in the meeting")
	}
	for index, t := range m.Participators {
		if t == user {
			m.Participators = append(m.Participators[:index], m.Participators[index+1:]...)
			break
		}
	}

	agendaDB.Where("username=?", user).Update(&u)
	if len(m.Participators) == 0 {
		agendaDB.Where("title=?", title).Delete(&m)
	} else {
		agendaDB.Where("title=?", title).Update(&m)
	}

	return nil
}

// ClearMeeting clear all the meeting one sponsored
func ClearMeeting(user string) {
	var u User
	var newMeetings []string

	agendaDB.Where("username=?", user).Get(&u)
	for _, t := range u.Meetings {
		var m Meeting
		agendaDB.Where("title=?", t).Get(&m)
		if m.Sponsor == user {
			agendaDB.Where("title=?", t).Delete(&m)
		} else {
			newMeetings = append(newMeetings, t)
		}
	}
	u.Meetings = newMeetings
	agendaDB.Where("username=?", user).Update(&u)
}

// QueryMeeting query one's meetings between a specified time interval
func QueryMeeting(user, start, end string) ([]Meeting, error) {
	var rst []Meeting
	startTime, ok1 := getTime(start)
	endTime, ok2 := getTime(end)

	if (!ok1) || (!ok2) {
		return []Meeting{}, errors.New("time format error: require format \"YYYY-MM-DD\"")
	}
	if !(validateTime(startTime, endTime)) {
		return []Meeting{}, errors.New("start time must before end time")
	}

	var u User
	agendaDB.Where("username=?", user).Get(&u)
	for _, t := range u.Meetings {
		var m Meeting
		agendaDB.Where("title=?", t).Get(&m)
		if !(m.StartTime.After(endTime) || m.EndTime.Before(startTime)) {
			rst = append(rst, m)
		}
	}

	return rst, nil
}

// -------------------------
// Validation Functions
// -------------------------

func validateTitle(title string) bool {
	var m Meeting
	got, err := agendaDB.Where("title=?", title).Get(&m)
	if got == false || err != nil {
		return true
	}
	return false
}

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

func validateTime(start, end time.Time) bool {
	if start.After(end) || start.Equal(end) {
		return false
	}
	return true
}

func validateNoConflicts(parts []string, start, end time.Time) bool {
	for _, part := range parts {
		var u User
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

// ------------------
// helpful function
// ------------------
func getTime(t string) (time.Time, bool) {
	tmpTime, err := time.Parse("2006-01-02", t)
	if err != nil {
		return time.Time{}, false
	}
	return tmpTime, true
}
