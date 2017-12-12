package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/v1/meetings", createMeetingHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/meetings", quitMeetingHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/v1/meetings", clearMeetingHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/v1/meetings", queryMeetingHandler(formatter)).Methods("GET")

	mx.HandleFunc("/v1/allusers", ListAllUsersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/users/getkey/{username,password}", GetUserByIdHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/users/{id}", GetUserKeyHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/newusers/{username,password}", GetUserKeyHandler(formatter)).Methods("POST")
}
