package shareddomain

import "github.com/samber/mo"

type Association[T any] struct {
	v mo.Option[T]
}

func (a *Association[T]) IsPresent() bool {
	return a.v.IsPresent()
}

func (a *Association[T]) IsAbsent() bool {
	return a.v.IsAbsent()
}

func (a *Association[T]) Get() (T, bool) {
	return a.v.Get()
}

func (a *Association[T]) MustGet() T {
	return a.v.MustGet()
}

func (a *Association[T]) Set(v T) {
	a.v = mo.Some(v)
}
