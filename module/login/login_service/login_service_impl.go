package login_service

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/login"
	"amikom-pedia-api/module/login/login_repository"
	"amikom-pedia-api/utils"
	"amikom-pedia-api/utils/token"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type LoginServiceImpl struct {
	TokenMaker      token.Maker
	LoginRepository login_repository.LoginRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewLoginService(tokenMaker token.Maker, loginRepository login_repository.LoginRepository, DB *sql.DB, validate *validator.Validate) LoginService {
	return &LoginServiceImpl{TokenMaker: tokenMaker, LoginRepository: loginRepository, DB: DB, Validate: validate}
}

func (loginService *LoginServiceImpl) Login(ctx context.Context, request login.LoginRequest) login.LoginResponse {
	err := loginService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := loginService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	loginDomain := domain.Login{
		UsernameOrEmail: request.UsernameOrEmail,
		Password:        request.Password,
	}

	result, err := loginService.LoginRepository.FindUserByUsernameOrEmail(ctx, tx, loginDomain)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	fmt.Println("result:", result)
	err = utils.CheckPassword(request.Password, result.Password)
	if err != nil {
		panic(exception.NewUnauthorizedError("wrong password"))
	}

	accessToken, err := loginService.TokenMaker.CreateToken(result.Username, result.UUID)

	helper.PanicIfError(err)

	return helper.ToLoginResponse(result, accessToken)
}
