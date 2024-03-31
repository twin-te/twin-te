package schoolcalendarport

import (
	"context"

	"cloud.google.com/go/civil"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
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
	DateAfterOrEqual  *civil.Date
	DateBeforeOrEqual *civil.Date
}

// ModuleDetail

type ListModuleDetailsConds struct {
	Year               *shareddomain.AcademicYear
	StartBeforeOrEqual *civil.Date
	EndAfterOrEqual    *civil.Date
}
