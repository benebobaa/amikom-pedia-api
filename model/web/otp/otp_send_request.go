package otp

type SendOtpRequest struct {
	RefCode string `json:"ref_code" validate:"required"`
}
