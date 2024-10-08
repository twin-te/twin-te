package schoolcalendarusecase

import (
	"context"
	"time"

	"cloud.google.com/go/civil"
	"github.com/samber/mo"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	schoolcalendarport "github.com/twin-te/twin-te/back/module/schoolcalendar/port"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (uc *impl) ListEvents(ctx context.Context, year shareddomain.AcademicYear) ([]*schoolcalendardomain.Event, error) {
	return uc.r.ListEvents(ctx, schoolcalendarport.ListEventsConds{
		DateAfterOrEqual: mo.Some(civil.Date{
			Year:  year.Int(),
			Month: time.April,
			Day:   1,
		}),
		DateBeforeOrEqual: mo.Some(civil.Date{
			Year:  year.Int() + 1,
			Month: time.March,
			Day:   31,
		}),
	}, sharedport.LockNone)
}

func (uc *impl) ListEventsByDate(ctx context.Context, date civil.Date) ([]*schoolcalendardomain.Event, error) {
	return uc.r.ListEvents(ctx, schoolcalendarport.ListEventsConds{
		DateAfterOrEqual:  mo.Some(date),
		DateBeforeOrEqual: mo.Some(date),
	}, sharedport.LockNone)
}
