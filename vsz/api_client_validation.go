package vsz

import (
	"gopkg.in/go-playground/validator.v9"
)

var validation *validator.Validate

func init() {
	validation = validator.New()
}

func Validator() *validator.Validate {
	return validation
}
