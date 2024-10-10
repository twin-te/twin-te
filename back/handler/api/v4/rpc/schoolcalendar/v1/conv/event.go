package schoolcalendarv1conv

import (
	"github.com/twin-te/twin-te/back/base"
	sharedconv "github.com/twin-te/twin-te/back/handler/api/v4/rpc/shared/conv"
	schoolcalendarv1 "github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/schoolcalendar/v1"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
)

func ToPBEvent(event *schoolcalendardomain.Event) (*schoolcalendarv1.Event, error) {
	pbEventType, err := ToPBEventType(event.Type)
	if err != nil {
		return nil, err
	}

	pbEvent := &schoolcalendarv1.Event{
		Id:          int32(event.ID.Int()),
		Type:        pbEventType,
		Date:        sharedconv.ToPBRFC3339FullDate(event.Date),
		Description: event.Description,
	}

	pbWeekday, err := base.OptionMapWithErr(event.ChangeTo, sharedconv.ToPBWeekday)
	if err != nil {
		return nil, err
	}
	pbEvent.ChangeTo = pbWeekday.ToPointer()

	return pbEvent, nil
}
