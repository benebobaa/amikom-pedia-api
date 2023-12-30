package service

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/user"
	"amikom-pedia-api/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (userServiceImpl *UserServiceImpl) Create(ctx context.Context, request user.CreateRequestUser) user.ResponseUser {
	err := userServiceImpl.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := userServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = userServiceImpl.UserRepository.FindById(ctx, tx, request.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	dataUser := domain.User{Id: request.Id, Username: request.Username, DisplayName: request.DisplayName, Email: request.Email, Password: request.Password}

	dataUser = userServiceImpl.UserRepository.Create(ctx, tx, dataUser)

	dataUser, err = userServiceImpl.UserRepository.FindById(ctx, tx, dataUser.Id)
	helper.PanicIfError(err)

	return helper.ToUserResponse(dataUser)
}

func (userServiceImpl *UserServiceImpl) Delete(ctx context.Context, userId string) {
	tx, err := userServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := userServiceImpl.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userServiceImpl.UserRepository.Delete(ctx, tx, category.Id)
}

func (userServiceImpl *UserServiceImpl) Update(ctx context.Context, request user.CreateUpdatePassword) user.ResponseUser {
	err := userServiceImpl.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := userServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userData, err := userServiceImpl.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userData.Password = request.Password

	userData = userServiceImpl.UserRepository.Update(ctx, tx, userData)

	return helper.ToUserResponse(userData)
}

func (userServiceImpl *UserServiceImpl) FindById(ctx context.Context, userId string) user.ResponseUser {
	tx, err := userServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := userServiceImpl.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(category)
}

func (userServiceImpl *UserServiceImpl) FindAll(ctx context.Context) []user.ResponseUser {
	tx, err := userServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := userServiceImpl.UserRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
