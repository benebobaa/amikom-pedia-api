package login_repository

import (
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
)

type LoginRepository interface {
	FindUserByUsernameOrEmail(ctx context.Context, tx *sql.Tx, login domain.Login) (domain.User, error)
}
