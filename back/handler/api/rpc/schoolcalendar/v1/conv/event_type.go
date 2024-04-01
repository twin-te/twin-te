package schoolcalendarv1conv

import (
	"fmt"

	schoolcalendarv1 "github.com/twin-te/twinte-back/handler/api/rpcgen/schoolcalendar/v1"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
)

func FromPBEventType(pbEventType schoolcalendarv1.EventType) (schoolcalendardomain.EventType, error) {
	switch pbEventType {
	case schoolcalendarv1.EventType_EVENT_TYPE_HOLIDAY:
		return schoolcalendardomain.EventTypeHoliday, nil
	case schoolcalendarv1.EventType_EVENT_TYPE_PUBLIC_HOLIDAY:
		return schoolcalendardomain.EventTypePublicHoliday, nil
	case schoolcalendarv1.EventType_EVENT_TYPE_EXAM:
		return schoolcalendardomain.EventTypeExam, nil
	case schoolcalendarv1.EventType_EVENT_TYPE_SUBSTITUTE_DAY:
		return schoolcalendardomain.EventTypeSubstituteDay, nil
	case schoolcalendarv1.EventType_EVENT_TYPE_OTHER:
		return schoolcalendardomain.EventTypeOther, nil
	}
	return 0, fmt.Errorf("invalid %#v", pbEventType)
}

func ToPBEventType(eventType schoolcalendardomain.EventType) (schoolcalendarv1.EventType, error) {
	switch eventType {
	case schoolcalendardomain.EventTypeHoliday:
		return schoolcalendarv1.EventType_EVENT_TYPE_HOLIDAY, nil
	case schoolcalendardomain.EventTypePublicHoliday:
		return schoolcalendarv1.EventType_EVENT_TYPE_PUBLIC_HOLIDAY, nil
	case schoolcalendardomain.EventTypeExam:
		return schoolcalendarv1.EventType_EVENT_TYPE_EXAM, nil
	case schoolcalendardomain.EventTypeSubstituteDay:
		return schoolcalendarv1.EventType_EVENT_TYPE_SUBSTITUTE_DAY, nil
	case schoolcalendardomain.EventTypeOther:
		return schoolcalendarv1.EventType_EVENT_TYPE_OTHER, nil
	}
	return 0, fmt.Errorf("invalid %#v", eventType)
}
