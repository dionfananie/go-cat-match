package main

import (
	router "web/go-cat-match/app"
	"web/go-cat-match/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()

	router.SetupRouter(r)

	r.Run(":8080")
}
