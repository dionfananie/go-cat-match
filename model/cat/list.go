package cat

import (
	"github.com/lib/pq"
)

type ListCat struct {
	Name        string
	Race        string
	Sex         string
	AgeInMonth  string
	Description string
	ImageUrls   pq.StringArray
	CreatedAt   string
	HasMatched  bool
	Id          string
	OwnerId     int64
}
