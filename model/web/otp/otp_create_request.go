package otp

type CreateRequestOtp struct {
	OTP_code string `json:"otp_code" validate:"required"`
	Ref_code string `json:"ref_code" validate:"required"`
}
