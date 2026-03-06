package calendarusecase

import (
	"context"
	"sort"
	"time"

	"cloud.google.com/go/civil"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

func (uc *impl) buildSchoolCalendarModules(ctx context.Context, year shareddomain.AcademicYear) ([]*calendardomain.SchoolCalendarModule, error) {
	modulesMsg, err := uc.schoolcalendar.ListModuleDetails(ctx, year)
	if err != nil {
		return nil, err
	}

	ms := make([]*calendardomain.SchoolCalendarModule, len(modulesMsg))
	for i, m := range modulesMsg {
		ms[i] = &calendardomain.SchoolCalendarModule{
			Module:     m.Module,
			Start:      m.Start,
			End:        m.End,
			Exceptions: make(map[time.Weekday][]civil.Date, 7),
			Additions:  make(map[time.Weekday][]civil.Date, 7),
		}
	}

	eventsMsg, err := uc.schoolcalendar.ListEvents(ctx, year)
	if err != nil {
		return nil, err
	}

	for _, e := range eventsMsg {
		if e.Type == schoolcalendardomain.EventTypeOther {
			continue
		}
		for _, m := range ms {
			if e.Date.Before(m.Start) || e.Date.After(m.End) {
				continue
			}
			if e.Type == schoolcalendardomain.EventTypeExam && m.Module != schoolcalendardomain.ModuleSpringA && m.Module != schoolcalendardomain.ModuleFallA {
				continue
			}
			eventWeekday := e.Date.In(calendardomain.JST).Weekday()
			if changeTo, ok := e.ChangeTo.Get(); ok {
				if changeTo != eventWeekday {
					m.AddException(eventWeekday, e.Date)
					m.Additions[changeTo] = append(m.Additions[changeTo], e.Date)
				}
			} else {
				m.AddException(eventWeekday, e.Date)
			}
		}
	}

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Start.Before(ms[j].Start)
	})

	return ms, nil
}
