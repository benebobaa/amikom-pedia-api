package otp

type OtpValidateRequest struct {
	RefCode  string `json:"ref_code" validate:"required"`
	OtpValue string `json:"otp_value" validate:"required,min=6,max=6"`
}
