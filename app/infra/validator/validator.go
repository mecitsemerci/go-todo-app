package validator

import "github.com/go-playground/validator/v10"

func Validate(s interface{}) error {
	v := validator.New()

	return v.Struct(s)
}
