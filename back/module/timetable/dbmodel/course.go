package timetabledbmodel

import (
	"time"

	"github.com/twin-te/twin-te/back/base"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

type Course struct {
	ID                string
	Year              int
	Code              string
	Name              string
	Instructors       string
	Credit            string
	Overview          string
	Remarks           string
	LastUpdatedAt     time.Time
	HasParseError     bool
	IsAnnual          bool
	RecommendedGrades []CourseRecommendedGrade
	Methods           []CourseMethod
	Schedules         []CourseSchedule

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CourseMethod struct {
	CourseID string
	Method   string
}

type CourseRecommendedGrade struct {
	CourseID         string
	RecommendedGrade int
}

type CourseSchedule struct {
	CourseID  string
	Module    string
	Day       string
	Period    int
	Locations string
}

func FromDBCourse(dbCourse *Course) (*timetabledomain.Course, error) {
	return timetabledomain.ConstructCourse(func(c *timetabledomain.Course) (err error) {
		c.ID, err = idtype.ParseCourseID(dbCourse.ID)
		if err != nil {
			return err
		}

		c.Year, err = shareddomain.ParseAcademicYear(dbCourse.Year)
		if err != nil {
			return err
		}

		c.Code, err = timetabledomain.ParseCode(dbCourse.Code)
		if err != nil {
			return err
		}

		c.Name, err = timetabledomain.ParseName(dbCourse.Name)
		if err != nil {
			return err
		}

		c.Instructors = dbCourse.Instructors

		c.Credit, err = timetabledomain.ParseCredit(dbCourse.Credit)
		if err != nil {
			return err
		}

		c.Overview = dbCourse.Overview
		c.Remarks = dbCourse.Remarks
		c.LastUpdatedAt = dbCourse.LastUpdatedAt
		c.HasParseError = dbCourse.HasParseError
		c.IsAnnual = dbCourse.IsAnnual

		if c.RecommendedGrades, err = base.MapWithErr(dbCourse.RecommendedGrades, FromDBRecommendedGrade); err != nil {
			return err
		}

		if c.Methods, err = base.MapWithErr(dbCourse.Methods, FromDBCourseMethod); err != nil {
			return err
		}

		if c.Schedules, err = base.MapWithErr(dbCourse.Schedules, FromDBCourseSchedule); err != nil {
			return err
		}

		return nil
	})
}

func ToDBCourse(course *timetabledomain.Course, withAssociations bool) *Course {
	dbCourse := &Course{
		ID:            course.ID.String(),
		Year:          course.Year.Int(),
		Code:          course.Code.String(),
		Name:          course.Name.String(),
		Instructors:   course.Instructors,
		Credit:        course.Credit.String(),
		Overview:      course.Overview,
		Remarks:       course.Remarks,
		LastUpdatedAt: course.LastUpdatedAt,
		HasParseError: course.HasParseError,
		IsAnnual:      course.IsAnnual,
	}

	if withAssociations {
		dbCourse.RecommendedGrades = base.MapWithArg(course.RecommendedGrades, course.ID, ToDBRecommendedGrade)
		dbCourse.Methods = base.MapWithArg(course.Methods, course.ID, ToDBCourseMethod)
		dbCourse.Schedules = base.MapWithArg(course.Schedules, course.ID, ToDBCourseSchedule)
	}

	return dbCourse
}

func FromDBCourseMethod(method CourseMethod) (timetabledomain.CourseMethod, error) {
	return timetabledomain.ParseCourseMethod(method.Method)
}

func ToDBCourseMethod(method timetabledomain.CourseMethod, courseID idtype.CourseID) CourseMethod {
	return CourseMethod{
		CourseID: courseID.String(),
		Method:   method.String(),
	}
}

func FromDBRecommendedGrade(dbRecommendedGrade CourseRecommendedGrade) (timetabledomain.RecommendedGrade, error) {
	return timetabledomain.ParseRecommendedGrade(dbRecommendedGrade.RecommendedGrade)
}

func ToDBRecommendedGrade(recommendedGrade timetabledomain.RecommendedGrade, courseID idtype.CourseID) CourseRecommendedGrade {
	return CourseRecommendedGrade{
		CourseID:         courseID.String(),
		RecommendedGrade: recommendedGrade.Int(),
	}
}

func FromDBCourseSchedule(dbCourseSchedule CourseSchedule) (timetabledomain.Schedule, error) {
	return timetabledomain.ParseSchedule(
		dbCourseSchedule.Module,
		dbCourseSchedule.Day,
		dbCourseSchedule.Period,
		dbCourseSchedule.Locations,
	)
}

func ToDBCourseSchedule(schedule timetabledomain.Schedule, courseID idtype.CourseID) CourseSchedule {
	return CourseSchedule{
		CourseID:  courseID.String(),
		Module:    schedule.Module.String(),
		Day:       schedule.Day.String(),
		Period:    schedule.Period.Int(),
		Locations: schedule.Locations,
	}
}
