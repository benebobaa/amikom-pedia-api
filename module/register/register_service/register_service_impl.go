package register_service

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/register"
	"amikom-pedia-api/module/register/register_repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type RegisterServiceImpl struct {
	RegisterRepository register_repository.RegisterRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewRegisterService(registerRepository register_repository.RegisterRepository, DB *sql.DB, validate *validator.Validate) *RegisterServiceImpl {
	return &RegisterServiceImpl{RegisterRepository: registerRepository, DB: DB, Validate: validate}
}

func (registerService RegisterServiceImpl) Create(ctx context.Context, requestRegister register.RegisterRequest) register.RegisterResponse {
	err := registerService.Validate.Struct(requestRegister)
	helper.PanicIfError(err)

	tx, err := registerService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	requestRegisterDomain := domain.Register{
		Email:    requestRegister.Email,
		Nim:      requestRegister.Nim,
		Name:     requestRegister.Name,
		Password: requestRegister.Password,
	}

	result := registerService.RegisterRepository.Create(ctx, tx, requestRegisterDomain)

	return helper.ToRegisterResponse(result)
}
