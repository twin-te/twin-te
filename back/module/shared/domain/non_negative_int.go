package shareddomain

import (
	"database/sql/driver"
	"fmt"
)

type NonNegativeInt int

func (nni NonNegativeInt) Int() int {
	return int(nni)
}

func (nni *NonNegativeInt) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil
	case int64:
		if src < 0 {
			return fmt.Errorf("expected non negative int to convert into NonNegativeInt, but got %d", src)
		}
		*nni = NonNegativeInt(src)
		return nil
	default:
		return fmt.Errorf("Scan: unable to scan type %T into NonNegativeInt", src)
	}
}

func (nni NonNegativeInt) Value() (driver.Value, error) {
	return int64(nni), nil
}

func NewNonNegativeIntParser(name string) func(int) (NonNegativeInt, error) {
	return func(i int) (NonNegativeInt, error) {
		if i < 0 {
			return 0, fmt.Errorf("%s must not be negative, but got %d", name, i)
		}
		return NonNegativeInt(i), nil
	}
}
