package domain

import "time"

type User struct {
	Id          string
	Username    string
	DisplayName string
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdateAt    time.Time
}
