package domain

import (
	"database/sql"
	"time"
)

type Register struct {
	ID              int
	Name            string
	Email           string
	Nim             string
	Password        string
	ConfirmPassword string
	RefCode         string
	IsVerified      bool
	EmailVerifiedAt sql.NullTime
	CreatedAt       time.Time
}
