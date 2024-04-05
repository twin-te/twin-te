package donationdomain

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

// SubscriptionPlan is identified by one of the following fields.
//   - ID
type SubscriptionPlan struct {
	ID     idtype.SubscriptionPlanID
	Name   string
	Amount int
}

func (sp *SubscriptionPlan) Clone() *SubscriptionPlan {
	ret := lo.ToPtr(*sp)
	return ret
}

// Subscription is identified by one of the following fields.
//   - ID
type Subscription struct {
	ID            idtype.SubscriptionID
	PaymentUserID idtype.PaymentUserID
	PlanID        idtype.SubscriptionPlanID
	IsActive      bool
	CreatedAt     time.Time

	PlanAssociation shareddomain.Association[*SubscriptionPlan]
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

func ConstructSubscription(fn func(s *Subscription) (err error)) (*Subscription, error) {
	s := new(Subscription)
	if err := fn(s); err != nil {
		return nil, err
	}

	if s.ID.IsZero() || s.PaymentUserID.IsZero() || s.PlanID.IsZero() || s.CreatedAt.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", s)
	}

	return s, nil
}
