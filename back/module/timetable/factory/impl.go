package timetablefactory

import (
	"database/sql"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetabledto "github.com/twin-te/twin-te/back/module/timetable/dto"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	"gorm.io/gorm"
)

var _ timetableport.Factory = (*impl)(nil)

type impl struct {
	db *gorm.DB
}

func (f *impl) NewCourse(courseWithoutID timetabledto.CourseWithoutID) (*timetabledomain.Course, error) {
	return timetabledomain.ConstructCourse(func(c *timetabledomain.Course) error {
		c.ID = idtype.NewCourseID()
		c.Year = courseWithoutID.Year
		c.Code = courseWithoutID.Code
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
	})
}

func (f *impl) NewRegisteredCourseFromCourse(userID idtype.UserID, course *timetabledomain.Course) (*timetabledomain.RegisteredCourse, error) {
	return timetabledomain.ConstructRegisteredCourse(func(rc *timetabledomain.RegisteredCourse) error {
		rc.ID = idtype.NewRegisteredCourseID()
		rc.UserID = userID
		rc.CourseID = &course.ID
		rc.Year = course.Year
		rc.CourseAssociation.Set(course)
		return nil
	})
}

func (f *impl) NewRegisteredCourseMannualy(
	userID idtype.UserID,
	year shareddomain.AcademicYear,
	name shareddomain.RequiredString,
	instructors string,
	credit timetabledomain.Credit,
	methods []timetabledomain.CourseMethod,
	schedules []timetabledomain.Schedule,
) (*timetabledomain.RegisteredCourse, error) {
	return timetabledomain.ConstructRegisteredCourse(func(rc *timetabledomain.RegisteredCourse) error {
		rc.ID = idtype.NewRegisteredCourseID()
		rc.UserID = userID
		rc.Year = year
		rc.Name = &name
		rc.Instructors = &instructors
		rc.Credit = &credit
		rc.Methods = &methods
		rc.Schedules = &schedules
		return nil
	})
}

func (f *impl) NewTag(
	userID idtype.UserID,
	name shareddomain.RequiredString,
) (*timetabledomain.Tag, error) {
	var result sql.NullInt16
	if err := f.db.Raw("SELECT max(position) FROM tags WHERE user_id = ?", userID.String()).
		Scan(&result).
		Error; err != nil {
		return nil, err
	}

	var position shareddomain.NonNegativeInt
	if result.Valid {
		maxPosition, err := timetabledomain.ParsePosition(int(result.Int16))
		if err != nil {
			return nil, err
		}
		position = maxPosition + 1
	}

	return timetabledomain.ConstructTag(func(t *timetabledomain.Tag) error {
		t.ID = idtype.NewTagID()
		t.UserID = userID
		t.Name = name
		t.Position = position
		return nil
	})
}

func New(db *gorm.DB) *impl {
	return &impl{db}
}
