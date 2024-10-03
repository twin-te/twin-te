package timetabledomain

import (
	"errors"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type CourseWithoutID struct {
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
}

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

func (c *Course) UpdateFromCourseWithoutID(courseWithoutID CourseWithoutID) error {
	if c.Year != courseWithoutID.Year || c.Code != courseWithoutID.Code {
		return errors.New("invalid course without id")
	}

	c.Name = courseWithoutID.Name
	c.Instructors = courseWithoutID.Instructors
	c.Credit = courseWithoutID.Credit
	c.Overview = courseWithoutID.Overview
	c.Remarks = courseWithoutID.Remarks
	c.LastUpdatedAt = courseWithoutID.LastUpdatedAt
	c.HasParseError = courseWithoutID.HasParseError
	c.IsAnnual = courseWithoutID.IsAnnual
	c.RecommendedGrades = courseWithoutID.RecommendedGrades
	c.Methods = courseWithoutID.Methods
	c.Schedules = courseWithoutID.Schedules

	return nil
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
