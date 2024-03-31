package schoolcalendarusecase

import (
	"context"
	"fmt"

	"cloud.google.com/go/civil"
	"github.com/twin-te/twinte-back/apperr"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
	schoolcalendarerr "github.com/twin-te/twinte-back/module/schoolcalendar/err"
	schoolcalendarport "github.com/twin-te/twinte-back/module/schoolcalendar/port"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

func (uc *impl) GetModuleDetails(ctx context.Context, year shareddomain.AcademicYear) ([]*schoolcalendardomain.ModuleDetail, error) {
	return uc.r.ListModuleDetails(ctx, schoolcalendarport.ListModuleDetailsConds{
		Year: &year,
	}, sharedport.LockNone)
}

func (uc *impl) GetModuleByDate(ctx context.Context, date civil.Date) (schoolcalendardomain.Module, error) {
	moduleDetails, err := uc.r.ListModuleDetails(ctx, schoolcalendarport.ListModuleDetailsConds{
		StartBeforeOrEqual: &date,
		EndAfterOrEqual:    &date,
	}, sharedport.LockNone)
	if err != nil {
		return 0, err
	}

	if len(moduleDetails) == 0 {
		return 0, apperr.New(schoolcalendarerr.CodeModuleNotFound, fmt.Sprintf("not found module corresponding to the date %s", date))
	}

	return moduleDetails[0].Module, nil
}
