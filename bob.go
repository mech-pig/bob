package bob

type Factory[T any] struct {
	makeDefault func() T
}

func (f Factory[T]) Build(overrides ...func(T) T) T {
	t := f.makeDefault()
	for _, override := range overrides {
		t = override(t)
	}
	return t
}

func (f Factory[T]) BuildMany(n int, overrides ...func(int, T) T) []T {
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

func (f Factory[T]) Override(overrides ...func(T) T) Factory[T] {
	return Factory[T]{
		makeDefault: func() T {
			return f.Build(overrides...)
		},
	}
}

func New[T any](makeDefault func() T) Factory[T] {
	return Factory[T]{makeDefault: makeDefault}
}
