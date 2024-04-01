package timetablev1conv

import (
	timetablev1 "github.com/twin-te/twinte-back/handler/api/rpcgen/timetable/v1"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

func FromPBSchedule(pbSchedule *timetablev1.Schedule) (timetabledomain.Schedule, error) {
	return timetabledomain.ConstructSchedule(func() (schedule timetabledomain.Schedule, err error) {
		schedule.Module, err = FromPBModule(pbSchedule.Module)
		if err != nil {
			return
		}

		schedule.Day, err = FromPBDay(pbSchedule.Day)
		if err != nil {
			return
		}

		if schedule.Day.IsNormal() {
			schedule.Period, err = timetabledomain.ParsePeriod(pbSchedule.Period)
			if err != nil {
				return
			}
		}

		schedule.Rooms = pbSchedule.Rooms

		return
	})
}

func ToPBSchedule(schedule timetabledomain.Schedule) (*timetablev1.Schedule, error) {
	pbModule, err := ToPBModule(schedule.Module)
	if err != nil {
		return nil, err
	}

	pbDay, err := ToPBDay(schedule.Day)
	if err != nil {
		return nil, err
	}

	pbSchedule := &timetablev1.Schedule{
		Module: pbModule,
		Day:    pbDay,
		Period: int32(schedule.Period.Int()),
		Rooms:  schedule.Rooms,
	}

	return pbSchedule, nil
}
