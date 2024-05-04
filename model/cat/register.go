package cat

type RegisterRequest struct {
	Name        string   `json:"name" binding:"required,min=1,max=30"`
	Race        string   `json:"race" binding:"required"`
	Sex         string   `json:"sex" binding:"required"`
	AgeInMonth  int      `json:"ageInMonth"`
	Description string   `json:"description" binding:"required,min=1,max=200"`
	ImageUrls   []string `json:"imageUrls" binding:"required"`
}
