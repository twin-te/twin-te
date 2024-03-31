package shareddomain

import (
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

func NewRequiredStringParser(name string) func(string) (RequiredString, error) {
	return func(s string) (RequiredString, error) {
		v := strings.TrimSpace(s)
		if v == "" {
			return "", fmt.Errorf("%s must not be empty string", name)
		}
		return RequiredString(v), nil
	}
}
