package validator

import (
	"reflect"
	"strings"

	"github.com/VideoCoin/cloud-api/rpc"
	"github.com/go-playground/locales"
	enLocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	enTrans "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	validate      *validator.Validate
	uniTranslator *ut.UniversalTranslator
	translator    *ut.Translator
	en            *locales.Translator
)

func init() {
	lt := enLocale.New()
	en = &lt
	uniTranslator = ut.New(*en, *en)
	uniEn, _ := uniTranslator.GetTranslator("en")
	translator = &uniEn

	validate = validator.New()
	enTrans.RegisterDefaultTranslations(validate, *translator)

	validate.RegisterTranslation(
		"email",
		*translator,
		RegisterUserEmailTranslation,
		UserEmailTranslation)
	validate.RegisterTranslation(
		"name",
		*translator,
		RegisterUserNameTranslation,
		UserNameTranslation)
	validate.RegisterTranslation(
		"password",
		*translator,
		RegisterUserPasswordTranslation,
		UserPasswordTranslation)
	validate.RegisterTranslation(
		"secure-password",
		*translator,
		RegisterSecureUserPasswordTranslation,
		SecureUserPasswordTranslation)
	validate.RegisterTranslation(
		"confirm-password",
		*translator,
		RegisterConfirmUserPasswordTranslation,
		ConfirmUserPasswordTranslation)

	validate.RegisterValidation("confirm-password", ValidateConfirmUserPassword)
	validate.RegisterValidation("secure-password", ValidateSecureUserPassword)

	validate.RegisterTranslation(
		"name",
		*translator,
		RegisterStreamNameTranslation,
		StreamNameTranslation)

	translateOverride(uniEn)
}

func translateOverride(trans ut.Translator) {
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "Required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}

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

		for _, err := range serr.(validator.ValidationErrors) {
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
