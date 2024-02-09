package domain

import (
	"database/sql"
	"time"
)

type Image struct {
	ID        sql.NullInt32
	UserID    sql.NullString
	FilePath  sql.NullString
	ImageType sql.NullString
	ImageUrl  sql.NullString
	PostID    sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}
