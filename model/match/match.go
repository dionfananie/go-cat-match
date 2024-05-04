package match

import "github.com/lib/pq"

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

type issuedBy struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type catDetail struct {
	Id          uint64         `json:"id"`
	Name        string         `json:"name"`
	Race        string         `json:"race"`
	Sex         string         `json:"sex"`
	Description string         `json:"description"`
	AgeInMonth  int            `json:"ageInMonth"`
	ImageUrls   pq.StringArray `json:"imageUrls"`
}

type MatchFullInfo struct {
	Id             int64     `json:"id"`
	IssuedBy       issuedBy  `json:"issuedBy"`
	MatchCatDetail catDetail `json:"matchCatDetail"`
	UserCatDetail  catDetail `json:"userCatDetail"`
	Message        string    `json:"message"`
	CreatedAt      string    `json:"createdAt"`
}
