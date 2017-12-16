package errors

import (
	"log"
	"os"
)

// ErrorMsg record the error messsage and print the error message
func ErrorMsg(usr, err string) {
	file, _ := os.OpenFile("data/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0)
	logger := log.New(file, usr+": ", log.Lshortfile)
	logger.Println(err)
	log.Println(err)

	file.Close()
}
