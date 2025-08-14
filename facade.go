package translation

type Facade interface {
	HasTranslator(id string) bool
	ListTranslators() []string
	GetTranslator(id string) (Translator, error)
	AddTranslator(id string, translator Translator) error
}

type facade struct {
	translatorFactory translatorFactory
}

func newFacade(
	translatorFactory translatorFactory,
) Facade {
	return &facade{
		translatorFactory: translatorFactory,
	}
}

func (facade facade) HasTranslator(
	id string,
) bool {
	return facade.translatorFactory.Has(id)
}

func (facade facade) ListTranslators() []string {
	return facade.translatorFactory.List()
}

func (facade facade) GetTranslator(
	id string,
) (Translator, error) {
	return facade.translatorFactory.Get(id)
}

func (facade facade) AddTranslator(
	id string,
	translator Translator,
) error {
	return facade.translatorFactory.Add(id, translator)
}
