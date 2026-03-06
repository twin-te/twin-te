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

func FromDBIcalSubscription(id, userID string) (idtype.IcalSubscriptionID, idtype.UserID, error) {
	subID, err := idtype.ParseIcalSubscriptionID(id)
	if err != nil {
		return idtype.IcalSubscriptionID{}, idtype.UserID{}, err
	}
	uid, err := idtype.ParseUserID(userID)
	if err != nil {
		return idtype.IcalSubscriptionID{}, idtype.UserID{}, err
	}
	return subID, uid, nil
}

func ToDBIcalSubscription(id idtype.IcalSubscriptionID, userID idtype.UserID) *IcalSubscription {
	return &IcalSubscription{
		ID:     id.String(),
		UserID: userID.String(),
	}
}
