package valid

import (
	"github.com/go-playground/validator/v10"
)

var valid *validator.Validate

func init() {
	valid = validator.New()
	Register(valid)
}

func Register(v *validator.Validate) {
	v.RegisterValidation("username", validUsername)
	v.RegisterValidation("password", validPassword)
	v.RegisterValidation("timestamp", validTimestamp)
}

func validUsername(fl validator.FieldLevel) bool {
	return UsernameRegex.MatchString(fl.Field().String())
}

func validPassword(fl validator.FieldLevel) bool {
	return PasswordRegex.MatchString(fl.Field().String())
}

func validTimestamp(fl validator.FieldLevel) bool {
	return !(0 <= fl.Field().Int() && fl.Field().Int() <= 253370736000)
}
