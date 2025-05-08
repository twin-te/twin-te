package calendarv1beta

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/mo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
)

func GetYear(c echo.Context) (shareddomain.AcademicYear, error) {
	yearParam := c.QueryParam("year")
	if yearParam == "" {
		now := time.Now()
		return shareddomain.NewAcademicYear(now.Year(), now.Month())
	}

	year, err := strconv.Atoi(yearParam)
	if err != nil {
		return 0, err
	}

	return shareddomain.ParseAcademicYear(year)
}

func GetTags(c echo.Context) ([]idtype.TagID, error) {
	ss, ok := c.QueryParams()["tags[]"]
	if !ok {
		return nil, nil
	}

	ids := make([]idtype.TagID, len(ss))
	for i, s := range ss {
		id, err := idtype.ParseTagID(s)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}

	return ids, nil
}

func FilterByTags(courses []*timetableappdto.RegisteredCourse, tags []idtype.TagID) []*timetableappdto.RegisteredCourse {
	m := make(map[idtype.TagID]struct{}, len(tags))
	for _, t := range tags {
		m[t] = struct{}{}
	}
	filtered := make([]*timetableappdto.RegisteredCourse, 0, len(courses))
	for _, c := range courses {
		for _, t := range c.TagIDs {
			if _, ok := m[t]; ok {
				filtered = append(filtered, c)
				break
			}
		}
	}
	return filtered
}

func (h *impl) ICSHandler(c echo.Context) error {
	year, err := GetYear(c)
	if err != nil {
		log.Printf("failed to get academic year: %v", err)
		return err
	}

	rdateParam := c.QueryParam("rdate")
	isRdateSupported := rdateParam == "" || rdateParam == "true"

	ctx := c.Request().Context()
	modules, err := h.GetSchoolCalendar(ctx, year)
	if err != nil {
		log.Printf("failed to get school calendar: %+v", err)
		return err
	}

	courses, err := h.timetable.ListRegisteredCourses(ctx, mo.Some(year))
	if err != nil {
		return err
	}

	tags, err := GetTags(c)
	if err != nil {
		log.Printf("failed to get tags: %+v", err)
		return err
	}

	if tags != nil {
		courses = FilterByTags(courses, tags)
	}

	var resp bytes.Buffer
	err = WriteICalendar(&resp, modules, courses, isRdateSupported)
	if err != nil {
		log.Printf("failed to write iCalendar: %+v", err)
		return err
	}

	return c.Stream(http.StatusOK, "application/octet-stream", &resp)
}
