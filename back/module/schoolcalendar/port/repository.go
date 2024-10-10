package schoolcalendarport

import (
	"context"

	"cloud.google.com/go/civil"
	"github.com/samber/mo"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

type Repository interface {
	ListEvents(ctx context.Context, filter EventFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*schoolcalendardomain.Event, error)
	CreateEvents(ctx context.Context, events ...*schoolcalendardomain.Event) error

	ListModuleDetails(ctx context.Context, filter ModuleDetailsFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*schoolcalendardomain.ModuleDetail, error)
	CreateModuleDetails(ctx context.Context, moduleDetails ...*schoolcalendardomain.ModuleDetail) error
}

type EventFilter struct {
	DateAfterOrEqual  mo.Option[civil.Date]
	DateBeforeOrEqual mo.Option[civil.Date]
}

type ModuleDetailsFilter struct {
	Year               mo.Option[shareddomain.AcademicYear]
	StartBeforeOrEqual mo.Option[civil.Date]
	EndAfterOrEqual    mo.Option[civil.Date]
}
