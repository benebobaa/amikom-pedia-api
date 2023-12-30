package repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
func (userRepository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(id,username,display_name,email,password) VALUES($1,$2,$3,$4,$5)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.DisplayName, user.Email, user.Password)

	helper.PanicIfError(err)

	//id, err := result.LastInsertId()

	//helper.PanicIfError(err)

	//user.Id = int(id)
	return user
}

func (userRepository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId string) {
	SQL := "DELETE from users WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(err)
}

func (userRepository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET password = $1,updated_at = CURRENT_TIMESTAMP WHERE id = $2"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (userRepository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (domain.User, error) {
	SQL := "SELECT id,username,display_name,email,password,created_at,updated_at FROM users WHERE id = $1"

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.DisplayName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdateAt)
		return user, nil
	} else {
		return user, errors.New("id not found")
	}
}

func (userRepository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id,username,display_name,email, password, created_at,updated_at FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.DisplayName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdateAt)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
