package otp_repository

import (
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
)

type OtpRepository interface {
	Create(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp
	FindByRefCode(ctx context.Context, tx *sql.Tx, refCode string) (domain.Otp, error)
	Update(ctx context.Context, tx *sql.Tx, otp domain.Otp) error
}
