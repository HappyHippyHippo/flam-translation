package translation

import (
	"go.uber.org/dig"

	"github.com/happyhippyhippo/flam"
)

type provider struct{}

func NewProvider() flam.Provider {
	return &provider{}
}

func (provider) Id() string {
	return providerId
}

func (provider) Register(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	var e error
	provide := func(constructor any, opts ...dig.ProvideOption) bool {
		e = container.Provide(constructor, opts...)
		return e == nil
	}

	_ = provide(newEnglishTranslatorCreator, dig.Group(TranslatorCreatorGroup)) &&
		provide(newTranslatorFactory) &&
		provide(newFacade)

	return e
}

func (provider) Close(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	return container.Invoke(func(
		translatorFactory translatorFactory,
	) error {
		if e := translatorFactory.Close(); e != nil {
			return e
		}

		return nil
	})
}
