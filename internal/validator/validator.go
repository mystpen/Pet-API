package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	
)

var UserNameRX = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9_.]{4,16}[a-zA-Z0-9])?$`)
// var PasswordRX = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`)

func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("password", ValidatePassword)

	return validate
}

func ValidatePassword(field validator.FieldLevel) bool {
	password := field.Field().Interface().(string)
	return password != ""
}

func ValidateUserName(field validator.FieldLevel) bool {
	username := field.Field().Interface().(string)

	return UserNameRX.MatchString(username)
}
