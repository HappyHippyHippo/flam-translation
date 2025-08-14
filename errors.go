package translation

import (
	"errors"

	flam "github.com/happyhippyhippo/flam"
)

var (
	ErrLanguageNotFound = errors.New("translator language not found")
)

func newErrNilReference(
	arg string,
) error {
	return flam.NewErrorFrom(
		flam.ErrNilReference,
		arg)
}

func newErrLanguageNotFound(
	language string,
) error {
	return flam.NewErrorFrom(
		ErrLanguageNotFound,
		language)
}
