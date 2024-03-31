package schoolcalendarrepository

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
	schoolcalendarport "github.com/twin-te/twinte-back/module/schoolcalendar/port"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

func (r *impl) ListEvents(ctx context.Context, conds schoolcalendarport.ListEventsConds, lock sharedport.Lock) ([]*schoolcalendardomain.Event, error) {
	events := r.events

	if conds.DateAfterOrEqual != nil {
		events = lo.Filter(events, func(event *schoolcalendardomain.Event, _ int) bool {
			return event.Date.After(*conds.DateAfterOrEqual) || event.Date == *conds.DateAfterOrEqual
		})
	}

	if conds.DateBeforeOrEqual != nil {
		events = lo.Filter(events, func(event *schoolcalendardomain.Event, _ int) bool {
			return event.Date.Before(*conds.DateBeforeOrEqual) || event.Date == *conds.DateBeforeOrEqual
		})
	}

	return base.MapByClone(events), nil
}

func (r *impl) CreateEvents(ctx context.Context, events ...*schoolcalendardomain.Event) error {
	ids := base.Map(events, func(event *schoolcalendardomain.Event) idtype.EventID {
		return event.ID
	})

	savedIDs := base.Map(r.events, func(event *schoolcalendardomain.Event) idtype.EventID {
		return event.ID
	})

	intersect := lo.Intersect(ids, savedIDs)
	if len(intersect) != 0 {
		return fmt.Errorf("duplicate ids: %v", intersect)
	}

	r.events = append(r.events, events...)

	return nil
}
