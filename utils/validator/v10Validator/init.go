package v10Validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	"github.com/naufalfmm/dayatani-farmer-api/utils/validator"
)

func NewV10() (validator.Validator, error) {
	langEn := en.New()
	langId := id.New()
	uni := ut.New(langEn, langEn, langId)
	trans, _ := uni.GetTranslator("en")

	validate := v10.New()
	if err := enTrans.RegisterDefaultTranslations(validate, trans); err != nil {
		return nil, err
	}

	return &CustomValidator{
		validator: validate,
	}, nil
}
