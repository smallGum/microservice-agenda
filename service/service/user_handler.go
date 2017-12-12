package service

import (
	"log"
	"microservice-agenda/service/entities"
	"net/http"
	"strconv"

	"github.com/unrolled/render"
)

func GetUserByIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
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
		formatter.JSON(w, http.StatusOK, entities.GetUserKey(req.Form["username"][0]))
	}
}

func ListAllUsersHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.JSON(w, http.StatusOK, entities.GetAllUsers())
	}
}

func LoginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		username := req.Form["username"][0]
		password := req.Form["password"][0]
		if entities.Login(username, password) {
			formatter.JSON(w, http.StatusOK, entities.NewUser(username, password))
		}
	}
}
