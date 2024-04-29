package auth

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=5,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}
