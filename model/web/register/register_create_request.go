package register

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Nim      string `json:"nim" validate:"required"`
	Password string `json:"password" validate:"required"`
}
