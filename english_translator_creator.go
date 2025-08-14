package translation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"

	flam "github.com/happyhippyhippo/flam"
)

type englishTranslatorCreator struct{}

func newEnglishTranslatorCreator() TranslatorCreator {
	return &englishTranslatorCreator{}
}

func (englishTranslatorCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == TranslatorDriverEnglish
}

func (englishTranslatorCreator) Create(
	_ flam.Bag,
) (Translator, error) {
	lang := en.New()
	translator, found := ut.New(lang, lang).GetTranslator("en")
	if !found {
		return nil, newErrLanguageNotFound("en")
	}

	return Translator(translator), nil
}
