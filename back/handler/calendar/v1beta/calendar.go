package calendarv1beta

import (
	"sort"
	"time"

	"cloud.google.com/go/civil"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	unifieddomain "github.com/twin-te/twin-te/back/module/unified/domain"
)

type Schedule struct {
	StartTime civil.DateTime
	EndTime   civil.DateTime

	Weekday time.Weekday
	Until   civil.DateTime

	Exceptions []civil.DateTime
	Additions  []civil.DateTime

	Location string
}

func nextWeekday(date civil.Date, weekday time.Weekday) civil.Date {
	current := date
	for {
		if current.In(time.UTC).Weekday() == weekday {
			return current
		}
		current = current.AddDays(1)
	}
}

func newCivilDateTime(date civil.Date, time civil.Time) civil.DateTime {
	return civil.DateTime{Date: date, Time: time}
}

func GetSchedules(modules []Module, ss []timetabledomain.Schedule) []Schedule {
	type item struct {
		ModuleStart timetabledomain.Module
		ModuleEnd   timetabledomain.Module

		Day time.Weekday

		PeriodStart timetabledomain.Period
		PeriodEnd   timetabledomain.Period

		Location string
	}

	items := make([]item, 0, len(ss))
	for _, s := range ss {
		if s.Day.IsSpecial() || s.Period < 1 || s.Period > 6 { // TODO: Support 7~8 period
			continue
		}
		items = append(items, item{
			ModuleStart: s.Module,
			ModuleEnd:   s.Module,
			Day:         s.Day.Weekday(),
			PeriodStart: s.Period,
			PeriodEnd:   s.Period,
			Location:    s.Locations,
		})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].ModuleStart != items[j].ModuleStart {
			return items[i].ModuleStart < items[j].ModuleStart
		}
		return items[i].PeriodStart < items[j].PeriodStart
	})

	removed := make(map[int]struct{}, len(items))

	for i, item := range items {
		if _, ok := removed[i]; ok {
			continue
		}
		for j := i + 1; j < len(items); j++ {
			v := items[j]
			if item.Day != v.Day {
				continue
			}
			if item.ModuleStart != v.ModuleStart {
				continue
			}
			if item.PeriodEnd+1 != v.PeriodStart {
				continue
			}
			item.PeriodEnd = v.PeriodEnd
			removed[j] = struct{}{}
		}
		items[i] = item
	}
	for i, item := range items {
		if _, ok := removed[i]; ok {
			continue
		}
		for j := i + 1; j < len(items); j++ {
			v := items[j]
			if item.Day != v.Day {
				continue
			}
			if item.PeriodStart != v.PeriodStart {
				continue
			}
			if item.PeriodEnd != v.PeriodEnd {
				continue
			}
			if item.ModuleEnd+1 != v.ModuleStart {
				continue
			}
			item.ModuleEnd = v.ModuleEnd
			removed[j] = struct{}{}
		}
		items[i] = item
	}

	result := make([]Schedule, 0, len(items)-len(removed))
	for i, item := range items {
		if _, ok := removed[i]; ok {
			continue
		}

		var startTime civil.DateTime
		var endTime civil.DateTime
		var exceptions []civil.DateTime
		var additions []civil.DateTime
		var until civil.DateTime

		for i, m := range modules {
			if m.Module == unifieddomain.TimetableModuleToSchoolCalendarModule[item.ModuleStart] {
				date := nextWeekday(m.Start, item.Day)
				startTime = newCivilDateTime(date, GetPeriodStart(item.PeriodStart))
				endTime = newCivilDateTime(date, GetPeriodEnd(item.PeriodEnd))
			}
			if startTime.IsZero() {
				continue
			}
			if m.Module > unifieddomain.TimetableModuleToSchoolCalendarModule[item.ModuleStart] {
				d := modules[i-1].End.AddDays(1)
				for d.Before(m.Start) {
					if d.In(jst).Weekday() == item.Day {
						exceptions = append(exceptions, newCivilDateTime(d, GetPeriodStart(item.PeriodStart)))
					}
					d = d.AddDays(1)
				}
			}
			for _, d := range m.Exceptions[item.Day] {
				exceptions = append(exceptions, newCivilDateTime(d, GetPeriodStart(item.PeriodStart)))
			}
			for _, d := range m.Additions[item.Day] {
				additions = append(additions, newCivilDateTime(d, GetPeriodStart(item.PeriodStart)))
			}
			if m.Module == unifieddomain.TimetableModuleToSchoolCalendarModule[item.ModuleEnd] {
				until = newCivilDateTime(m.End, civil.Time{Hour: 23, Minute: 59})
				break
			}
		}

		result = append(result, Schedule{
			StartTime:  startTime,
			EndTime:    endTime,
			Weekday:    item.Day,
			Until:      until,
			Exceptions: exceptions,
			Additions:  additions,
			Location:   item.Location,
		})
	}
	return result
}
