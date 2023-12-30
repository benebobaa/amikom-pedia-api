package user

type CreateRequestUser struct {
	Id          string `validate:"required"`
	Username    string `validate:"required"`
	DisplayName string `validate:"required"`
	Email       string `validate:"required"`
	Password    string `validate:"required"`
}
