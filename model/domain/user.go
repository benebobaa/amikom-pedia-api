package domain

import "time"

type User struct {
	UUID      string
	Email     string
	Nim       string
	Name      string
	Username  string
	Bio       string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	RefCode   string
}
