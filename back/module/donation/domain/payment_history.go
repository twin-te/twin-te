package donationdomain

import (
	"fmt"
	"time"

	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

type PaymentType int

func (pt PaymentType) IsZero() bool {
	return pt == 0
}

const (
	PaymentTypeOneTime PaymentType = iota + 1
	PaymentTypeSubscription
)

type PaymentStatus int

func (ps PaymentStatus) IsZero() bool {
	return ps == 0
}

const (
	PaymentStatusPending PaymentStatus = iota + 1
	PaymentStatusCanceled
	PaymentStatusSucceeded
)

// PaymentHistory is identified by one of the following fields.
//   - ID
type PaymentHistory struct {
	ID            idtype.PaymentHistoryID
	PaymentUserID *idtype.PaymentUserID
	Type          PaymentType
	Status        PaymentStatus
	Amount        int
	CreatedAt     time.Time

	EntityBeforeUpdated *PaymentHistory
}

func ConstructPaymentHistory(fn func(ph *PaymentHistory) (err error)) (*PaymentHistory, error) {
	ph := new(PaymentHistory)
	if err := fn(ph); err != nil {
		return nil, err
	}

	if ph.ID.IsZero() || ph.Type.IsZero() || ph.Status.IsZero() || ph.Amount <= 0 || ph.CreatedAt.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", ph)
	}

	return ph, nil
}
