package entity

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	UserName string `xorm:"username" pk`
	Password string `xorm:"password"`
	Email    string `xorm:"email"`
	Tel      string `xorm:"telephone"`
}

type UserOp struct {
	OpTime    time.Time `xorm:"notnull pk 'operation_time'"`
	UserName  string    `xorm:"username"`
	Operation string    `xorm:"operation"`
}

type LoginInfo struct {
	CurrentUser string    `xorm:"notnull pk 'current_user'"`
	LoginTime   time.Time `xorm:"login_time"`
}

// -----------------
// User commands
// -----------------

func NewUser(username string, password string) User {
	var user User
	user.UserName = username
	user.Password = password
	user.Email = ""
	user.Tel = ""
	return user
}

func Register(username string, password string) bool {
	if validUsername(username) {
		user := NewUser(username, password)
		_, err := agendaDB.Insert(&user)
		if err != nil {
			log.Fatal("something wrong occured when register new user")
		}

		recordOperation(username, "register")
		fmt.Println("register successfully")
		return true
	} else {
		fmt.Println("invalid register information")
		return false
	}
}

func Login(username string, password string) bool {
	cu := new(LoginInfo)
	total, _ := agendaDB.Count(cu)
	if total != 0 {
		log.Fatal("you have already logged in, to switch to another account," +
			"you must log out first")
		return false
	}

	var user User
	agendaDB.Where("username=?", username).Get(&user)
	if user.UserName == username && user.Password == password {
		loginOne := LoginInfo{
			CurrentUser: user.UserName,
			LoginTime:   time.Now(),
		}
		agendaDB.Insert(&loginOne)

		recordOperation(username, "login")
		return true
	} else {
		fmt.Println("wrong username or password")
		return false
	}
}

func SetEmail(email, cuName string) {
	var user User
	agendaDB.Where("username = ?", cuName).Get(&user)
	user.Email = email
	agendaDB.Where("username = ?", cuName).Update(&user)

	recordOperation(user.UserName, "set email")
}

func SetTelephone(tel, cuName string) {
	var user User
	agendaDB.Where("username = ?", cuName).Get(&user)
	user.Tel = tel
	agendaDB.Where("username = ?", cuName).Update(&user)

	recordOperation(user.UserName, "set telephone")
}

func Logout(cuName string) {
	var cu LoginInfo
	agendaDB.Where("current_user = ?", cuName).Delete(&cu)

	recordOperation(cuName, "logout")
}

func LookupAllUser(cuName string) {
	users := getAllUsers()
	fmt.Println("there are", len(users), " users:")
	fmt.Println("--------------------------")
	for _, user := range users {
		fmt.Println("user:" + user.UserName)
		fmt.Println("email:" + user.Email)
		fmt.Println("tel:" + user.Tel)
		fmt.Println("--------------------------")
	}

	recordOperation(cuName, "check all users")
}

func CancelAccount(cuName string) bool {
	meetings := getOnesMeetings(cuName)
	for _, meeting := range meetings {
		if done := QuitMeeting(meeting, cuName); !done {
			return false
		}
	}
	ClearAllMeetings(cuName)

	var user User
	var login LoginInfo
	agendaDB.Where("username = ?", cuName).Delete(&user)
	agendaDB.Where("current_user = ?", cuName).Delete(&login)

	recordOperation(cuName, "cancel this account")

	return true
}

// --------------------
// helpful functions
// --------------------

func validUsername(username string) bool {
	user := new(User)
	has, err := agendaDB.Where("username=?", username).Get(user)
	if err != nil {
		log.Fatal("something wrong occured when validate the register name")
	}
	if has {
		return false
	} else {
		return true
	}
}

func GetCurrentUser() (string, bool) {
	currentUser := make([]LoginInfo, 0)
	err := agendaDB.Find(&currentUser)
	if err != nil {
		log.Fatal("something wrong occured when get current user info")
	}
	if len(currentUser) == 0 {
		fmt.Println("please login first!")
		return "", false
	}

	return currentUser[0].CurrentUser, true
}

func getAllUsers() []User {
	users := make([]User, 0)
	err := agendaDB.Find(&users)
	if err != nil {
		log.Fatal("something wrong occured when get all users")
	}
	return users
}

func getOnesMeetings(name string) []string {
	var meetings Participation
	got, err := agendaDB.Where("username = ?", name).Get(&meetings)
	if !got || err != nil {
		return []string{}
	}
	return meetings.Meetings
}

func recordOperation(username, op string) {
	operation := UserOp{
		OpTime:    time.Now(),
		UserName:  username,
		Operation: op,
	}
	agendaDB.Insert(&operation)
}
