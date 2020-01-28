package vsz

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	pkgValidator *validator.Validate
)

func init() {
	pkgValidator = validator.New()
}

// PackageValidator returns the global validator instance used by this client
func PackageValidator() *validator.Validate {
	return pkgValidator
}
