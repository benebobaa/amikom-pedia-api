package register_repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
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
