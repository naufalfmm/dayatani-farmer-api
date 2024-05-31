package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/naufalfmm/dayatani-farmer-api/utils/validator/v10Validator/v10Err"
)

func HandleBindError(req any, err error) error {
	if _, ok := err.(validator.ValidationErrors); ok {
		return v10Err.NewFromValidationErrors(req, err)
	}

	return err
}
