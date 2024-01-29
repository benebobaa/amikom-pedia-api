package user

type SetNewPasswordRequest struct {
	RefCode         string `json:"ref_code" validate:"required"`
	Password        string `json:"password" validate:"required,min=8,containsany,containsuppercase,containslowercase,containsnumeric"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
