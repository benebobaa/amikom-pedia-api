package test

import (
	"amikom-pedia-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	password := "testPassword"
	wrongPassword := "wrongPassword"
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)

	t.Run("correct password", func(t *testing.T) {
		isCorrect := utils.CheckPasswordHash(password, hash)
		assert.True(t, isCorrect)
	})

	t.Run("incorrect password", func(t *testing.T) {
		isCorrect := utils.CheckPasswordHash(wrongPassword, hash)
		assert.False(t, isCorrect)
	})
}

func ValidateAmikomEmail(field validator.FieldLevel) bool {
	email := field.Field().Interface().(string)
	return strings.HasSuffix(email, "@student.amikom.ac.id") || strings.HasSuffix(email, "@amikom.ac.id")
}

func TestValidate(t *testing.T) {
	validate := validator.New()
	_ = validate.RegisterValidation("amikom", ValidateAmikomEmail)

	type Register struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email,amikom"`
		Password string `json:"password" validate:"required"`
	}

	request := Register{
		Name:     "test",
		Email:    "hanif@gmail.com",
		Password: "test",
	}

	err := validate.Struct(request)
	if err != nil {
		t.Error(err.Error())
	}
}
