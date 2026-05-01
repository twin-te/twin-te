package calendardbmodel

import (
	"time"

	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type IcalSubscription struct {
	ID        string
	UserID    string
	
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToDBIcalSubscription(id idtype.IcalSubscriptionID, userID idtype.UserID) *IcalSubscription {
	return &IcalSubscription{
		ID:     id.String(),
		UserID: userID.String(),
	}
}
