package cat

import (
	"github.com/lib/pq"
)

type ListCat struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Race        string         `json:"race"`
	Sex         string         `json:"sex"`
	AgeInMonth  int            `json:"ageInMonth"`
	Description string         `json:"description"`
	ImageUrls   pq.StringArray `json:"imageUrls"`
	HasMatched  bool           `json:"hasMatched"`
	OwnerId     int64          `json:"ownerId"`
	CreatedAt   string         `json:"createdAt"`
}
