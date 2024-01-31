package user_service

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/user"
	"amikom-pedia-api/module/otp/otp_repository"
	"amikom-pedia-api/module/user/user_repository"
	"amikom-pedia-api/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
	OtpRepository  otp_repository.OtpRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository user_repository.UserRepository, otpRepository otp_repository.OtpRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, OtpRepository: otpRepository, DB: DB, Validate: validate}
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

func (userService *UserServiceImpl) Update(ctx context.Context, uuid string, requestUser user.UpdateRequestUser) user.ResponseUser {
	err := userService.Validate.Struct(requestUser)
	helper.PanicIfError(err)

	fmt.Println("requestUser : ", requestUser)
	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := userService.UserRepository.FindByUUID(ctx, tx, domain.User{UUID: uuid})
	helper.PanicIfError(err)

	requestUserDomain := domain.User{
		UUID:     user.UUID,
		Name:     requestUser.Name,
		Username: requestUser.Username,
		Bio:      requestUser.Bio,
	}

	result := userService.UserRepository.Update(ctx, tx, requestUserDomain)
	return helper.ToUserResponse(result)
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

func (userService *UserServiceImpl) ForgotPassword(ctx context.Context, email string) user.ForgotPasswordResponse {
	err := userService.Validate.Var(email, "required,email")
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	userEmail := domain.User{Email: email}

	result, err := userService.UserRepository.FindByEmail(ctx, tx, userEmail)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	resultNullString := sql.NullString{String: result.UUID, Valid: true}

	otpData := domain.Otp{
		RefCode:   utils.RandomCombineIntAndString(),
		OtpValue:  strconv.FormatInt(utils.RandomInt(100000, 999999), 10),
		ExpiredAt: time.Now().Add(time.Minute * 5),
		UUID:      resultNullString,
	}

	resultOTP := userService.OtpRepository.Create(ctx, tx, otpData)

	return helper.ToSetNewPasswordResponse(resultOTP)

}

func (userService *UserServiceImpl) SetNewPassword(ctx context.Context, requestSetNewPassword user.SetNewPasswordRequest) {
	err := userService.Validate.Struct(requestSetNewPassword)
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	hashedPassword, err := utils.HashPassword(requestSetNewPassword.Password)
	helper.PanicIfError(err)

	result, err := userService.OtpRepository.FindByRefCode(ctx, tx, requestSetNewPassword.RefCode)
	helper.PanicIfError(err)

	requestSetNewPasswordDomain := domain.User{
		UUID:     result.UUID.String,
		Password: hashedPassword,
	}

	userService.UserRepository.SetNewPassword(ctx, tx, requestSetNewPasswordDomain)
}

func (userService *UserServiceImpl) UpdatePassword(ctx context.Context, uuid string, newPasswordRequest user.UpdatePasswordRequest) error {
	err := userService.Validate.Struct(newPasswordRequest)
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := userService.UserRepository.FindByUUID(ctx, tx, domain.User{UUID: uuid})
	helper.PanicIfError(err)

	if !utils.CheckPasswordHash(newPasswordRequest.CurrentPassword, user.Password) {
		return errors.New("current password does not match")
	}

	hashedPassword, err := utils.HashPassword(newPasswordRequest.NewPassword)
	helper.PanicIfError(err)

	user.Password = hashedPassword
	err = userService.UserRepository.UpdatePassword(ctx, tx, user, newPasswordRequest.CurrentPassword, newPasswordRequest.NewPassword)
	if err != nil {
		return err
	}

	return nil
}
