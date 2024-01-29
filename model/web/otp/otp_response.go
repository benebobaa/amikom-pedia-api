package otp

import (
	"time"
)

type CreateResponseResetPasswordOtp struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	RefCode   string    `json:"ref_code"`
}

type CreateResponseWithToken struct {
	AccessToken string `json:"access_token"`
}
