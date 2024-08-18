package base

import "golang.org/x/exp/slices"

func ToPtrWithErr[T any](x T, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	return &x, nil
}

func ToDoublePtrWithErr[T any](x T, err error) (**T, error) {
	return ToPtrWithErr(ToPtrWithErr(x, err))
}

func EqualPtr[T comparable](p1, p2 *T) bool {
	if p1 == nil && p2 == nil {
		return true
	}

	if p1 == nil || p2 == nil {
		return false
	}

	return *p1 == *p2
}

func EqualPtrBy[T any](p1, p2 *T, fn func(v1, v2 T) bool) bool {
	if p1 == nil && p2 == nil {
		return true
	}

	if p1 == nil || p2 == nil {
		return false
	}

	return fn(*p1, *p2)
}

func EqualSlicePtr[T comparable](p1, p2 *[]T) bool {
	if p1 == nil && p2 == nil {
		return true
	}

	if p1 == nil || p2 == nil {
		return false
	}

	return slices.Equal(*p1, *p2)
}
