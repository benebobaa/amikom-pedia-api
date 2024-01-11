package login_service

import (
	"amikom-pedia-api/model/web/login"
	"context"
)

type LoginService interface {
	Login(ctx context.Context, request login.LoginRequest) login.LoginResponse
}
