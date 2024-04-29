package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error

	DB, err = sql.Open("postgres", "host=localhost port=5432 user=dionfananie password=yoloyolo dbname=catmatchdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	println("database connected!")
}
