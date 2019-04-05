package validator

import (
	ut "github.com/go-playground/universal-translator"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
	validator "gopkg.in/go-playground/validator.v9"
)

func RegisterSecurePasswordTranslation(ut ut.Translator) error {
	return ut.Add("secure-password", "Password is too simple", true)
}

func SecurePasswordTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("secure-password", fe.Field())
	return t
}

func RegisterConfirmPasswordTranslation(ut ut.Translator) error {
	return ut.Add("confirm-password", "Password does not match", true)
}

func ConfirmPasswordTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("confirm-password", fe.Field())
	return t
}

func ValidateConfirmPassword(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()

	currentField, currentKind, ok := fl.GetStructFieldOK()
	if !ok || currentKind != kind {
		return false
	}

	return field.String() == currentField.String()
}

func ValidateSecurePassword(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.String() == "" {
		return false
	}

	s := zxcvbn.PasswordStrength(field.String(), nil)
	if s.Score <= 2 {
		return false
	}

	return true
}
