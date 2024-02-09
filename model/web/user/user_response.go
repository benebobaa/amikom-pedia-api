package user

import (
	"amikom-pedia-api/model/web/image"
	"time"
)

type ResponseUser struct {
	UUID      string                      `json:"uuid"`
	Email     string                      `json:"email"`
	Nim       string                      `json:"nim"`
	Name      string                      `json:"name"`
	Username  string                      `json:"username"`
	Bio       string                      `json:"bio"`
	Images    []image.CreateImageResponse `json:"images"`
	CreatedAt time.Time                   `json:"created_at"`
	UpdatedAt time.Time                   `json:"updated_at"`
}
