package timetabledomain

import (
	"fmt"
	"regexp"
)

// Credit represents course's credit.
// Zero value is invalid.
type Credit string

func (credit Credit) String() string {
	return string(credit)
}

func (credit Credit) IsZero() bool {
	return credit == ""
}

var creditRegexp = regexp.MustCompile(`^\d{1,2}\.[05]$`)

func ParseCredit(s string) (Credit, error) {
	if creditRegexp.MatchString(s) {
		return Credit(s), nil
	}
	return "", fmt.Errorf("failed to parse Credit %v", s)
}
