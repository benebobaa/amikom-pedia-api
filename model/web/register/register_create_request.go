package register

type RegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email,amikom"`
	Nim             string `json:"nim" validate:"required"`
	Password        string `json:"password" validate:"required,min=8,containsany,containsuppercase,containslowercase,containsnumeric"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
