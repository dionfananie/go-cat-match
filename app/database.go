
package app

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "os"
  )
  

const (
    host     = "localhost"
    port     = 5432
    user_name = "dionfananie"
    password = "yoloyolo"
    dbname   = "cat_match"
)

func NewDb() {
  
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user_name, password, dbname)
    println(psqlInfo)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!");
   
}