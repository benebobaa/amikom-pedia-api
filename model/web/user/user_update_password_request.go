package user

type UpdatePasswordRequest struct {
	UUID            string `json:"uuid" validate:"required"`
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
}
