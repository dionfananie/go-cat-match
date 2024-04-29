package main

import (
	router "web/go-cat-match/app"
	"web/go-cat-match/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()
	// Setup Auth Middleware

	router.SetupRouter(r)

	r.Run("localhost:8080")
}