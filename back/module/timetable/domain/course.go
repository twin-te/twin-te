package timetabledomain

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
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

	BeforeUpdated mo.Option[*Course]
}

func (c *Course) Clone() *Course {
	ret := lo.ToPtr(*c)
	ret.RecommendedGrades = base.CopySlice(c.RecommendedGrades)
	ret.Methods = base.CopySlice(c.Methods)
	ret.Schedules = base.CopySlice(c.Schedules)
	return ret
}

func (c *Course) BeforeUpdateHook() {
	c.BeforeUpdated = mo.Some(c.Clone())
}

type CourseDataToUpdate struct {
	Name              mo.Option[shareddomain.RequiredString]
	Instructors       mo.Option[string]
	Credit            mo.Option[Credit]
	Overview          mo.Option[string]
	Remarks           mo.Option[string]
	LastUpdatedAt     mo.Option[time.Time]
	HasParseError     mo.Option[bool]
	IsAnnual          mo.Option[bool]
	RecommendedGrades mo.Option[[]RecommendedGrade]
	Methods           mo.Option[[]CourseMethod]
	Schedules         mo.Option[[]Schedule]
}

func (c *Course) Update(data CourseDataToUpdate) {
	if name, ok := data.Name.Get(); ok {
		c.Name = name
	}

	if instructors, ok := data.Instructors.Get(); ok {
		c.Instructors = instructors
	}

	if credit, ok := data.Credit.Get(); ok {
		c.Credit = credit
	}

	if overview, ok := data.Overview.Get(); ok {
		c.Overview = overview
	}

	if remarks, ok := data.Remarks.Get(); ok {
		c.Remarks = remarks
	}

	if lastUpdatedAt, ok := data.LastUpdatedAt.Get(); ok {
		c.LastUpdatedAt = lastUpdatedAt
	}

	if hasParseError, ok := data.HasParseError.Get(); ok {
		c.HasParseError = hasParseError
	}

	if isAnnual, ok := data.IsAnnual.Get(); ok {
		c.IsAnnual = isAnnual
	}

	if recommendedGrades, ok := data.RecommendedGrades.Get(); ok {
		c.RecommendedGrades = recommendedGrades
	}

	if methods, ok := data.Methods.Get(); ok {
		c.Methods = methods
	}

	if schedules, ok := data.Schedules.Get(); ok {
		c.Schedules = schedules
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
