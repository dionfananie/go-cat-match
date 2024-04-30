package database

import (
	"database/sql"
	"fmt"
	"log"
	"web/go-cat-match/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	var dbCredentials = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PARAMS)

	DB, err = sql.Open("postgres", dbCredentials)
	if err != nil {
		log.Fatal(err)
	}
	println("database connected!")
}
