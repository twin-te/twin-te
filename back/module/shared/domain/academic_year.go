package shareddomain

import (
	"fmt"
	"time"

	"cloud.google.com/go/civil"
)

// AcademicYear starts in April and ends in March.
// Zero value is invalid.
type AcademicYear int

func (year AcademicYear) Int() int {
	return int(year)
}

func (year AcademicYear) IsZero() bool {
	return year == 0
}

func ParseAcademicYear(i int) (AcademicYear, error) {
	if i <= 0 {
		return 0, fmt.Errorf("failed to parse AcademicYear %v", i)
	}
	return AcademicYear(i), nil
}

func NewAcademicYear(year int, month time.Month) (AcademicYear, error) {
	if month < time.April {
		year -= 1
	}
	return ParseAcademicYear(year)
}

func NewAcademicYearFromDate(date civil.Date) (AcademicYear, error) {
	return NewAcademicYear(date.Year, date.Month)
}
