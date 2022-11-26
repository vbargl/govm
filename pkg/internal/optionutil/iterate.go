package optionutil

func Iterate[T any, Opt ~func(*T)](value *T, options []Opt) {
	for _, opt := range options {
		opt(value)
	}
}
