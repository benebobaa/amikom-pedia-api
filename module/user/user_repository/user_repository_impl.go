package user_repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (userRepo *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `INSERT INTO "user" (email, nim, name, username, bio, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid, created_at, updated_at`
	fmt.Println("SQL:", SQL)
	row := tx.QueryRowContext(ctx, SQL, user.Email, user.Nim, user.Name, user.Username, user.Bio, user.Password)
	fmt.Println("row:", row)
	err := row.Scan(&user.UUID, &user.CreatedAt, &user.UpdatedAt)
	fmt.Println("err:", err)
	helper.PanicIfError(err)

	return user
}

func (userRepo *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) {
}

func (userRepo *UserRepositoryImpl) FindByUUID(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := `SELECT uuid, email, nim, name, username, bio, password, created_at, updated_at FROM "user" WHERE uuid = $1`

	rows, err := tx.QueryContext(ctx, SQL, user.UUID)
	fmt.Println("err:", err)
	helper.PanicIfError(err)
	defer rows.Close()

	userData := domain.User{}
	if rows.Next() {
		rows.Scan(&userData.UUID, &userData.Email, &userData.Nim, &userData.Name, &userData.Username, &userData.Bio, &userData.Password, &userData.CreatedAt, &userData.UpdatedAt)
		return userData, nil
	} else {
		return userData, errors.New("user not found")
	}
}

func (userRepo *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, user domain.User) []domain.User {
	//TODO implement me
	panic("implement me")
}
