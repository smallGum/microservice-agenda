package handler

import (
	"net/http"

	"microservice-agenda/service/entities"

	"github.com/unrolled/render"
)

func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		username := req.Form["username"][0]
		password := req.Form["password"][0]
		if entities.Login() {
			formatter(w,http.StatusOK,"log in successfully")
		}else{
	 		formatter(w,http.StatusOK,"log in failed")
		}
	}
}
