package donationdomain

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

// PaymentUser is identified by one of the following fields.
//   - ID
//   - UserID
type PaymentUser struct {
	ID          idtype.PaymentUserID
	UserID      idtype.UserID
	DisplayName *string
	Link        *string

	EntityBeforeUpdated *PaymentUser
}

func (pu *PaymentUser) Clone() *PaymentUser {
	ret := lo.ToPtr(*pu)

	if pu.DisplayName != nil {
		ret.DisplayName = lo.ToPtr(*pu.DisplayName)
	}

	if pu.Link != nil {
		ret.Link = lo.ToPtr(*pu.Link)
	}

	return ret
}

func (pu *PaymentUser) BeforeUpdateHook() {
	pu.EntityBeforeUpdated = pu.Clone()
}

type PaymentUserDataToUpdate struct {
	DisplayName *string
	Link        *string
}

func (pu *PaymentUser) Update(data PaymentUserDataToUpdate) {
	if data.DisplayName != nil {
		pu.DisplayName = data.DisplayName
	}

	if data.Link != nil {
		pu.Link = data.Link
	}
}

func ConstructPaymentUser(fn func(pu *PaymentUser) (err error)) (*PaymentUser, error) {
	pu := new(PaymentUser)
	if err := fn(pu); err != nil {
		return nil, err
	}

	if pu.ID.IsZero() || pu.UserID.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", pu)
	}

	return pu, nil
}
