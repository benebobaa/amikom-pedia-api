package domain

import "time"

type Otp struct {
	ID        int
	User_rid  int
	UUID      string
	OTP_code  string
	ExpiredAt time.Time
	CreatedAt time.Time
	Ref_code  string
}
