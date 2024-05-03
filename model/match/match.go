package match

type MatchCatRequest struct {
	MatchCatId int    `json:"matchCatId" binding:"required,numeric"`
	UserCatId  int    `json:"userCatId" binding:"required,numeric"`
	Message    string `json:"message" binding:"required,min=5,max=120"`
}

type status string

const (
	STATUS_PENDING  status = "pending"
	STATUS_APPROVED status = "approved"
	STATUS_REJECTED status = "rejected"
)

type Match struct {
	Id           int64
	IssuedUserId int64
	Status       status
	UserCatId    int64
	MatchCatId   int64
	Message      string
	CreatedAt    string
}
