package vsz

// API Version: v8_1

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	pkgValidator *validator.Validate
)

func init() {
	pkgValidator = validator.New()
}
