package bob

type Builder[T any] struct {
	makeDefault func() T
}

// Build generates a new instance and applies the provided
// override functions from left to right.
func (f Builder[T]) Build(overrides ...func(T) T) T {
	t := f.makeDefault()
	for _, override := range overrides {
		t = override(t)
	}
	return t
}

// Builds generates `n` instances and applies the provided
// override functions from left to right to each of them.
func (f Builder[T]) BuildMany(n int, overrides ...func(int, T) T) []T {
	tt := make([]T, 0, n)
	for i := 0; i < n; i++ {
		t := f.makeDefault()
		for _, override := range overrides {
			t = override(i, t)
		}
		tt = append(tt, t)
	}
	return tt
}

// Override returns a derived Builder obtained by appliyng the provided
// functions to every instance that is generated.
func (f Builder[T]) Override(overrides ...func(T) T) Builder[T] {
	return Builder[T]{
		makeDefault: func() T {
			return f.Build(overrides...)
		},
	}
}

// Returns a new builder
func New[T any](makeDefault func() T) Builder[T] {
	return Builder[T]{makeDefault: makeDefault}
}
