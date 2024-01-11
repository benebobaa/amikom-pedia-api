package otp

import (
	"database/sql"
	"time"
)

type CreateRequestOtp struct {
	OtpValue  string         `json:"otp_code" validate:"required"`
	RefCode   string         `json:"ref_code" validate:"required"`
	ExpiredAt time.Time      `json:"expired_at" validate:"required"`
	UserRid   sql.NullInt32  `json:"user_rid" validate:"required"`
	UUID      sql.NullString `json:"uuid" validate:"required"`
}
