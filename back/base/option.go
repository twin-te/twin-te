package base

import (
	"fmt"

	"github.com/samber/mo"
)

func OptionMap[T, R any](o mo.Option[T], transform func(T) R) mo.Option[R] {
	if v, ok := o.Get(); ok {
		return mo.Some(transform(v))
	}
	return mo.None[R]()
}

func OptionMapWithErr[T, R any](o mo.Option[T], transform func(T) (R, error)) (mo.Option[R], error) {
	value, ok := o.Get()
	if !ok {
		return mo.None[R](), nil
	}

	result, err := transform(value)
	if err != nil {
		return mo.None[R](), err
	}

	return mo.Some(result), nil
}

func OptionMapByString[T fmt.Stringer](o mo.Option[T]) mo.Option[string] {
	if value, ok := o.Get(); ok {
		return mo.Some(value.String())
	}
	return mo.None[string]()
}

func OptionEqualBy[T any](o1, o2 mo.Option[T], comparison func(T, T) bool) bool {
	switch {
	case o1.IsPresent() && o2.IsPresent():
		return comparison(o1.MustGet(), o2.MustGet())
	case o1.IsAbsent() && o2.IsAbsent():
		return true
	default:
		return false
	}
}

func OptionCloneBy[T any](o mo.Option[T], clone func(T) T) mo.Option[T] {
	if value, ok := o.Get(); ok {
		return mo.Some(clone(value))
	}
	return mo.None[T]()
}

func SomeWithErr[T any](value T, err error) (mo.Option[T], error) {
	if err != nil {
		return mo.None[T](), err
	}
	return mo.Some(value), nil
}
