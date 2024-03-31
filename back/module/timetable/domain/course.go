package timetabledomain

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

// Course is identified by one of the following fields.
//   - ID
//   - Year and Code
type Course struct {
	ID                idtype.CourseID
	Year              shareddomain.AcademicYear
	Code              Code
	Name              shareddomain.RequiredString
	Instructors       string
	Credit            Credit
	Overview          string
	Remarks           string
	LastUpdatedAt     time.Time
	HasParseError     bool
	IsAnnual          bool
	RecommendedGrades []RecommendedGrade
	Methods           []CourseMethod
	Schedules         []Schedule

	EntityBeforeUpdated *Course
}

func (c *Course) Clone() *Course {
	ret := lo.ToPtr(*c)
	ret.RecommendedGrades = base.CopySlice(c.RecommendedGrades)
	ret.Methods = base.CopySlice(c.Methods)
	ret.Schedules = base.CopySlice(c.Schedules)
	return ret
}

func (c *Course) BeforeUpdateHook() {
	c.EntityBeforeUpdated = c.Clone()
}

type CourseDataToUpdate struct {
	Name              *shareddomain.RequiredString
	Instructors       *string
	Credit            *Credit
	Overview          *string
	Remarks           *string
	LastUpdatedAt     *time.Time
	HasParseError     *bool
	IsAnnual          *bool
	RecommendedGrades *[]RecommendedGrade
	Methods           *[]CourseMethod
	Schedules         *[]Schedule
}

func (c *Course) Update(data CourseDataToUpdate) {
	if data.Name != nil {
		c.Name = *data.Name
	}

	if data.Instructors != nil {
		c.Instructors = *data.Instructors
	}

	if data.Credit != nil {
		c.Credit = *data.Credit
	}

	if data.Overview != nil {
		c.Overview = *data.Overview
	}

	if data.Remarks != nil {
		c.Remarks = *data.Remarks
	}

	if data.LastUpdatedAt != nil {
		c.LastUpdatedAt = *data.LastUpdatedAt
	}

	if data.HasParseError != nil {
		c.HasParseError = *data.HasParseError
	}

	if data.IsAnnual != nil {
		c.IsAnnual = *data.IsAnnual
	}

	if data.RecommendedGrades != nil {
		c.RecommendedGrades = *data.RecommendedGrades
	}

	if data.Methods != nil {
		c.Methods = *data.Methods
	}

	if data.Schedules != nil {
		c.Schedules = *data.Schedules
	}
}

func ConstructCourse(fn func(c *Course) (err error)) (*Course, error) {
	c := new(Course)
	if err := fn(c); err != nil {
		return nil, err
	}

	if c.ID.IsZero() ||
		c.Year.IsZero() ||
		c.Code.IsZero() ||
		c.Name.IsZero() ||
		c.Credit.IsZero() ||
		c.LastUpdatedAt.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", c)
	}

	return c, nil
}
