package calendarv1beta

import (
	"cloud.google.com/go/civil"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func GetPeriodStart(period timetabledomain.Period) civil.Time {
	switch period {
	case 1:
		return civil.Time{Hour: 8, Minute: 40}
	case 2:
		return civil.Time{Hour: 10, Minute: 10}
	case 3:
		return civil.Time{Hour: 12, Minute: 15}
	case 4:
		return civil.Time{Hour: 13, Minute: 45}
	case 5:
		return civil.Time{Hour: 15, Minute: 15}
	case 6:
		return civil.Time{Hour: 16, Minute: 45}
	default:
		return civil.Time{Hour: 0, Minute: 0}
	}
}

func GetPeriodEnd(period timetabledomain.Period) civil.Time {
	switch period {
	case 1:
		return civil.Time{Hour: 9, Minute: 55}
	case 2:
		return civil.Time{Hour: 11, Minute: 25}
	case 3:
		return civil.Time{Hour: 13, Minute: 30}
	case 4:
		return civil.Time{Hour: 15, Minute: 00}
	case 5:
		return civil.Time{Hour: 16, Minute: 30}
	case 6:
		return civil.Time{Hour: 18, Minute: 00}
	default:
		return civil.Time{Hour: 23, Minute: 59}
	}
}
