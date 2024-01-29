package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"unicode"
)

func validateAmikomEmail(field validator.FieldLevel) bool {
	email := field.Field().Interface().(string)
	return strings.HasSuffix(email, "@student.amikom.ac.id") || strings.HasSuffix(email, "@amikom.ac.id")
}

func containsAny(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	specialChars := "!@#$%^&*()_+~`"
	for _, char := range specialChars {
		if strings.ContainsRune(field, char) {
			return true
		}
	}
	return false
}

func containsLowercase(fl validator.FieldLevel) bool {
	return strings.ToLower(fl.Field().String()) != fl.Field().String()
}

func containsUppercase(fl validator.FieldLevel) bool {
	return strings.ToUpper(fl.Field().String()) != fl.Field().String()
}

func containsNumeric(fl validator.FieldLevel) bool {
	return strings.IndexFunc(fl.Field().String(), unicode.IsNumber) != -1
}

func CustomValidator() *validator.Validate {
	validate := validator.New()
	_ = validate.RegisterValidation("amikom", validateAmikomEmail)
	_ = validate.RegisterValidation("containsany", containsAny)
	_ = validate.RegisterValidation("containslowercase", containsLowercase)
	_ = validate.RegisterValidation("containsuppercase", containsUppercase)
	_ = validate.RegisterValidation("containsnumeric", containsNumeric)

	return validate
}
