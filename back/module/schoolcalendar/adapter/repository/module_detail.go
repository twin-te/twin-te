package schoolcalendarrepository

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	schoolcalendarport "github.com/twin-te/twin-te/back/module/schoolcalendar/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (r *impl) ListModuleDetails(ctx context.Context, filter schoolcalendarport.ModuleDetailsFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*schoolcalendardomain.ModuleDetail, error) {
	moduleDetails := applyModuleDetailFilter(r.moduleDetails, filter)
	return base.MapByClone(moduleDetails), nil
}

func (r *impl) CreateModuleDetails(ctx context.Context, moduleDetails ...*schoolcalendardomain.ModuleDetail) error {
	ids := base.Map(moduleDetails, func(moduleDetail *schoolcalendardomain.ModuleDetail) idtype.ModuleDetailID {
		return moduleDetail.ID
	})

	savedIDs := base.Map(r.moduleDetails, func(moduleDetail *schoolcalendardomain.ModuleDetail) idtype.ModuleDetailID {
		return moduleDetail.ID
	})

	intersect := lo.Intersect(ids, savedIDs)
	if len(intersect) != 0 {
		return fmt.Errorf("duplicate ids: %v", intersect)
	}

	r.moduleDetails = append(r.moduleDetails, moduleDetails...)

	return nil
}

func applyModuleDetailFilter(moduleDetails []*schoolcalendardomain.ModuleDetail, filter schoolcalendarport.ModuleDetailsFilter) []*schoolcalendardomain.ModuleDetail {
	if year, ok := filter.Year.Get(); ok {
		moduleDetails = lo.Filter(moduleDetails, func(moduleDetail *schoolcalendardomain.ModuleDetail, _ int) bool {
			return moduleDetail.Year == year
		})
	}

	if startBeforeOrEqual, ok := filter.StartBeforeOrEqual.Get(); ok {
		moduleDetails = lo.Filter(moduleDetails, func(moduleDetail *schoolcalendardomain.ModuleDetail, _ int) bool {
			return moduleDetail.Start.Before(startBeforeOrEqual) || moduleDetail.Start == startBeforeOrEqual
		})
	}

	if endAfterOrEqual, ok := filter.EndAfterOrEqual.Get(); ok {
		moduleDetails = lo.Filter(moduleDetails, func(moduleDetail *schoolcalendardomain.ModuleDetail, _ int) bool {
			return moduleDetail.End.After(endAfterOrEqual) || moduleDetail.End == endAfterOrEqual
		})
	}

	return moduleDetails
}
