package service

import (
	"fmt"
	"log"
	"microservice-agenda/service/entities"
	"net/http"
	"strconv"

	"github.com/unrolled/render"
)

func GetUserByIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Write([]byte("this is get user by id handler"))
		if len(req.Form["id"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		} else {
			temp, err := strconv.ParseInt(req.Form["id"][0], 10, 64)
			if err != nil {
				log.Fatal("something wrong occured in getUserById")
			}
			formatter.JSON(w, http.StatusOK, entities.GetUserById(temp))
		}
	}
}

func GetUserKeyHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Write([]byte("this is get user key handler"))
		formatter.JSON(w, http.StatusOK, entities.GetUserKey(req.Form["username"][0]))
	}
}

func ListAllUsersHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Write([]byte("this is list all user handler"))
		formatter.JSON(w, http.StatusOK, entities.GetAllUsers())
	}
}

func LoginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Write([]byte("this is log in handler"))
		username := req.Form["username"][0]
		password := req.Form["password"][0]
		w.Write([]byte("get " + username + password))
		if entities.Login(username, password) {
			fmt.Println("logined")
			formatter.JSON(w, http.StatusOK, entities.NewUser(username, password))
		}
	}
}

func registerHndler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Write([]byte("this is register handler"))
		username := req.Form["username"][0]
		password := req.Form["password"][0]
		w.Write([]byte("get " + username + password))
		user := entities.NewUser(username, password)
		if entities.Register(username, password) {
			fmt.Println("register success")
			formatter.JSON(w, http.StatusOK, user)
		} else {
			w.Write([]byte("register failed"))
			fmt.Println("register failed")
		}
	}
}

//--------------------------------------------------------------
func testHndler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello world"))
	}
}
