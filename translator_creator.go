package translation

import (
	flam "github.com/happyhippyhippo/flam"
)

type TranslatorCreator interface {
	Accept(config flam.Bag) bool
	Create(config flam.Bag) (Translator, error)
}
