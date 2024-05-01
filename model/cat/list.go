package cat

import (
	"database/sql"

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
	OwnerId     sql.NullString
}
