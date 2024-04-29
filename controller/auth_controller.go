package controller

import (
	"net/http"
	"web/go-cat-match/database"
	"web/go-cat-match/utils/jwt"
	"web/go-cat-match/utils/password"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=5,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO : Complete this function
	var Name string
	var Email string
	var Pass string
	var UserId uint64
	println("user checking: ", req.Email)

	// query := "SELECT name, password as pass, email from users WHERE email = ?"
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
	println("UserId", Name)
	println("UserId", Email)
	println("UserId", UserId)

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
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword := password.Hash(req.Password)

	var UserId uint64
	err := database.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", req.Name, req.Email, hashedPassword).Scan(&UserId)

	//TODO: HANDLE EMAIL UNIQUE VALIDATION (409)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: UserId,
	})

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "data": gin.H{
		"email":       req.Email,
		"name":        req.Name,
		"accessToken": accessToken,
	}})
}
