package unifiedusecase

import (
	"context"
	"time"

	"cloud.google.com/go/civil"
	"github.com/samber/lo"
	"github.com/samber/mo"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	schoolcalendarmodule "github.com/twin-te/twin-te/back/module/schoolcalendar"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	unifiedmodule "github.com/twin-te/twin-te/back/module/unified"
	unifieddomain "github.com/twin-te/twin-te/back/module/unified/domain"
)

var _ unifiedmodule.UseCase = (*impl)(nil)

type impl struct {
	accessController authmodule.AccessController

	schoolCalendarUseCase schoolcalendarmodule.UseCase
	timetableUseCase      timetablemodule.UseCase
}

func (uc *impl) GetByDate(ctx context.Context, date civil.Date) (events []*schoolcalendardomain.Event, module schoolcalendardomain.Module, registeredCourses []*timetableappdto.RegisteredCourse, err error) {
	_, err = uc.accessController.Authenticate(ctx)
	if err != nil {
		return
	}

	events, err = uc.schoolCalendarUseCase.ListEventsByDate(ctx, date)
	if err != nil {
		return
	}

	module, err = uc.schoolCalendarUseCase.GetModuleByDate(ctx, date)
	if err != nil {
		return
	}

	if lo.SomeBy(events, func(event *schoolcalendardomain.Event) bool {
		return lo.Contains([]schoolcalendardomain.EventType{
			schoolcalendardomain.EventTypeHoliday,
			schoolcalendardomain.EventTypePublicHoliday,
		}, event.Type)
	}) {
		return
	}

	if lo.SomeBy(events, func(event *schoolcalendardomain.Event) bool {
		return event.IsSpringAExam() || event.IsSpringCExam() || event.IsFallAExam() || event.IsFallCExam()
	}) {
		return
	}

	if module == schoolcalendardomain.ModuleSpringVacation && date.Month == time.April {
		return
	}

	if module == schoolcalendardomain.ModuleWinterVacation {
		return
	}

	weekday := date.In(time.Local).Weekday()

	for _, event := range events {
		if event.Type == schoolcalendardomain.EventTypeSubstituteDay {
			weekday = event.ChangeTo.MustGet()
		}
	}

	academicYear, err := shareddomain.NewAcademicYearFromDate(date)
	if err != nil {
		return
	}

	registeredCourses, err = uc.timetableUseCase.ListRegisteredCourses(ctx, mo.Some(academicYear))
	if err != nil {
		return
	}

	registeredCourses = lo.Filter(registeredCourses, func(registeredCourse *timetableappdto.RegisteredCourse, index int) bool {
		return lo.SomeBy(registeredCourse.Schedules, func(schedule timetabledomain.Schedule) bool {
			return schedule.IsNormal() && module == unifieddomain.TimetableModuleToSchoolCalendarModule[schedule.Module] && schedule.Day.Weekday() == weekday
		})
	})

	return
}

func New(accessController authmodule.AccessController, schoolCalendarUseCase schoolcalendarmodule.UseCase, timetableUseCase timetablemodule.UseCase) *impl {
	return &impl{
		accessController:      accessController,
		schoolCalendarUseCase: schoolCalendarUseCase,
		timetableUseCase:      timetableUseCase,
	}
}
