package domain

import "time"

type Register struct {
	ID              int
	Name            string
	Email           string
	Nim             string
	Password        string
	RefCode         string
	IsVerified      bool
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
}
