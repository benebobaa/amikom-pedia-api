package user_repository

import (
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindByUUID(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	SetNewPassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User, oldPassword string, newPassword string) error
}
