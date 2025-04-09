package calendarv1beta

import (
	"context"
	"sort"
	"time"

	"cloud.google.com/go/civil"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

type Module struct {
	Module schoolcalendardomain.Module
	Start  civil.Date
	End    civil.Date

	Exceptions map[time.Weekday][]civil.Date
	Additions  map[time.Weekday][]civil.Date
}

func (m Module) addException(weekday time.Weekday, date civil.Date) {
	for _, d := range m.Exceptions[weekday] {
		if d == date {
			return // Already exists
		}
	}
	m.Exceptions[weekday] = append(m.Exceptions[weekday], date)
}

func (h *impl) GetSchoolCalendar(ctx context.Context, year shareddomain.AcademicYear) ([]Module, error) {
	modulesMsg, err := h.schoolcalendar.ListModuleDetails(ctx, year)
	if err != nil {
		return nil, err
	}

	ms := make([]Module, len(modulesMsg))
	for i, m := range modulesMsg {
		ms[i] = Module{
			Module:     m.Module,
			Start:      m.Start,
			End:        m.End,
			Exceptions: make(map[time.Weekday][]civil.Date, 7),
			Additions:  make(map[time.Weekday][]civil.Date, 7),
		}
	}

	eventsMsg, err := h.schoolcalendar.ListEvents(ctx, year)
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
			eventWeekday := e.Date.In(jst).Weekday()
			if changeTo, ok := e.ChangeTo.Get(); ok {
				if changeTo != eventWeekday { // Do not add exception and addition to same date
					m.addException(eventWeekday, e.Date)
					m.Additions[changeTo] = append(m.Additions[changeTo], e.Date)
				}
			} else {
				m.addException(eventWeekday, e.Date)
			}
		}
	}

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Start.Before(ms[j].Start)
	})

	return ms, nil
}
