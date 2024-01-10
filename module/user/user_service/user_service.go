package user_service

import (
	"amikom-pedia-api/model/web/user"
	"context"
)

type UserService interface {
	Create(ctx context.Context, requestUser user.CreateRequestUser) user.ResponseUser
	Update(ctx context.Context) user.ResponseUser
	FindByUUID(ctx context.Context, uuid string) user.ResponseUser
	FindAll(ctx context.Context) []user.ResponseUser
}
