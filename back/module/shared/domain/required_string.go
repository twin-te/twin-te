package shareddomain

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// RequiredString represents non-empty string.
// Zero value is invalid.
type RequiredString string

func (rs RequiredString) String() string {
	return string(rs)
}

func (rs RequiredString) IsZero() bool {
	return rs == ""
}

func (rs *RequiredString) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil
	case string:
		if src == "" {
			return fmt.Errorf("expected non empty string to convert into RequiredString, but got %s", src)
		}
		*rs = RequiredString(src)
		return nil
	default:
		return fmt.Errorf("Scan: unable to scan type %T into RequiredString", src)
	}
}

func (rs RequiredString) Value() (driver.Value, error) {
	return string(rs), nil
}

func NewRequiredStringParser(name string) func(string) (RequiredString, error) {
	return func(s string) (RequiredString, error) {
		v := strings.TrimSpace(s)
		if v == "" {
			return "", fmt.Errorf("%s must not be empty string", name)
		}
		return RequiredString(v), nil
	}
}
