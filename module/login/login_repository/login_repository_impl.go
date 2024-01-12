package login_repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type LoginRepositoryImpl struct{}

func NewLoginRepository() LoginRepository {
	return &LoginRepositoryImpl{}
}

func (loginRepo *LoginRepositoryImpl) FindUserByUsernameOrEmail(ctx context.Context, tx *sql.Tx, login domain.Login) (domain.User, error) {
	SQL := `SELECT uuid, email, nim, name, username, bio, password, created_at, updated_at FROM "user" WHERE username = $1 OR email = $2`

	rows, err := tx.QueryContext(ctx, SQL, login.UsernameOrEmail, login.UsernameOrEmail)
	fmt.Println("err:", err)
	helper.PanicIfError(err)
	defer rows.Close()

	userData := domain.User{}
	if rows.Next() {
		err = rows.Scan(&userData.UUID, &userData.Email, &userData.Nim, &userData.Name, &userData.Username, &userData.Bio, &userData.Password, &userData.CreatedAt, &userData.UpdatedAt)
		helper.PanicIfError(err)
		return userData, nil
	} else {
		return userData, errors.New("username or email not found")
	}

}
