package timetablev1conv

import (
	"fmt"

	timetablev1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/timetable/v1"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func FromPBDay(pbDay timetablev1.Day) (timetabledomain.Day, error) {
	switch pbDay {
	case timetablev1.Day_DAY_SUN:
		return timetabledomain.DaySun, nil
	case timetablev1.Day_DAY_MON:
		return timetabledomain.DayMon, nil
	case timetablev1.Day_DAY_TUE:
		return timetabledomain.DayTue, nil
	case timetablev1.Day_DAY_WED:
		return timetabledomain.DayWed, nil
	case timetablev1.Day_DAY_THU:
		return timetabledomain.DayThu, nil
	case timetablev1.Day_DAY_FRI:
		return timetabledomain.DayFri, nil
	case timetablev1.Day_DAY_SAT:
		return timetabledomain.DaySat, nil
	case timetablev1.Day_DAY_INTENSIVE:
		return timetabledomain.DayIntensive, nil
	case timetablev1.Day_DAY_APPOINTMENT:
		return timetabledomain.DayAppointment, nil
	case timetablev1.Day_DAY_ANY_TIME:
		return timetabledomain.DayAnyTime, nil
	case timetablev1.Day_DAY_NT:
		return timetabledomain.DayNT, nil
	}
	return 0, fmt.Errorf("invalid %#v", pbDay)
}

func ToPBDay(day timetabledomain.Day) (timetablev1.Day, error) {
	switch day {
	case timetabledomain.DaySun:
		return timetablev1.Day_DAY_SUN, nil
	case timetabledomain.DayMon:
		return timetablev1.Day_DAY_MON, nil
	case timetabledomain.DayTue:
		return timetablev1.Day_DAY_TUE, nil
	case timetabledomain.DayWed:
		return timetablev1.Day_DAY_WED, nil
	case timetabledomain.DayThu:
		return timetablev1.Day_DAY_THU, nil
	case timetabledomain.DayFri:
		return timetablev1.Day_DAY_FRI, nil
	case timetabledomain.DaySat:
		return timetablev1.Day_DAY_SAT, nil
	case timetabledomain.DayIntensive:
		return timetablev1.Day_DAY_INTENSIVE, nil
	case timetabledomain.DayAppointment:
		return timetablev1.Day_DAY_APPOINTMENT, nil
	case timetabledomain.DayAnyTime:
		return timetablev1.Day_DAY_ANY_TIME, nil
	case timetabledomain.DayNT:
		return timetablev1.Day_DAY_NT, nil
	}
	return 0, fmt.Errorf("invalid %#v", day)
}
