package schoolcalendarusecase

import (
	"context"
	"fmt"

	"cloud.google.com/go/civil"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/apperr"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	schoolcalendarerr "github.com/twin-te/twin-te/back/module/schoolcalendar/err"
	schoolcalendarport "github.com/twin-te/twin-te/back/module/schoolcalendar/port"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (uc *impl) ListModuleDetails(ctx context.Context, year shareddomain.AcademicYear) ([]*schoolcalendardomain.ModuleDetail, error) {
	return uc.r.ListModuleDetails(ctx, schoolcalendarport.ModuleDetailsFilter{
		Year: mo.Some(year),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
}

func (uc *impl) GetModuleByDate(ctx context.Context, date civil.Date) (schoolcalendardomain.Module, error) {
	moduleDetails, err := uc.r.ListModuleDetails(ctx, schoolcalendarport.ModuleDetailsFilter{
		StartBeforeOrEqual: mo.Some(date),
		EndAfterOrEqual:    mo.Some(date),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
	if err != nil {
		return 0, err
	}

	if len(moduleDetails) == 0 {
		return 0, apperr.New(schoolcalendarerr.CodeModuleNotFound, fmt.Sprintf("not found module corresponding to the date %s", date))
	}

	return moduleDetails[0].Module, nil
}
