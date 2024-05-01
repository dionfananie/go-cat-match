package cat

import (
	"database/sql"

	"github.com/lib/pq"
)

type Sex string

type Race string


type ListCat struct {
	Name        string
	Race        Race
	Sex         Sex
	AgeInMonth  string
	Description string
	ImageUrls   pq.StringArray
	CreatedAt   string
	HasMatched  bool
	Id          string
	OwnerId     sql.NullString
}
