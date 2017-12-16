package entity

import (
	"os"
	"testing"
)

// TestMeeting test user's function
func TestUser(t *testing.T) {
	os.Mkdir("data", 0755)
	dbFile := "data/cli-agenda.db"
	InitializeDB(dbFile)

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

	cu, _ := GetCurrentUser()

	t.Log("[usertest]: cancelling user Jack")
	if !CancelAccount(cu) {
		t.Fatal("cancel user failure")
	}
	t.Log("cancel user success")
}
