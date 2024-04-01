package sharedconv

import (
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"

	"github.com/twin-te/twinte-back/handler/api/rpcgen/sharedpb"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
)

func FromPBAcadimicYear(pbYear *sharedpb.AcademicYear) (shareddomain.AcademicYear, error) {
	if pbYear == nil {
		return 0, errors.New("academic year must be present")
	}
	return shareddomain.ParseAcademicYear(int(pbYear.Value))
}

func ToPBAcademicYear(year shareddomain.AcademicYear) *sharedpb.AcademicYear {
	return &sharedpb.AcademicYear{Value: int32(year)}
}

func FromPBUUID[T ~[16]byte](pbUUID *sharedpb.UUID, parser func(s string) (T, error)) (T, error) {
	if pbUUID == nil {
		return T{}, errors.New("uuid must be present")
	}
	return parser(pbUUID.Value)
}

func ToPBUUID[T ~[16]byte](id T) *sharedpb.UUID {
	return &sharedpb.UUID{Value: uuid.UUID(id).String()}
}

func FromPBRFC3339DateTime(pbDateTime *sharedpb.RFC3339DateTime) (time.Time, error) {
	if pbDateTime == nil {
		return time.Time{}, errors.New("date time must be present")
	}
	return time.Parse(time.RFC3339Nano, pbDateTime.Value)
}

func ToPBRFC3339DateTime(t time.Time) *sharedpb.RFC3339DateTime {
	return &sharedpb.RFC3339DateTime{Value: t.Format(time.RFC3339Nano)}
}

func FromPBRFC3339FullDate(pbFullDate *sharedpb.RFC3339FullDate) (civil.Date, error) {
	if pbFullDate == nil {
		return civil.Date{}, errors.New("full date must be present")
	}
	return civil.ParseDate(pbFullDate.Value)
}

func ToPBRFC3339FullDate(cd civil.Date) *sharedpb.RFC3339FullDate {
	return &sharedpb.RFC3339FullDate{Value: cd.String()}
}

func FromPBWeekday(pbWeekday sharedpb.Weekday) (time.Weekday, error) {
	switch pbWeekday {
	case sharedpb.Weekday_WEEKDAY_SUNDAY:
		return time.Sunday, nil
	case sharedpb.Weekday_WEEKDAY_MONDAY:
		return time.Monday, nil
	case sharedpb.Weekday_WEEKDAY_TUESDAY:
		return time.Tuesday, nil
	case sharedpb.Weekday_WEEKDAY_WEDNESDAY:
		return time.Wednesday, nil
	case sharedpb.Weekday_WEEKDAY_THURSDAY:
		return time.Thursday, nil
	case sharedpb.Weekday_WEEKDAY_FRIDAY:
		return time.Friday, nil
	case sharedpb.Weekday_WEEKDAY_SATURDAY:
		return time.Saturday, nil
	}
	return 0, fmt.Errorf("invalid %#v", pbWeekday)
}

func ToPBWeekday(weekday time.Weekday) (sharedpb.Weekday, error) {
	switch weekday {
	case time.Sunday:
		return sharedpb.Weekday_WEEKDAY_SUNDAY, nil
	case time.Monday:
		return sharedpb.Weekday_WEEKDAY_MONDAY, nil
	case time.Tuesday:
		return sharedpb.Weekday_WEEKDAY_TUESDAY, nil
	case time.Wednesday:
		return sharedpb.Weekday_WEEKDAY_WEDNESDAY, nil
	case time.Thursday:
		return sharedpb.Weekday_WEEKDAY_THURSDAY, nil
	case time.Friday:
		return sharedpb.Weekday_WEEKDAY_FRIDAY, nil
	case time.Saturday:
		return sharedpb.Weekday_WEEKDAY_SATURDAY, nil
	}
	return 0, fmt.Errorf("invalid %#v", weekday)
}
