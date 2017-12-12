package handler

import (
	"log"
	"net/http"
	"strconv"

	"microservice-agenda/service/entities"

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
