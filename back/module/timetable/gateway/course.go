package timetablegateway

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/twin-te/twin-te/back/base"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
)

func (g *impl) GetCourseWithoutIDsFromKdB(ctx context.Context, year shareddomain.AcademicYear) ([]timetableport.CourseWithoutID, error) {
	kdbCourses, err := loadKdBCourseFromJSON(g.kdbJSONFilePath)
	if err != nil {
		return nil, err
	}

	return base.MapWithArgAndErr(kdbCourses, year, parseKdbCoure)
}

type kdbCourse struct {
	Code              string        `json:"code"`
	Name              string        `json:"name"`
	Instructors       string        `json:"instructors"`
	Credit            string        `json:"credit"`
	Overview          string        `json:"overview"`
	Remarks           string        `json:"remarks"`
	LastUpdatedAt     time.Time     `json:"lastUpdatedAt"`
	HasParseError     bool          `json:"hasParseError"`
	IsAnnual          bool          `json:"isAnnual"`
	RecommendedGrades []int         `json:"recommendedGrades"`
	Methods           []string      `json:"methods"`
	Schedules         []kdbSchedule `json:"schedules"`
}

type kdbSchedule struct {
	Module string `json:"module"`
	Day    string `json:"day"`
	Period int    `json:"period"`
	Rooms  string `json:"rooms"`
}

func loadKdBCourseFromJSON(kdbJsonFilePath string) (ret []*kdbCourse, err error) {
	data, err := os.ReadFile(kdbJsonFilePath)
	if err != nil {
		return
	}
	return ret, json.Unmarshal(data, &ret)
}

func parseKdbCoure(kdbCourse *kdbCourse, year shareddomain.AcademicYear) (courseWithoutID timetableport.CourseWithoutID, err error) {
	courseWithoutID.Year = year

	courseWithoutID.Code, err = timetabledomain.ParseCode(kdbCourse.Code)
	if err != nil {
		return
	}

	courseWithoutID.Name, err = timetabledomain.ParseName(kdbCourse.Name)
	if err != nil {
		return
	}

	courseWithoutID.Instructors = kdbCourse.Instructors

	courseWithoutID.Credit, err = timetabledomain.ParseCredit(kdbCourse.Credit)
	if err != nil {
		return
	}

	courseWithoutID.Overview = kdbCourse.Overview
	courseWithoutID.Remarks = kdbCourse.Remarks
	courseWithoutID.LastUpdatedAt = kdbCourse.LastUpdatedAt
	courseWithoutID.HasParseError = kdbCourse.HasParseError
	courseWithoutID.IsAnnual = kdbCourse.IsAnnual

	courseWithoutID.RecommendedGrades, err = base.MapWithErr(kdbCourse.RecommendedGrades, timetabledomain.ParseRecommendedGrade)
	if err != nil {
		return
	}

	courseWithoutID.Methods, err = base.MapWithErr(kdbCourse.Methods, timetabledomain.ParseCourseMethod)
	if err != nil {
		return
	}

	courseWithoutID.Schedules, err = base.MapWithErr(kdbCourse.Schedules, func(kdbSchedule kdbSchedule) (timetabledomain.Schedule, error) {
		return timetabledomain.ParseSchedule(
			kdbSchedule.Module,
			kdbSchedule.Day,
			kdbSchedule.Period,
			kdbSchedule.Rooms,
		)
	})
	if err != nil {
		return
	}

	return
}
