package domain

import (
	"database/sql"
	"time"
)

type Otp struct {
	ID                int
	UserRid           sql.NullInt32
	UUID              sql.NullString
	Nim               sql.NullString
	Name              sql.NullString
	Password          sql.NullString
	EmailUserRegister sql.NullString
	EmailUser         sql.NullString
	OtpValue          string
	RefCode           string
	ExpiredAt         time.Time
	CreatedAt         time.Time
}
