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
}

func validUsername(fl validator.FieldLevel) bool {
	return UsernameRegex.Match(fl.Field().Bytes())
}

func validPassword(fl validator.FieldLevel) bool {
    return PasswordRegex.Match(fl.Field().Bytes())
}
