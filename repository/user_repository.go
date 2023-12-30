package repository

import (
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, userId string)
	Update(ctx context.Context, tx *sql.Tx, category domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
