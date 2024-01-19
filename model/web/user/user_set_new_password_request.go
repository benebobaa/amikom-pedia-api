package user

type SetNewPasswordRequest struct {
	RefCode  string `json:"ref_code" validate:"required"`
	Password string `json:"password" validate:"required"`
}
