package calendardbmodel

import (
	"time"

	"github.com/twin-te/twin-te/back/base"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type IcalSubscription struct {
	ID     string
	UserID string
	Mode   string

	TargetTags []IcalSubscriptionTargetTag

	CreatedAt time.Time
	UpdatedAt time.Time
}

type IcalSubscriptionTargetTag struct {
	IcalSubscriptionID string
	TagID              string
}

func ToDBIcalSubscription(s *calendardomain.IcalSubscription) *IcalSubscription {
	return &IcalSubscription{
		ID:     s.ID.String(),
		UserID: s.UserID.String(),
		Mode:   s.Mode.String(),
		TargetTags: base.Map(s.TargetTagIDs, func(tagID idtype.TagID) IcalSubscriptionTargetTag {
			return IcalSubscriptionTargetTag{
				IcalSubscriptionID: s.ID.String(),
				TagID:              tagID.String(),
			}
		}),
	}
}

func FromDBIcalSubscription(dbSub *IcalSubscription) (*calendardomain.IcalSubscription, error) {
	return calendardomain.ConstructIcalSubscription(func(s *calendardomain.IcalSubscription) (err error) {
		s.ID, err = idtype.ParseIcalSubscriptionID(dbSub.ID)
		if err != nil {
			return err
		}

		s.UserID, err = idtype.ParseUserID(dbSub.UserID)
		if err != nil {
			return err
		}

		s.Mode, err = calendardomain.ParseIcalSubscriptionMode(dbSub.Mode)
		if err != nil {
			return err
		}

		s.TargetTagIDs, err = base.MapWithErr(dbSub.TargetTags, func(t IcalSubscriptionTargetTag) (idtype.TagID, error) {
			return idtype.ParseTagID(t.TagID)
		})
		if err != nil {
			return err
		}

		return nil
	})
}
