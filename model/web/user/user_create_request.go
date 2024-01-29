package user

type CreateRequestUser struct {
	Email           string `json:"email" validate:"required,email"`
	Nim             string `json:"nim" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Bio             string `json:"bio"`
	Password        string `json:"password" validate:"required,min=8,containsany,containsuppercase,containslowercase,containsnumeric"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
