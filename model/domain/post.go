package domain

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        string
	Content   string
	UserId    string
	RefPostId sql.NullString
	Images    []Image
	CreatedAt time.Time
	UpdatedAt time.Time
}
