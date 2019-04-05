package validator

import (
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

func RegisterEmailTranslation(ut ut.Translator) error {
	return ut.Add("email", "Enter a valid email address", true)
}

func EmailTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("email", fe.Field())
	return t
}
