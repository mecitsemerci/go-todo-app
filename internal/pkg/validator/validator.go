package validator

import "github.com/go-playground/validator/v10"

//Validate given object
func Validate(s interface{}) error {
	v := validator.New()

	return v.Struct(s)
}
