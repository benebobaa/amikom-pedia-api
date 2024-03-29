package register_service

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/register"
	"amikom-pedia-api/module/otp/otp_repository"
	"amikom-pedia-api/module/register/register_repository"
	"amikom-pedia-api/utils"
	"amikom-pedia-api/utils/mail"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

type RegisterServiceImpl struct {
	RegisterRepository register_repository.RegisterRepository
	OtpRepository      otp_repository.OtpRepository
	GmailSender        mail.EmailSender
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewRegisterService(registerRepository register_repository.RegisterRepository, otpRepository otp_repository.OtpRepository, gmailSender mail.EmailSender, DB *sql.DB, validate *validator.Validate) RegisterService {
	return &RegisterServiceImpl{RegisterRepository: registerRepository, OtpRepository: otpRepository, GmailSender: gmailSender, DB: DB, Validate: validate}
}

func (registerService *RegisterServiceImpl) Create(ctx context.Context, requestRegister register.RegisterRequest) register.RegisterResponse {
	err := registerService.Validate.Struct(requestRegister)
	helper.PanicIfError(err)

	tx, err := registerService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	hashedPassword, err := utils.HashPassword(requestRegister.Password)
	helper.PanicIfError(err)

	requestRegisterDomain := domain.Register{
		Email:    requestRegister.Email,
		Nim:      requestRegister.Nim,
		Name:     requestRegister.Name,
		Password: hashedPassword,
	}

	_, err = registerService.RegisterRepository.FindByEmail(ctx, tx, requestRegister.Email)

	if err != nil {
		panic(exception.NewEmailAlreadyVerified(err.Error()))
	}

	result := registerService.RegisterRepository.Create(ctx, tx, requestRegisterDomain)

	resultNulId := sql.NullInt32{Int32: int32(result.ID), Valid: true}

	otpData := domain.Otp{
		RefCode:   utils.RandomCombineIntAndString(),
		OtpValue:  strconv.FormatInt(utils.RandomInt(100000, 999999), 10),
		ExpiredAt: time.Now().Add(time.Minute * 5),
		UserRid:   resultNulId,
	}

	resultOTP := registerService.OtpRepository.Create(ctx, tx, otpData)

	subject, content, toEmail := mail.GetSenderParamEmailRegist(result.Email, resultOTP.OtpValue)

	err = registerService.GmailSender.SendEmail(subject, content, toEmail, []string{}, []string{}, []string{})
	fmt.Println("Error", err)
	helper.PanicIfError(err)
	return helper.ToRegisterResponse(result, resultOTP)
}
