package v10Validator

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) ValidateStruct(i interface{}) error {
	return cv.validator.Struct(i)
}

func (cv *CustomValidator) Engine() interface{} {
	return cv.validator
}
