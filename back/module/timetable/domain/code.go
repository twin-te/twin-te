package timetabledomain

import (
	"fmt"
	"regexp"
)

// Code represents course's code.
// Zero value is invalid.
type Code string

func (code Code) String() string {
	return string(code)
}

func (code Code) IsZero() bool {
	return code == ""
}

var codeRegexp = regexp.MustCompile(`^[0-9A-Z]{7}$`)

func ParseCode(s string) (Code, error) {
	if codeRegexp.MatchString(s) {
		return Code(s), nil
	}
	return "", fmt.Errorf("failed to parse Code %#v", s)
}
