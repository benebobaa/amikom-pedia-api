package user_service

import (
	"amikom-pedia-api/model/web/user"
	"context"
)

type UserService interface {
	Create(ctx context.Context, requestUser user.CreateRequestUser) user.ResponseUser
	Update(ctx context.Context, uuid string, requestUser user.UpdateRequestUser) user.ResponseUser
	Delete(ctx context.Context, uuid string)
	FindByUUID(ctx context.Context, uuid string) user.ResponseUser
	FindAll(ctx context.Context) []user.ResponseUser
	ForgotPassword(ctx context.Context, email string) user.ForgotPasswordResponse
	SetNewPassword(ctx context.Context, requestSetNewPassword user.SetNewPasswordRequest)
	UpdatePassword(ctx context.Context, uuid string, newPasswordRequest user.UpdatePasswordRequest) error
}
