package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("postgres", "postgres:password@localhost:5432/catmatchdb")
	if err != nil {
		log.Fatal(err)
	}
}
