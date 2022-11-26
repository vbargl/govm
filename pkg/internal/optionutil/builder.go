package optionutil

type Builder[T any, Opt ~func(*T)] struct {
	defaultOptions []Opt
	givenOptions   []Opt
}

func (b *Builder[T, Opt]) AddDefaults(options ...Opt) {
	b.defaultOptions = append(b.defaultOptions, options...)
}

func (b *Builder[T, Opt]) AddGiven(options ...Opt) {
	b.givenOptions = append(b.givenOptions, options...)
}

func (b *Builder[T, Opt]) Build(value *T) {
	Iterate(value, b.defaultOptions)
	Iterate(value, b.givenOptions)
}
