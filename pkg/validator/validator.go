package validator

import "github.com/go-playground/validator/v10"

var validate *validator.Validate = validator.New()

func GetValidate() *validator.Validate {
	return validate
}
