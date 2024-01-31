package domain

import "database/sql"

type Image struct {
	ID        int
	UserID    sql.NullString
	FilePath  string
	ImageType string
	ImageUrl  string
	PostID    sql.NullString
	CreatedAt string
	UpdatedAt string
}
