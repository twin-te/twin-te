package donationdomain

import (
	"errors"
	"fmt"
	"time"

	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

// Subscription is identified by one of the following fields.
//   - ID
type Subscription struct {
	ID            idtype.SubscriptionID
	PaymentUserID idtype.PaymentUserID
	Plans         []*SubscriptionPlan
	IsActive      bool
	CreatedAt     time.Time
}

type SubscriptionPlan struct {
	ID     idtype.SubscriptionPlanID
	Name   string
	Amount int
}

func ConstructSubscription(fn func(s *Subscription) (err error)) (*Subscription, error) {
	s := new(Subscription)
	if err := fn(s); err != nil {
		return nil, err
	}

	if len(s.Plans) == 0 {
		return nil, errors.New("subscription must have at least one plan")
	}

	if s.ID.IsZero() || s.PaymentUserID.IsZero() || s.CreatedAt.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", s)
	}

	return s, nil
}

func ConstructSubscriptionPlan(fn func(sp *SubscriptionPlan) (err error)) (*SubscriptionPlan, error) {
	sp := new(SubscriptionPlan)
	if err := fn(sp); err != nil {
		return nil, err
	}

	if sp.ID.IsZero() || sp.Name == "" || sp.Amount <= 0 {
		return nil, fmt.Errorf("failed to construct %+v", sp)
	}

	return sp, nil
}
