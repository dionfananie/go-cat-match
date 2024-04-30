package controller

import (
	"fmt"
	"net/http"
	"web/go-cat-match/database"
	"web/go-cat-match/model/auth"
	"web/go-cat-match/utils/jwt"
	"web/go-cat-match/utils/password"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func Login(c *gin.Context) {
	var req auth.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Name string
	var Email string
	var Pass string
	var UserId uint64

	err := database.DB.QueryRow("SELECT id, name, password, email from users WHERE email = $1", req.Email).Scan(&UserId, &Name, &Pass, &Email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return

	}

	errs := password.Verify(Pass, req.Password)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request doesnt pass validation"})
		return

	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: UserId,
	})

	c.JSON(http.StatusOK, gin.H{"message": "User logged successfully", "data": gin.H{
		"email":       Email,
		"name":        Name,
		"accessToken": accessToken,
	}})
}

func Register(c *gin.Context) {
	var req auth.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword := password.Hash(req.Password)

	var emailExists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", req.Email).Scan(&emailExists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if emailExists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	var UserId uint64
	err = database.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", req.Name, req.Email, hashedPassword).Scan(&UserId)
	if err != nil {

		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code)
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: UserId,
	})

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "data": gin.H{
		"email":       req.Email,
		"name":        req.Name,
		"accessToken": accessToken,
	}})
}
