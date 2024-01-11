package otp_service

import (
	"amikom-pedia-api/model/web/otp"
	"context"
)

type OtpService interface {
	Create(ctx context.Context, request otp.CreateRequestOtp) otp.CreateResponseOTP
	Validation(ctx context.Context, request otp.OtpValidateRequest)
}
