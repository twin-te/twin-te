package timetabledomain

import "fmt"

// RecommendedGrade is between 1 and 6.
// Zero value is invalid.
type RecommendedGrade int

func (grade RecommendedGrade) Int() int {
	return int(grade)
}

func (grade RecommendedGrade) IsZero() bool {
	return grade == 0
}

func ParseRecommendedGrade(i int) (RecommendedGrade, error) {
	if 1 <= i && i <= 6 {
		return RecommendedGrade(i), nil
	}
	return 0, fmt.Errorf("failed to parse RecommendedGrade %v", i)
}
