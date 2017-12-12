package handler

import (
	"net/http"

	"microservice-agenda/service/entities"

	"github.com/unrolled/render"
)

func listAllUsers(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.JSON(w, http.StatusOK, entities.GetAllUsers())
	}
}
