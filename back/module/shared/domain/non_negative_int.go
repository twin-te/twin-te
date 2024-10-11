package shareddomain

import (
	"fmt"
)

type NonNegativeInt int

func (nni NonNegativeInt) Int() int {
	return int(nni)
}

func NewNonNegativeIntParser(name string) func(int) (NonNegativeInt, error) {
	return func(i int) (NonNegativeInt, error) {
		if i < 0 {
			return 0, fmt.Errorf("%s must not be negative, but got %d", name, i)
		}
		return NonNegativeInt(i), nil
	}
}
