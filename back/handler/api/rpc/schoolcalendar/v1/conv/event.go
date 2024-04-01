package schoolcalendarv1conv

import (
	sharedconv "github.com/twin-te/twinte-back/handler/api/rpc/shared/conv"
	schoolcalendarv1 "github.com/twin-te/twinte-back/handler/api/rpcgen/schoolcalendar/v1"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
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

	if event.ChangeTo != nil {
		pbWeekday, err := sharedconv.ToPBWeekday(*event.ChangeTo)
		if err != nil {
			return nil, err
		}
		pbEvent.ChangeTo = &pbWeekday
	}

	return pbEvent, nil
}
