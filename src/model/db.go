package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	database, err := sql.Open("postgres", "postgres://test:test@localhost:15432/test")
	if err != nil {
		log.Fatal("Cannot find database. Received error: " + err.Error())
	} else {
		db = database
	}
}
