package entities

import (
	"fmt"
	"log"
)

type User struct {
	UserName string   `xorm:"username" pk`
	Password string   `xorm:"password"`
	Email    string   `xorm:"email"`
	Tel      string   `xorm:"telephone"`
	Meetings []string `xorm:"meetings"`
}

var currentUser User
var Users map[int64]User

func NewUser(username string, password string) User {
	var user User
	user.UserName = username
	user.Password = password
	user.Email = ""
	user.Tel = ""
	return user
}

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

func Register(username string, password string) bool {
	if validUsername(username) {
		user := NewUser(username, password)
		_, err := agendaDB.Insert(user)
		if err != nil {
			log.Fatal("something wrong occured when register new user")
		}
		fmt.Println("register successfully")
		return true
	} else {
		fmt.Println("invalid register information")
		return false
	}
}

func Login(username string, password string) bool {
	var user User
	agendaDB.Where("username=?", username).Get(&user)
	if user.UserName == username && user.Password == password {
		var loginOne LoginInfo
		loginOne.UserName = user.UserName
		agendaDB.Insert(loginOne)
		return true
	} else {
		fmt.Println("wrong username or password")
		return false
	}
}

func logout() bool {
	var user User
	agendaDB.Where("username = ?", currentUser.UserName).Get(&user)
	if user.UserName == currentUser.UserName {
		fmt.Println("log out successfully")
		agendaDB.Where("username = ?", currentUser.UserName).Delete(user)
		return true
	} else {
		fmt.Println("you haven't logined yet!")
		return false
	}
}

func setEmail(email string) {
	user := NewUser(currentUser.UserName, currentUser.Password)
	user.Email = email
	_, err := agendaDB.Id(currentUser.UserName).Update(user)
	if err != nil {
		log.Fatal("something wrong occur when update email")
	}
}

func setTelephone(tel string) {
	user := NewUser(currentUser.UserName, currentUser.Password)
	user.Tel = tel
	_, err := agendaDB.Id(currentUser.UserName).Update(user)
	if err != nil {
		log.Fatal("something wrong occur when update telephone")
	}
}

func GetUserKey(username string) LoginInfo {
	var user LoginInfo
	agendaDB.Where("UserName=?", username).Get(&user)
	return user
}

func GetUserById(id int64) User {
	var logined LoginInfo
	_, err := agendaDB.Where("Key=?", id).Get(&logined)
	if err != nil {
		log.Fatal("something wrong occured in getUserbyId")
	}
	username := logined.UserName
	var user User
	_, err = agendaDB.Where("UserName=?", username).Get(&user)
	if err != nil {
		log.Fatal("something wrong occured in getUserbyId")
	}
	return user
}

func GetAllUsers() []User {
	users := make([]User, 0)
	err := agendaDB.Find(&users)
	if err != nil {
		log.Fatal("something wrong occured when get all users")
	}
	return users
}
