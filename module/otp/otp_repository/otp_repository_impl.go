package otp_repository

import (
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
)

type OtpRepositoryImpl struct{}

func NewOtpRepository() OtpRepository {
	return &OtpRepositoryImpl{}
}

func (otpRepo *OtpRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp {
	SQL := `INSERT INTO "otp" (user_id, otp_value, ref_code) `
}

func (otpRepo *OtpRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp {
	//TODO implement me
	panic("implement me")
}
