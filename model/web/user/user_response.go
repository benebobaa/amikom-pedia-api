package user

import "time"

type ResponseUser struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}
