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
	// Event

	ListEvents(ctx context.Context, conds ListEventsConds, lock sharedport.Lock) ([]*schoolcalendardomain.Event, error)
	CreateEvents(ctx context.Context, events ...*schoolcalendardomain.Event) error

	// ModuleDetail

	ListModuleDetails(ctx context.Context, conds ListModuleDetailsConds, lock sharedport.Lock) ([]*schoolcalendardomain.ModuleDetail, error)
	CreateModuleDetails(ctx context.Context, moduleDetails ...*schoolcalendardomain.ModuleDetail) error
}

// Event

type ListEventsConds struct {
	DateAfterOrEqual  mo.Option[civil.Date]
	DateBeforeOrEqual mo.Option[civil.Date]
}

// ModuleDetail

type ListModuleDetailsConds struct {
	Year               mo.Option[shareddomain.AcademicYear]
	StartBeforeOrEqual mo.Option[civil.Date]
	EndAfterOrEqual    mo.Option[civil.Date]
}
