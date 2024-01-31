package user

type UpdateRequestUser struct {
	UUID     string `json:"-"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Bio      string `json:"bio"`
}
