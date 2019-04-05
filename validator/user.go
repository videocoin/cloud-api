package validator

import (
	"reflect"

	"github.com/VideoCoin/cloud-api/rpc"
	"github.com/VideoCoin/cloud-api/users/v1"
	validatorcore "gopkg.in/go-playground/validator.v9"
)

func ValidateCreateUserRequest(r *v1.CreateUserRequest) *rpc.MultiValidationError {
	trans := *translator

	verrs := &rpc.MultiValidationError{}

	serr := validate.Struct(r)
	if serr != nil {
		verrs.Errors = []*rpc.ValidationError{}

		for _, err := range serr.(validatorcore.ValidationErrors) {
			field, _ := reflect.TypeOf(r).Elem().FieldByName(err.Field())
			jsonField := extractValueFromTag(field.Tag.Get("json"))
			verr := &rpc.ValidationError{
				Field:   jsonField,
				Message: err.Translate(trans),
			}
			verrs.Errors = append(verrs.Errors, verr)
		}

		return verrs
	}

	return nil
}

func ValidateLoginUserRequest(r *v1.LoginUserRequest) *rpc.MultiValidationError {
	trans := *translator

	verrs := &rpc.MultiValidationError{}

	serr := validate.Struct(r)
	if serr != nil {
		verrs.Errors = []*rpc.ValidationError{}

		for _, err := range serr.(validatorcore.ValidationErrors) {
			field, _ := reflect.TypeOf(r).Elem().FieldByName(err.Field())
			jsonField := extractValueFromTag(field.Tag.Get("json"))
			verr := &rpc.ValidationError{
				Field:   jsonField,
				Message: err.Translate(trans),
			}
			verrs.Errors = append(verrs.Errors, verr)
		}

		return verrs
	}

	return nil
}
