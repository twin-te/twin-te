package timetabledbmodel

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

type RegisteredCourse struct {
	ID          string
	UserID      string
	Year        int
	CourseID    mo.Option[string]
	Name        mo.Option[string]
	Instructors mo.Option[string]
	Credit      mo.Option[string]
	Methods     mo.Option[string]
	Schedules   sql.Null[string] // error will be occurred if mo.Option is used
	Memo        string
	Attendance  int
	Absence     int
	Late        int
	Tags        []RegisteredCourseTag

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RegisteredCourseTag struct {
	RegisteredCourseID string
	TagID              string
}

func FromDBRegisteredCourse(dbRegisteredCourse *RegisteredCourse) (*timetabledomain.RegisteredCourse, error) {
	return timetabledomain.ConstructRegisteredCourse(func(registeredCourse *timetabledomain.RegisteredCourse) (err error) {
		registeredCourse.ID, err = idtype.ParseRegisteredCourseID(dbRegisteredCourse.ID)
		if err != nil {
			return err
		}

		registeredCourse.UserID, err = idtype.ParseUserID(dbRegisteredCourse.UserID)
		if err != nil {
			return err
		}

		registeredCourse.Year, err = shareddomain.ParseAcademicYear(dbRegisteredCourse.Year)
		if err != nil {
			return err
		}

		registeredCourse.CourseID, err = base.OptionMapWithErr(dbRegisteredCourse.CourseID, idtype.ParseCourseID)
		if err != nil {
			return err
		}

		registeredCourse.Name, err = base.OptionMapWithErr(dbRegisteredCourse.Name, timetabledomain.ParseName)
		if err != nil {
			return err
		}

		registeredCourse.Instructors = dbRegisteredCourse.Instructors

		registeredCourse.Credit, err = base.OptionMapWithErr(dbRegisteredCourse.Credit, timetabledomain.ParseCredit)
		if err != nil {
			return err
		}

		registeredCourse.Methods, err = base.OptionMapWithErr(dbRegisteredCourse.Methods, FromDBRegisteredCourseMethods)
		if err != nil {
			return err
		}

		registeredCourse.Schedules, err = base.OptionMapWithErr(dbhelper.NullToOption(dbRegisteredCourse.Schedules), FromDBRegisteredCourseSchedules)
		if err != nil {
			return err
		}

		registeredCourse.Memo = dbRegisteredCourse.Memo

		registeredCourse.Attendance, err = timetabledomain.ParseAttendance(dbRegisteredCourse.Attendance)
		if err != nil {
			return
		}

		registeredCourse.Absence, err = timetabledomain.ParseAbsence(dbRegisteredCourse.Absence)
		if err != nil {
			return
		}

		registeredCourse.Late, err = timetabledomain.ParseLate(dbRegisteredCourse.Late)
		if err != nil {
			return
		}

		registeredCourse.TagIDs, err = base.MapWithErr(dbRegisteredCourse.Tags, FromDBRegisteredCourseTag)
		if err != nil {
			return err
		}

		return nil
	})
}

func ToDBRegisteredCourse(registeredCourse *timetabledomain.RegisteredCourse, withAssociations bool) (*RegisteredCourse, error) {
	dbRegisteredCourse := &RegisteredCourse{
		ID:          registeredCourse.ID.String(),
		UserID:      registeredCourse.UserID.String(),
		Year:        registeredCourse.Year.Int(),
		CourseID:    base.OptionMapByString(registeredCourse.CourseID),
		Name:        base.OptionMapByString(registeredCourse.Name),
		Instructors: registeredCourse.Instructors,
		Credit:      base.OptionMapByString(registeredCourse.Credit),
		Methods:     base.OptionMap(registeredCourse.Methods, ToDBRegisteredCourseMethods),
		Memo:        registeredCourse.Memo,
		Attendance:  registeredCourse.Attendance.Int(),
		Absence:     registeredCourse.Absence.Int(),
		Late:        registeredCourse.Late.Int(),
	}

	var err error
	stringOption, err := base.OptionMapWithErr(registeredCourse.Schedules, ToDBRegisteredCourseSchedulesJSON)
	if err != nil {
		return nil, err
	}
	dbRegisteredCourse.Schedules = dbhelper.OptionToNull(stringOption)

	if withAssociations {
		dbRegisteredCourse.Tags = base.MapWithArg(registeredCourse.TagIDs, registeredCourse.ID, ToDBRegisteredCourseTag)
	}

	return dbRegisteredCourse, nil
}

type dbRegisteredCourseSchedule struct {
	Module    string `json:"module"`
	Day       string `json:"day"`
	Period    int    `json:"period"`
	Locations string `json:"locations"`
}

func FromDBRegisteredCourseMethods(dbMethods string) ([]timetabledomain.CourseMethod, error) {
	if dbMethods == "{}" {
		return nil, nil
	}

	dbMethods = strings.TrimPrefix(dbMethods, "{")
	dbMethods = strings.TrimSuffix(dbMethods, "}")

	return base.MapWithErr(strings.Split(dbMethods, ","), timetabledomain.ParseCourseMethod)
}

func ToDBRegisteredCourseMethods(methods []timetabledomain.CourseMethod) string {
	return fmt.Sprintf("{%s}", strings.Join(base.MapByString(methods), ","))
}

func FromDBRegisteredCourseSchedules(data string) ([]timetabledomain.Schedule, error) {
	var dbRegisteredCourseSchedules []dbRegisteredCourseSchedule

	if err := json.Unmarshal([]byte(data), &dbRegisteredCourseSchedules); err != nil {
		return nil, err
	}

	return base.MapWithErr(dbRegisteredCourseSchedules, func(dbRegisteredCourseSchedule dbRegisteredCourseSchedule) (timetabledomain.Schedule, error) {
		return timetabledomain.ParseSchedule(
			dbRegisteredCourseSchedule.Module,
			dbRegisteredCourseSchedule.Day,
			dbRegisteredCourseSchedule.Period,
			dbRegisteredCourseSchedule.Locations,
		)
	})
}

func ToDBRegisteredCourseSchedulesJSON(schedules []timetabledomain.Schedule) (string, error) {
	dbRegisteredCourseSchedules := base.Map(schedules, func(schedule timetabledomain.Schedule) *dbRegisteredCourseSchedule {
		return &dbRegisteredCourseSchedule{
			Module:    schedule.Module.String(),
			Day:       schedule.Day.String(),
			Period:    schedule.Period.Int(),
			Locations: schedule.Locations,
		}
	})

	data, err := json.Marshal(dbRegisteredCourseSchedules)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func FromDBRegisteredCourseTag(dbRegisteredCourseTag RegisteredCourseTag) (idtype.TagID, error) {
	return idtype.ParseTagID(dbRegisteredCourseTag.TagID)
}

func ToDBRegisteredCourseTag(tagID idtype.TagID, registeredCourseID idtype.RegisteredCourseID) RegisteredCourseTag {
	return RegisteredCourseTag{
		RegisteredCourseID: registeredCourseID.String(),
		TagID:              tagID.String(),
	}
}
