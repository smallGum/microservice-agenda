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

// NewMeeting create a new meeting and add to database
func NewMeeting(title, start, end, spon string, parts []string) (Meeting, error) {
	newM := Meeting{}
	if !(validateTitle(title)) {
		return newM, errors.New("title has existed")
	}
	if !(validateParticipators(parts, title)) {
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
	agendaDB.Id(spon).Get(&user)
	user.Meetings = append(user.Meetings, title)
	agendaDB.Id(spon).Update(&user)
	for _, part := range parts {
		agendaDB.Id(part).Get(&user)
		user.Meetings = append(user.Meetings, title)
		agendaDB.Id(part).Update(&user)
	}

	return newM, nil
}

// -------------------------
// Validation Functions
// -------------------------

func validateTitle(title string) bool {
	var m Meeting
	got, _ := agendaDB.Id(title).Get(&m)
	return !got
}

func validateParticipators(parts []string, title string) bool {
	var user User
	for _, part := range parts {
		got, _ := agendaDB.Id(title).Get(&user)
		if !got {
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
	var m Meeting
	var u User
	for _, part := range parts {
		agendaDB.Id(part).Get(&u)
		for _, ms := range u.Meetings {
			agendaDB.Id(ms).Get(&m)
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
