package user

type UpdatePasswordRequest struct {
	UUID               string `json:"uuid"`
	CurrentPassword    string `json:"current_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required,min=8,containsany,containsuppercase,containslowercase,containsnumeric"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required,eqfield=NewPassword"`
}
