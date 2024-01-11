package register_service

import (
	"amikom-pedia-api/model/web/register"
	"context"
)

type RegisterService interface {
	Create(ctx context.Context, requestRegister register.RegisterRequest) register.RegisterResponse
}
