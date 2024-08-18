package donationdomain

import (
	"fmt"
	"net/url"

	"github.com/samber/lo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var ParseDisplayName = shareddomain.NewRequiredStringParser("display name")

type Link string

func (l Link) String() string {
	return string(l)
}

func (l *Link) StringPtr() *string {
	if l == nil {
		return nil
	}
	return lo.ToPtr(l.String())
}

func ParseLink(s string) (Link, error) {
	uri, err := url.ParseRequestURI(s)
	if err != nil || !uri.IsAbs() {
		return "", fmt.Errorf("failed to parse link %#v", s)
	}
	return Link(s), nil
}

// PaymentUser is identified by one of the following fields.
//   - ID
//   - UserID
type PaymentUser struct {
	ID          idtype.PaymentUserID
	UserID      idtype.UserID
	DisplayName *shareddomain.RequiredString
	Link        *Link

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
	DisplayName **shareddomain.RequiredString
	Link        **Link
}

func (pu *PaymentUser) Update(data PaymentUserDataToUpdate) {
	if data.DisplayName != nil {
		pu.DisplayName = *data.DisplayName
	}

	if data.Link != nil {
		pu.Link = *data.Link
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
