package user_service

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/user"
	"amikom-pedia-api/module/user/user_repository"
	"amikom-pedia-api/utils"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository user_repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (userService *UserServiceImpl) Create(ctx context.Context, requestUser user.CreateRequestUser) user.ResponseUser {
	err := userService.Validate.Struct(requestUser)
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	hashedPassword, err := utils.HashPassword(requestUser.Password)
	helper.PanicIfError(err)

	requestUserDomain := domain.User{
		Email:    requestUser.Email,
		Nim:      requestUser.Nim,
		Name:     requestUser.Name,
		Username: requestUser.Username,
		Bio:      requestUser.Bio,
		Password: hashedPassword}

	result := userService.UserRepository.Create(ctx, tx, requestUserDomain)

	return helper.ToUserResponse(result)
}

func (userService *UserServiceImpl) Update(ctx context.Context) user.ResponseUser {
	//TODO implement me
	panic("implement me")
}

func (userService *UserServiceImpl) FindByUUID(ctx context.Context, uuid string) user.ResponseUser {
	err := userService.Validate.Var(uuid, "required")

	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()

	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	userUUID := domain.User{UUID: uuid}

	result, err := userService.UserRepository.FindByUUID(ctx, tx, userUUID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(result)
}

func (userService *UserServiceImpl) FindAll(ctx context.Context) []user.ResponseUser {
	tx, err := userService.DB.Begin()

	helper.PanicIfError(err)

	users := userService.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}

func (userService *UserServiceImpl) Delete(ctx context.Context, uuid string) {
	err := userService.Validate.Var(uuid, "required")

	helper.PanicIfError(err)
	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	userUUID, err := userService.UserRepository.FindByUUID(ctx, tx, domain.User{UUID: uuid})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userService.UserRepository.Delete(ctx, tx, domain.User{UUID: userUUID.UUID})
}
