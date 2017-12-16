package entity

import (
	"os"
	"testing"
)

// TestMeeting test meetings' function
func TestMeeting(t *testing.T) {
	os.Mkdir("data", 0755)
	InitializeDB("data/agenda.db")

	m1 := &Meeting{
		Title:         "test_meeting_title1",
		Participators: []string{"Alice", "Bob"},
		Sponsor:       "Jack",
	}

	m2 := &Meeting{
		Title:         "test_meeting_title2",
		Participators: []string{"Jack", "Bob"},
		Sponsor:       "Alice",
	}

	u1 := &User{UserName: "Jack"}
	u2 := &User{UserName: "Bob"}
	u3 := &User{UserName: "Alice"}
	agendaDB.Insert(u1)
	agendaDB.Insert(u2)
	agendaDB.Insert(u3)

	t.Log("[meetingtest] creating meeting1")
	NewMeeting(m1.Title, "2017-01-01", "2017-01-02", m1.Sponsor, m1.Participators)
	t.Log("[meetingtest] creating meeting2")
	NewMeeting(m2.Title, "2017-01-03", "2017-01-05", m2.Sponsor, m2.Participators)
	t.Log("[meetingtest] querying meeting")
	m := queryMeeting("Alice", "2017-01-01", "2017-01-04")
	if len(m) != 2 {
		t.Fatal("meeting creation failure")
	}
	t.Log("meeting creation success")

	t.Log("[meetingtest] quiting meeting")
	QuitMeeting("test_meeting_title1", "Alice")
	m = queryMeeting("Alice", "2017-01-02", "2017-01-04")
	if len(m) != 1 {
		t.Fatal("meeting quit failure")
	}
	t.Log("meeting quit success")

	t.Log("[meetingtest] clearing meeting")
	ClearAllMeetings("Alice")
	t.Log("[meetingtest] querying meeting")
	m = queryMeeting("Jack", "2017-01-02", "2017-01-04")
	if len(m) != 1 {
		t.Fatal("meeting clear failure")
	}
	t.Log("meeting clear success")
}
