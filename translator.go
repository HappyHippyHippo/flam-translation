package translation

import (
	ut "github.com/go-playground/universal-translator"
)

type Translator interface {
	ut.Translator
}
