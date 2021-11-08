package validators

import (
	validators "github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateStruct(v interface{}) error
}
type validator struct {
	validator *validators.Validate
}

func NewValidator() Validator {
	return &validator{
		validator: validators.New(),
	}
}

func (vld *validator) ValidateStruct(v interface{}) error {
	return vld.validator.Struct(v)
}
