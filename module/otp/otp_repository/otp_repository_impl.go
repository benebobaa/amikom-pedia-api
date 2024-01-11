package otp_repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type OtpRepositoryImpl struct{}

func NewOtpRepository() OtpRepository {
	return &OtpRepositoryImpl{}
}

func (otpRepo *OtpRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp {
	SQL := `INSERT INTO "otp" (user_id, user_rid, otp_value, expired_at, ref_code) VALUES ($1, $2, $3, $4, $5) RETURNING ref_code`
	rows := tx.QueryRowContext(ctx, SQL, otp.UUID, otp.UserRid, otp.OtpValue, otp.ExpiredAt, otp.RefCode)
	err := rows.Scan(&otp.RefCode)
	helper.PanicIfError(err)

	return otp
}

func (otpRepo *OtpRepositoryImpl) FindByRefCode(ctx context.Context, tx *sql.Tx, refCode string) (domain.Otp, error) {
	SQL := "SELECT id, ref_code, otp_value, expired_at, created_at, user_rid, user_id FROM otp WHERE ref_code = $1 AND expired_at > NOW()"
	rows, err := tx.QueryContext(ctx, SQL, refCode)
	helper.PanicIfError(err)
	defer rows.Close()

	otp := domain.Otp{}

	if rows.Next() {
		errScan := rows.Scan(&otp.ID, &otp.RefCode, &otp.OtpValue, &otp.ExpiredAt, &otp.CreatedAt, &otp.UserRid, &otp.UUID)
		fmt.Println("errScan", errScan)
		helper.PanicIfError(errScan)
		fmt.Println("otp", otp)
		return otp, nil
	} else {
		fmt.Println("otp", otp)
		return otp, errors.New("otp not found or expired")
	}
}
