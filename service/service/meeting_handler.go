package service

import (
	"net/http"

	"github.com/unrolled/render"

	"microservice-agenda/service/entities"
)

func createMeetingHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		// deal with bad request
		if len(req.Form["key"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"key required!"})
			return
		}
		if len(req.Form["title"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"title required!"})
			return
		}
		if len(req.Form["participators"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"participators required!"})
			return
		}
		if len(req.Form["startTime"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"startTime required!"})
			return
		}
		if len(req.Form["endTime"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"endTime required!"})
			return
		}

		// check if key is valid
		currentUser, err := entities.CheckKey(req.Form["key"][0])
		if err != nil {
			formatter.JSON(w, http.StatusUnauthorized, struct{ ErrorIndo string }{err.Error()})
			return
		}

		// create a new meeting
		newMeeting, err := entities.NewMeeting(req.Form["title"][0], req.Form["startTime"][0], req.Form["endTime"][0], currentUser, req.Form["participators"])
		if err != nil {
			formatter.JSON(w, http.StatusServiceUnavailable, struct{ ErrorIndo string }{err.Error()})
			return
		}

		formatter.JSON(w, http.StatusCreated, newMeeting)
	}
}

func quitMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		// deal with bad request
		if len(req.Form["key"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"key required!"})
			return
		}
		if len(req.Form["title"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"title required!"})
			return
		}

		// check if key is valid
		currentUser, err := entities.CheckKey(req.Form["key"][0])
		if err != nil {
			formatter.JSON(w, http.StatusUnauthorized, struct{ ErrorIndo string }{err.Error()})
			return
		}

		// query a meeting
		err = entities.QuitMeeting(currentUser, req.Form["title"][0])
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, struct{ ErrorIndo string }{err.Error()})
			return
		}

		formatter.JSON(w, http.StatusOK, struct{ ErrorIndo string }{"success!"})
	}
}

func clearMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		// deal with bad request
		if len(req.Form["key"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"key required!"})
			return
		}

		// check if key is valid
		currentUser, err := entities.CheckKey(req.Form["key"][0])
		if err != nil {
			formatter.JSON(w, http.StatusUnauthorized, struct{ ErrorIndo string }{err.Error()})
			return
		}

		// query a meeting
		entities.ClearMeeting(currentUser)
		formatter.JSON(w, http.StatusOK, struct{ ErrorIndo string }{"success!"})
	}
}

func queryMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		// deal with bad request
		if len(req.Form["key"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"key required!"})
			return
		}
		if len(req.Form["startTime"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"startTime required!"})
			return
		}
		if len(req.Form["endTime"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"endTime required!"})
			return
		}

		// check if key is valid
		currentUser, err := entities.CheckKey(req.Form["key"][0])
		if err != nil {
			formatter.JSON(w, http.StatusUnauthorized, struct{ ErrorIndo string }{err.Error()})
			return
		}

		newMeeting, err := entities.QueryMeeting(currentUser, req.Form["startTime"][0], req.Form["endTime"][0])
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, struct{ ErrorIndo string }{err.Error()})
			return
		}

		formatter.JSON(w, http.StatusOK, newMeeting)
	}
}
