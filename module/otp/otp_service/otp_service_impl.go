package otp_service

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/otp"
	"amikom-pedia-api/module/otp/otp_repository"
	"amikom-pedia-api/module/register/register_repository"
	"amikom-pedia-api/module/user/user_repository"
	"amikom-pedia-api/utils"
	"amikom-pedia-api/utils/mail"
	"amikom-pedia-api/utils/token"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type OtpServiceImpl struct {
	OtpRepository      otp_repository.OtpRepository
	RegisterRepository register_repository.RegisterRepository
	UserRepository     user_repository.UserRepository
	GmailSender        mail.EmailSender
	DB                 *sql.DB
	Validate           *validator.Validate
	TokenMaker         token.Maker
}

func NewOtpService(otpRepository otp_repository.OtpRepository, registerRepository register_repository.RegisterRepository, userRepository user_repository.UserRepository, gmailSender mail.EmailSender, DB *sql.DB, validate *validator.Validate, tokenMaker token.Maker) OtpService {
	return &OtpServiceImpl{OtpRepository: otpRepository, RegisterRepository: registerRepository, UserRepository: userRepository, GmailSender: gmailSender, DB: DB, Validate: validate, TokenMaker: tokenMaker}
}

func (otpService *OtpServiceImpl) Create(ctx context.Context, request otp.CreateRequestOtp) otp.CreateResponseOTP {
	err := otpService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := otpService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	otpData := domain.Otp{
		RefCode:   request.RefCode,
		OtpValue:  request.OtpValue,
		ExpiredAt: request.ExpiredAt,
		UserRid:   request.UserRid,
	}

	result := otpService.OtpRepository.Create(ctx, tx, otpData)

	return helper.ToOtpResponse(result)
}

func (otpService *OtpServiceImpl) Validation(ctx context.Context, request otp.OtpValidateRequest) otp.CreateResponseWithToken {
	err := otpService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := otpService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := otpService.OtpRepository.FindByRefCode(ctx, tx, request.RefCode)

	if err != nil {
		fmt.Println("otp not found")
		panic(exception.NewOtpError(err.Error()))
	}

	if result.ExpiredAt.Before(time.Now()) {
		fmt.Println("otp expired")
		panic(exception.NewOtpError("OTP expired. Please request a new one."))
	}

	if result.OtpValue != request.OtpValue {
		fmt.Println("otp wrong")
		panic(exception.NewOtpError("Invalid OTP. Please enter a valid one."))
	}

	if result.UserRid.Valid {
		requestUserDomain := domain.User{
			Email:    result.EmailUserRegister.String,
			Nim:      result.Nim.String,
			Name:     result.Name.String,
			Username: utils.RandomString(7),
			Password: result.Password.String}

		paramsUpdate := domain.Register{
			IsVerified:      true,
			EmailVerifiedAt: sql.NullTime{Time: time.Now(), Valid: true},
			ID:              int(result.UserRid.Int32),
		}
		otpService.RegisterRepository.Update(ctx, tx, paramsUpdate)
		userCreate := otpService.UserRepository.Create(ctx, tx, requestUserDomain)
		accessToken, err := otpService.TokenMaker.CreateToken(userCreate.Username, userCreate.UUID)
		helper.PanicIfError(err)

		return helper.ToOtpResponseWithToken(sql.NullString{Valid: true, String: accessToken})
	}
	if result.UUID.Valid {
		return helper.ToOtpResponseWithToken(sql.NullString{Valid: false})
	}

	return helper.ToOtpResponseWithToken(sql.NullString{Valid: false})

}

func (otpService *OtpServiceImpl) SendOtp(ctx context.Context, request otp.SendOtpRequest) error {
	err := otpService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := otpService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	result, err := otpService.OtpRepository.FindByRefCode(ctx, tx, request.RefCode)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//var toEmail string

	if result.UserRid.Valid {
		//toEmail = result.EmailUserRegister.String
		//err = otpService.sendingOtp(toEmail, result.OtpValue)
		if err != nil {
			panic(exception.NewOtpError(err.Error()))
		}
		return nil
	} else if result.UUID.Valid {
		//toEmail = result.EmailUser.String
		//err = otpService.sendingOtp(toEmail, result.OtpValue)
		if err != nil {
			panic(exception.NewOtpError(err.Error()))
		}
		return nil
	}

	return errors.New("email not found")
}

func (otpService *OtpServiceImpl) ResendOtp(ctx context.Context, request otp.SendOtpRequest) error {
	err := otpService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := otpService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	newOtp := utils.GenerateNewOtpValue()

	result, err := otpService.OtpRepository.FindByRefCode(ctx, tx, request.RefCode)
	if err != nil {
		panic(exception.NewOtpError(err.Error()))
	}

	result.OtpValue = newOtp
	err = otpService.OtpRepository.Update(ctx, tx, result)
	if err != nil {
		return err
	}

	var toEmail string

	if result.UserRid.Valid {
		toEmail = result.EmailUserRegister.String

		subject, content, toEmail := mail.GetSenderParamEmailRegist(toEmail, newOtp)
		err = otpService.GmailSender.SendEmail(subject, content, toEmail, []string{}, []string{}, []string{})

		if err != nil {
			panic(exception.NewOtpError(err.Error()))
		}
		return nil
	} else if result.UUID.Valid {
		toEmail = result.EmailUser.String

		subject, content, toEmail := mail.GetSenderParamEmailForgotPass(toEmail, newOtp)
		err = otpService.GmailSender.SendEmail(subject, content, toEmail, []string{}, []string{}, []string{})

		if err != nil {
			panic(exception.NewOtpError(err.Error()))
		}
		return nil
	}

	return errors.New("email not found")

}
