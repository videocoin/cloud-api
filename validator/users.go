package validator

import (
	"github.com/VideoCoin/cloud-api/rpc"
	v1 "github.com/VideoCoin/cloud-api/users/v1"
	ut "github.com/go-playground/universal-translator"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
	validator "gopkg.in/go-playground/validator.v9"
)

func RegisterUserEmailTranslation(ut ut.Translator) error {
	return ut.Add("email", "Enter a valid email address", true)
}

func UserEmailTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("email", fe.Field())
	return t
}

func RegisterUserNameTranslation(ut ut.Translator) error {
	return ut.Add("name", "Enter a name", true)
}

func UserNameTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("name", fe.Field())
	return t
}

func RegisterSecureUserPasswordTranslation(ut ut.Translator) error {
	return ut.Add("secure-password", "Password is too simple", true)
}

func SecureUserPasswordTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("secure-password", fe.Field())
	return t
}

func RegisterConfirmUserPasswordTranslation(ut ut.Translator) error {
	return ut.Add("confirm-password", "Password does not match", true)
}

func ConfirmUserPasswordTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("confirm-password", fe.Field())
	return t
}

func ValidateConfirmUserPassword(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()

	currentField, currentKind, ok := fl.GetStructFieldOK()
	if !ok || currentKind != kind {
		return false
	}

	return field.String() == currentField.String()
}

func ValidateSecureUserPassword(fl validator.FieldLevel) bool {
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

func ValidateCreateUserRequest(r *v1.CreateUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}

func ValidateLoginUserRequest(r *v1.LoginUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}

func ValidateStartRecoveryUserRequest(r *v1.StartRecoveryUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}

func ValidateRecoverUserRequest(r *v1.RecoverUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}

func ValidateResetPasswordUserRequest(r *v1.ResetPasswordUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}
