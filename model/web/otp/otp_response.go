package otp

type ResponseOTP struct {
	ID        int    `json:"id"`
	User_rid  int    `json:"user_rid"`
	UUID      string `json:"uuid"`
	OTP_code  string `json:"otp_code"`
	ExpiredAt string `json:"expired_at"`
	CreatedAt string `json:"created_at"`
	Ref_code  string `json:"ref_code"`
}
