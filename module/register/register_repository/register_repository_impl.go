package register_repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RegisterRepositoryImpl struct{}

func NewRegisterRepository() RegisterRepository {
	return &RegisterRepositoryImpl{}
}

func (registerRepo *RegisterRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, register domain.Register) domain.Register {
	SQL := `INSERT INTO "user_registration" (email, nim, name, password) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	row := tx.QueryRowContext(ctx, SQL, register.Email, register.Nim, register.Name, register.Password)
	err := row.Scan(&register.ID, &register.CreatedAt)
	helper.PanicIfError(err)

	return register
}

func (registerRepo *RegisterRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, register domain.Register) domain.Register {
	SQL := `UPDATE "user_registration"
			SET is_verified = $1, email_verified_at = $2
			WHERE id = $3`

	_, err := tx.ExecContext(ctx, SQL, register.IsVerified, register.EmailVerifiedAt, register.ID)

	helper.PanicIfError(err)

	return register
}

func (registerRepo *RegisterRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.Register, error) {
	SQL := `SELECT is_verified, email_verified_at FROM "user_registration" WHERE email = $1`

	row, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer row.Close()

	register := domain.Register{}
	if row.Next() {
		err := row.Scan(&register.IsVerified, &register.EmailVerifiedAt)
		if register.IsVerified {
			fmt.Printf("Email %s already verified\n", email)
			return register, errors.New("email already verified")
		}
		helper.PanicIfError(err)
	} else {
		return register, nil
	}

	return register, nil
}

func (registerRepo *RegisterRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (domain.Register, error) {
	SQL := `SELECT email FROM "user_registration" WHERE id = $1`

	row, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer row.Close()

	register := domain.Register{}
	if row.Next() {
		err := row.Scan(&register.Email)
		helper.PanicIfError(err)

	} else {
		return register, errors.New("email not found")
	}
	return register, nil
}
