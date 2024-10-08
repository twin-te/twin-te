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

func (r *impl) ListEvents(ctx context.Context, filter schoolcalendarport.EventFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*schoolcalendardomain.Event, error) {
	events := applyEventFilter(r.events, filter)
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

func applyEventFilter(events []*schoolcalendardomain.Event, filter schoolcalendarport.EventFilter) []*schoolcalendardomain.Event {
	if dateAfterOrEqual, ok := filter.DateAfterOrEqual.Get(); ok {
		events = lo.Filter(events, func(event *schoolcalendardomain.Event, _ int) bool {
			return event.Date.After(dateAfterOrEqual) || event.Date == dateAfterOrEqual
		})
	}

	if dateBeforeOrEqual, ok := filter.DateBeforeOrEqual.Get(); ok {
		events = lo.Filter(events, func(event *schoolcalendardomain.Event, _ int) bool {
			return event.Date.Before(dateBeforeOrEqual) || event.Date == dateBeforeOrEqual
		})
	}

	return events
}
