package user

type CreateUpdatePassword struct {
	Id       string `validate:"required"`
	Password string `validate:"required"`
}
