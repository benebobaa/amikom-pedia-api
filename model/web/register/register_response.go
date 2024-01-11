package register

import "time"

type RegisterResponse struct {
	ID              int       `json:"id"`
	Name            string    `json:"name" validate:"required"`
	Email           string    `json:"email" validate:"required,email"`
	Nim             string    `json:"nim" validate:"required"`
	RefCode         string    `json:"ref_code"`
	IsVerified      bool      `json:"is_verified"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	CreatedAt       time.Time `json:"created_at"`
}
