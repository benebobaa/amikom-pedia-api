package login

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required,amikom"`
	Password        string `json:"password" validate:"required"`
}
