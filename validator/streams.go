package validator

import (
	"github.com/VideoCoin/cloud-api/rpc"
	v1 "github.com/VideoCoin/cloud-api/streams/v1"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

func RegisterStreamNameTranslation(ut ut.Translator) error {
	return ut.Add("name", "Enter a name", true)
}

func StreamNameTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("name", fe.Field())
	return t
}

func ValidateCreateStreamRequest(r *v1.CreateStreamRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}
