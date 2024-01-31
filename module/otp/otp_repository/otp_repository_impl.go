package otp_repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type OtpRepositoryImpl struct{}

func NewOtpRepository() OtpRepository {
	return &OtpRepositoryImpl{}
}

func (otpRepo *OtpRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp {
	SQL := `INSERT INTO "otp" (user_id, user_rid, otp_value, expired_at, ref_code) VALUES ($1, $2, $3, $4, $5) RETURNING ref_code, user_id, user_rid`
	rows := tx.QueryRowContext(ctx, SQL, otp.UUID, otp.UserRid, otp.OtpValue, otp.ExpiredAt, otp.RefCode)
	err := rows.Scan(&otp.RefCode, &otp.UUID, &otp.UserRid)
	helper.PanicIfError(err)

	return otp
}

func (otpRepo *OtpRepositoryImpl) FindByRefCode(ctx context.Context, tx *sql.Tx, refCode string) (domain.Otp, error) {
	SQL := `SELECT o.id, o.ref_code, o.otp_value, o.expired_at, o.created_at, o.user_rid, o.user_id, ur.nim ,ur.name, ur.password, ur.email AS user_registration_email, u.email AS user_email
				FROM
					"otp" AS o
				LEFT JOIN
					"user_registration" AS ur ON o.user_rid = ur.id
				LEFT JOIN
					"user" AS u ON o.user_id = u.uuid
				WHERE
					o.ref_code = $1`

	rows, err := tx.QueryContext(ctx, SQL, refCode)
	helper.PanicIfError(err)
	defer rows.Close()

	otp := domain.Otp{}

	if rows.Next() {
		errScan := rows.Scan(&otp.ID, &otp.RefCode, &otp.OtpValue, &otp.ExpiredAt, &otp.CreatedAt, &otp.UserRid, &otp.UUID, &otp.Nim, &otp.Name, &otp.Password, &otp.EmailUserRegister, &otp.EmailUser)
		helper.PanicIfError(errScan)
		return otp, nil
	} else {
		return otp, errors.New("refferal code not found")
	}
}

func (otpRepo *OtpRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, otp domain.Otp) error {
	SQL := `UPDATE "otp" SET otp_value = $1, expired_at = $2 WHERE ref_code = $3`
	_, err := tx.ExecContext(ctx, SQL, otp.OtpValue, otp.ExpiredAt, otp.RefCode)
	return err
}
