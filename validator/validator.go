package validator

import (
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
		RegisterEmailTranslation,
		EmailTranslation)
	validate.RegisterTranslation(
		"secure-password",
		*translator,
		RegisterSecurePasswordTranslation,
		SecurePasswordTranslation)
	validate.RegisterTranslation(
		"confirm-password",
		*translator,
		RegisterConfirmPasswordTranslation,
		ConfirmPasswordTranslation)

	validate.RegisterValidation("confirm-password", ValidateConfirmPassword)
	validate.RegisterValidation("secure-password", ValidateSecurePassword)

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
