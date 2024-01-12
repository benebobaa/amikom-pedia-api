package otp_service

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/otp"
	"amikom-pedia-api/module/otp/otp_repository"
	"amikom-pedia-api/module/register/register_repository"
	"amikom-pedia-api/module/user/user_repository"
	"amikom-pedia-api/utils/mail"
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
}

func NewOtpService(otpRepository otp_repository.OtpRepository, registerRepository register_repository.RegisterRepository, gmailSender mail.EmailSender, DB *sql.DB, validate *validator.Validate) OtpService {
	return &OtpServiceImpl{OtpRepository: otpRepository, RegisterRepository: registerRepository, GmailSender: gmailSender, DB: DB, Validate: validate}
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

func (otpService *OtpServiceImpl) Validation(ctx context.Context, request otp.OtpValidateRequest) {
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

	paramsUpdate := domain.Register{
		IsVerified:      true,
		EmailVerifiedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:              int(result.UserRid.Int32),
	}

	otpService.RegisterRepository.Update(ctx, tx, paramsUpdate)
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

	var toEmail string

	if result.UserRid.Valid {
		toEmail = result.EmailUserRegister.String
		err = otpService.sendingOtp(toEmail, result.OtpValue)
		if err != nil {
			panic(exception.NewOtpError(err.Error()))
		}
		return nil
	} else if result.UUID.Valid {
		toEmail = result.EmailUser.String
		err = otpService.sendingOtp(toEmail, result.OtpValue)
		if err != nil {
			panic(exception.NewOtpError(err.Error()))
		}
		return nil
	}

	return errors.New("email not found")
}

func (otpService *OtpServiceImpl) sendingOtp(email string, otpCode string) error {
	subject := "OTP Verification for Amikom Pedia"

	// Use fmt.Sprintf to dynamically insert values into the content string
	content := fmt.Sprintf(`
        <h1>Hello %s,</h1>

        <p>We're excited to have you on board with Amikom Pedia! As part of our security measures, please use the following One-Time Password (OTP) to verify your account:</p>

        <h2>%s</h2>

        <p>This OTP is valid for a single use and will expire in 1 minute. Please do not share it with anyone for security reasons.</p>

        <p>If you did not attempt to create an account with Amikom Pedia, please disregard this email. Your account security is important to us.</p>

        <p>Thank you for choosing Amikom Pedia!</p>
    `, email, otpCode)

	to := []string{email}
	err := otpService.GmailSender.SendEmail(subject, content, to, []string{}, []string{}, []string{})
	if err != nil {
		return errors.New("failed to send email")
	}
	return nil
}
