package service

import (
	"amikom-pedia-api/model/web/user"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request user.CreateRequestUser) user.ResponseUser
	Update(ctx context.Context, request user.CreateUpdatePassword) user.ResponseUser
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) user.ResponseUser
	FindAll(ctx context.Context) []user.ResponseUser
}
