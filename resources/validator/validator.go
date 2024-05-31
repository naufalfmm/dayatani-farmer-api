package validator

import (
	"github.com/naufalfmm/dayatani-farmer-api/utils/validator"
	"github.com/naufalfmm/dayatani-farmer-api/utils/validator/v10Validator"
)

func NewValidator() (validator.Validator, error) {
	return v10Validator.NewV10()
}
