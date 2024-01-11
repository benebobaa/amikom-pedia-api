package otp_service

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/otp"
	"amikom-pedia-api/module/otp/otp_repository"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type OtpServiceImpl struct {
	OtpRepository otp_repository.OtpRepository
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewOtpService(otpRepository otp_repository.OtpRepository, DB *sql.DB, validate *validator.Validate) OtpService {
	return &OtpServiceImpl{OtpRepository: otpRepository, DB: DB, Validate: validate}
}

func (otpServices *OtpServiceImpl) Create(ctx context.Context, request otp.CreateRequestOtp) otp.CreateResponseOTP {
	err := otpServices.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := otpServices.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	otpData := domain.Otp{
		RefCode:   request.RefCode,
		OtpValue:  request.OtpValue,
		ExpiredAt: request.ExpiredAt,
		UserRid:   request.UserRid,
	}

	result := otpServices.OtpRepository.Create(ctx, tx, otpData)

	return helper.ToOtpResponse(result)
}

func (otpServices *OtpServiceImpl) Validation(ctx context.Context, request otp.OtpValidateRequest) {
	err := otpServices.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := otpServices.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := otpServices.OtpRepository.FindByRefCode(ctx, tx, request.RefCode)

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
}
