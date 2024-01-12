package otp

import "time"

type CreateResponseOTP struct {
	ID        int       `json:"id"`
	UserRid   int       `json:"user_rid"`
	UUID      string    `json:"uuid"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	RefCode   string    `json:"ref_code"`
}
