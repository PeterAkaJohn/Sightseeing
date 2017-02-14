package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	database, err := sql.Open("postgres", "user=sightseeing dbname=sightseeing password=S1gHtSeeing sslmode=disable")
	if err != nil {
		log.Fatal("Cannot find database. Received error: " + err.Error())
	} else {
		db = database
	}
}
