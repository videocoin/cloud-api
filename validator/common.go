package validator

import (
	"reflect"
	"strings"

	"github.com/VideoCoin/cloud-api/rpc"
	validatorcore "gopkg.in/go-playground/validator.v9"
)

func extractValueFromTag(tag string) string {
	values := strings.Split(tag, ",")
	return values[0]
}

func ValidateRequest(r interface{}) *rpc.MultiValidationError {
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
