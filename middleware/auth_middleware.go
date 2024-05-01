package middleware

import (
	"net/http"
	"strings"
	"web/go-cat-match/utils/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty Header Authorization"})
		return
	}

	tokenParts := strings.Split(tokenString, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": tokenString})
		return
	}

	token := tokenParts[1]

	payload, err := jwt.Verify(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("userId", payload.UserId)

	c.Next()
}
