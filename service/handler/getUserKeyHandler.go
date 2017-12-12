package handler

import (
	"net/http"

	"microservice-agenda/service/entities"

	"github.com/unrolled/render"
)

func getUserKeyHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.JSON(w, http.StatusOK, entities.GetUserKey(req.Form["username"][0]))
	}
}
