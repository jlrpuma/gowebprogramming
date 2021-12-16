package data

import (
	"database/sql"
	"log"
)

// exported member to get access to the dabatase
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode= disable")
	if err != nil {
		log.Fatal(err)
	}
	//TODO: why do this fuction have a return statement ?
	return
}
