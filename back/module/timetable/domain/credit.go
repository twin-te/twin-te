package timetabledomain

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

// Credit represents course's credit.
// Zero value is invalid.
type Credit string

func (credit Credit) String() string {
	return string(credit)
}

func (credit Credit) Float() float64 {
	return lo.Must(strconv.ParseFloat(credit.String(), 32))
}

func (credit *Credit) FloatPtr() *float64 {
	if credit == nil {
		return nil
	}
	return lo.ToPtr(credit.Float())
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
