package types

import (
	vCop "gopkg.in/go-playground/validator.v9"
)

var validator *vCop.Validate

func init() {
	validator = vCop.New()
}

// Validate - validates an object based on it's tags
func Validate(t interface{}) error {
	return validator.Struct(t)
}
