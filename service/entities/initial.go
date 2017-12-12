package entities

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var agendaDB *xorm.Engine

// InitializeDB initialize database
func InitializeDB(dbFile string) {
	var err error
	agendaDB, err = xorm.NewEngine("sqlite3", dbFile)
	checkErr(err)
	// TODO: add all tables
	err = agendaDB.Sync2(new(LoginInfo))
	err = agendaDB.Sync2(new(User))

	checkErr(err)
}

// error detection
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
