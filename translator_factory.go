package translation

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
)

type translatorFactory flam.Factory[Translator]

type translatorFactoryArgs struct {
	dig.In

	Creators      []TranslatorCreator `group:"flam.translation.translators.creator"`
	FactoryConfig flam.FactoryConfig
}

func newTranslatorFactory(
	args translatorFactoryArgs,
) (translatorFactory, error) {
	var creators []flam.ResourceCreator[Translator]
	for _, creator := range args.Creators {
		creators = append(creators, creator)
	}

	return flam.NewFactory(
		creators,
		PathTranslators,
		args.FactoryConfig,
		nil)
}
