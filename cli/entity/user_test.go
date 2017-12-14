package entity

import (
	"testing"
)

// TestMeeting test user's function
func TestUser(t *testing.T) {
	InitAllUsers()
	InitAllMeetings()

	t.Log("[usertest]: registering new user Jack")
	if !Register("Jack", "123456") {
		t.Fatal("register failure")
	}
	t.Log("register success")

	t.Log("[usertest]: Jack login")
	if !Login("Jack", "123456") {
		t.Fatal("login failure")
	}
	t.Log("login success")

	InitAllUsers()
	u := GetCurrentUser()

	t.Log("[usertest]: setting Jack's email")
	if !u.SetEmail("Jack@ubuntu.com") {
		t.Fatal("set email failure")
	}
	t.Log("set email success")

	t.Log("[usertest]: setting Jack's telephone")
	if !u.SetTelephone("12345678902") {
		t.Fatal("set telephone failure")
	}
	t.Log("set telephone success")

	t.Log("[usertest]: cancelling user Jack")
	if !u.CancelAccount() {
		t.Fatal("cancel user failure")
	}
	t.Log("cancel user success")
}
