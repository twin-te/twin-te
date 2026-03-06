package calendardomain

import (
	"sort"
	"time"

	"cloud.google.com/go/civil"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	unifieddomain "github.com/twin-te/twin-te/back/module/unified/domain"
)

type Schedule struct {
	StartTime  civil.DateTime
	EndTime    civil.DateTime
	Weekday    time.Weekday
	Until      civil.DateTime
	Exceptions []civil.DateTime
	Additions  []civil.DateTime
	Location   string
}

func nextWeekday(date civil.Date, weekday time.Weekday) civil.Date {
	current := date
	for {
		if current.In(JST).Weekday() == weekday {
			return current
		}
		current = current.AddDays(1)
	}
}

func newCivilDateTime(date civil.Date, t civil.Time) civil.DateTime {
	return civil.DateTime{Date: date, Time: t}
}

func GetSchedules(modules []*SchoolCalendarModule, ss []timetabledomain.Schedule) []Schedule {
	type item struct {
		moduleStart timetabledomain.Module
		moduleEnd   timetabledomain.Module
		day         time.Weekday
		periodStart timetabledomain.Period
		periodEnd   timetabledomain.Period
		location    string
	}

	items := make([]item, 0, len(ss))
	for _, s := range ss {
		if s.Day.IsSpecial() || s.Period < 1 || s.Period > 6 {
			continue
		}
		items = append(items, item{
			moduleStart: s.Module,
			moduleEnd:   s.Module,
			day:         s.Day.Weekday(),
			periodStart: s.Period,
			periodEnd:   s.Period,
			location:    s.Locations,
		})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].moduleStart != items[j].moduleStart {
			return items[i].moduleStart < items[j].moduleStart
		}
		if items[i].day != items[j].day {
			return items[i].day < items[j].day
		}
		return items[i].periodStart < items[j].periodStart
	})

	removed := make(map[int]struct{}, len(items))

	for i := range items {
		if _, ok := removed[i]; ok {
			continue
		}
		item := &items[i]
		for j := i + 1; j < len(items); j++ {
			v := items[j]
			if item.day != v.day || item.moduleStart != v.moduleStart || item.location != v.location || item.periodEnd+1 != v.periodStart {
				continue
			}
			item.periodEnd = v.periodEnd
			removed[j] = struct{}{}
		}
	}
	for i := range items {
		if _, ok := removed[i]; ok {
			continue
		}
		item := &items[i]
		for j := i + 1; j < len(items); j++ {
			v := items[j]
			if item.day != v.day || item.periodStart != v.periodStart || item.periodEnd != v.periodEnd || item.location != v.location || item.moduleEnd+1 != v.moduleStart {
				continue
			}
			item.moduleEnd = v.moduleEnd
			removed[j] = struct{}{}
		}
	}

	result := make([]Schedule, 0, len(items)-len(removed))
	for i, item := range items {
		if _, ok := removed[i]; ok {
			continue
		}

		var startTime, endTime civil.DateTime
		var exceptions, additions []civil.DateTime
		var until civil.DateTime

		for idx, m := range modules {
			if m.Module == unifieddomain.TimetableModuleToSchoolCalendarModule[item.moduleStart] {
				date := nextWeekday(m.Start, item.day)
				startTime = newCivilDateTime(date, GetPeriodStart(item.periodStart))
				endTime = newCivilDateTime(date, GetPeriodEnd(item.periodEnd))
			}
			if startTime.IsZero() {
				continue
			}
			if m.Module > unifieddomain.TimetableModuleToSchoolCalendarModule[item.moduleStart] {
				if idx == 0 {
					continue
				}
				d := modules[idx-1].End.AddDays(1)
				for d.Before(m.Start) {
					if d.In(JST).Weekday() == item.day {
						exceptions = append(exceptions, newCivilDateTime(d, GetPeriodStart(item.periodStart)))
					}
					d = d.AddDays(1)
				}
			}
			for _, d := range m.Exceptions[item.day] {
				exceptions = append(exceptions, newCivilDateTime(d, GetPeriodStart(item.periodStart)))
			}
			for _, d := range m.Additions[item.day] {
				additions = append(additions, newCivilDateTime(d, GetPeriodStart(item.periodStart)))
			}
			if m.Module == unifieddomain.TimetableModuleToSchoolCalendarModule[item.moduleEnd] {
				until = newCivilDateTime(m.End, civil.Time{Hour: 23, Minute: 59})
				break
			}
		}

		if startTime.IsZero() {
			continue
		}

		result = append(result, Schedule{
			StartTime:  startTime,
			EndTime:    endTime,
			Weekday:    item.day,
			Until:      until,
			Exceptions: exceptions,
			Additions:  additions,
			Location:   item.location,
		})
	}
	return result
}
