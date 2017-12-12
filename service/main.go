package main

import (
	"os"

	"microservice-agenda/service/entities"
	"microservice-agenda/service/service"

	flag "github.com/spf13/pflag"
)

// PORT port for server to listen
const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	dbFile := flag.StringP("database", "d", "", "sqlite3 database file for loading data")
	flag.Parse()

	if len(*pPort) != 0 {
		port = *pPort
	}
	if len(*dbFile) == 0 {
		os.Mkdir("data", 0755)
		*dbFile = "data/agenda.db"
	}

	entities.InitializeDB(*dbFile)
	server := service.NewServer()
	server.Run(":" + port)
}
