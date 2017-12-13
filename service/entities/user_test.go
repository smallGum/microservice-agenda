package entities

import (
	"os"
	"testing"
)

// TestMeeting test user's function
func TestUser(t *testing.T) {
	os.Mkdir("data", 0755)
	InitializeDB("data/agenda.db")

	t.Log("[usertest] registering 3 users: Jackie, Darth and Lucy")
	if !Register("Jackie", "123456") {
		t.Fatal("Jackie register failure")
	}
	if !Register("Darth", "123456") {
		t.Fatal("Darth register failure")
	}
	if !Register("Lucy", "123456") {
		t.Fatal("Lucy register failure")
	}
	t.Log("register success")

	t.Log("[usertest] Jackie login")
	if !Login("Jackie", "123456") {
		t.Fatal("Jackie login failure")
	}
	t.Log("Jackie login success")

	t.Log("[usertest] getting Jackie's key")
	firstOne := GetUserKey("Jackie")
	if firstOne.UserName != "Jackie" {
		t.Fatal("get Jackie's key failure")
	}
	t.Log("Jackie's key: ", firstOne.Key)

	t.Log("[usertest] getting logined user Jackie by id")
	u := GetUserById(1)
	if u.UserName != "Jackie" {
		t.Fatal("get user by id failure")
	}
	t.Log("get user by id success")

	t.Log("getting all users")
	us := GetAllUsers()
	if len(us) != 3 {
		t.Fatal("get all user failure")
	}
	t.Log("get all user success")
}
