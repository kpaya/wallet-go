package entity

import (
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt"
)

var Validate *validator.Validate
var Trans ut.Translator

func init() {
	Validate = validator.New()
	pt := pt.New()
	uni := ut.New(pt, pt)
	Trans, _ := uni.GetTranslator("pt")
	pt_translations.RegisterDefaultTranslations(Validate, Trans)
}
