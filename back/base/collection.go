package base

import (
	"fmt"

	"github.com/samber/lo"
)

func CopySlice[T any](collection []T) []T {
	ret := make([]T, len(collection))
	copy(ret, collection)
	return ret
}

func Map[T any, R any](collection []T, iteratee func(item T) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item)
	}

	return result
}

func MapWithArg[T, A, R any](collection []T, arg A, iteratee func(item T, arg A) R) (result []R) {
	result = make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, arg)
	}

	return
}

func MapWithErr[T any, R any](collection []T, iteratee func(item T) (R, error)) (result []R, err error) {
	result = make([]R, len(collection))

	for i, item := range collection {
		if result[i], err = iteratee(item); err != nil {
			return
		}
	}

	return
}

func MapWithArgAndErr[T, A, R any](collection []T, arg A, iteratee func(item T, arg A) (R, error)) (result []R, err error) {
	result = make([]R, len(collection))

	for i, item := range collection {
		if result[i], err = iteratee(item, arg); err != nil {
			return
		}
	}

	return
}

func MapByString[T fmt.Stringer](collectioin []T) []string {
	result := make([]string, len(collectioin))

	for i, item := range collectioin {
		result[i] = item.String()
	}

	return result
}

func MapByClone[T interface{ Clone() T }](collectioin []T) []T {
	result := make([]T, len(collectioin))

	for i, item := range collectioin {
		result[i] = item.Clone()
	}

	return result
}

func ForEachWithErr[T any, R any](collection []T, iteratee func(item T) error) (err error) {
	for _, item := range collection {
		if err = iteratee(item); err != nil {
			return
		}
	}
	return
}

func ForEachWithArgAndErr[T, A any](collection []T, arg A, iteratee func(item T, arg A) error) (err error) {
	for _, item := range collection {
		if err = iteratee(item, arg); err != nil {
			return
		}
	}
	return
}

func FindByString[T fmt.Stringer](collection []T, element string) (T, bool) {
	for _, item := range collection {
		if item.String() == element {
			return item, true
		}
	}

	var result T
	return result, false
}

// HaveSameElements verifies whether two slices have exactly identical elements.
// The order of the elements does not matter.
// Each slices is expected to have no duplicate elements.
func HaveSameElements[T comparable](collectionA, collectionB []T) bool {
	if len(collectionA) != len(collectionB) {
		return false
	}
	intersect := lo.Intersect(collectionA, collectionB)
	return len(intersect) == len(collectionA)
}
