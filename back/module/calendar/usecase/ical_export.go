package calendarusecase

import (
	"bytes"
	"context"

	"github.com/samber/mo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
)

func filterCoursesByTags(courses []*timetableappdto.RegisteredCourse, tagIDs []idtype.TagID) []*timetableappdto.RegisteredCourse {
	if len(tagIDs) == 0 {
		return courses
	}
	m := make(map[idtype.TagID]struct{}, len(tagIDs))
	for _, t := range tagIDs {
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

func (uc *impl) ExportTimetableToICal(ctx context.Context, year shareddomain.AcademicYear, tagIDs []idtype.TagID, isRdateSupported bool) ([]byte, error) {
	modules, err := uc.buildSchoolCalendarModules(ctx, year)
	if err != nil {
		return nil, err
	}

	courses, err := uc.timetable.ListRegisteredCourses(ctx, mo.Some(year))
	if err != nil {
		return nil, err
	}

	courses = filterCoursesByTags(courses, tagIDs)

	var buf bytes.Buffer
	if err := uc.writeICalendar(&buf, modules, courses, isRdateSupported); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
