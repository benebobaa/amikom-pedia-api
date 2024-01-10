package user

import "time"

type ResponseUser struct {
	UUID      string    `json:"uuid"`
	Email     string    `json:"email"`
	Nim       string    `json:"nim"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
