package validator

import (
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

func RegisterNameTranslation(ut ut.Translator) error {
	return ut.Add("name", "Enter a name", true)
}

func NameTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("name", fe.Field())
	return t
}
