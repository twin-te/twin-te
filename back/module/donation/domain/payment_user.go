package donationdomain

import (
	"fmt"
	"net/url"

	"github.com/samber/lo"
	"github.com/samber/mo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var ParseDisplayName = shareddomain.NewRequiredStringParser("display name")

type Link string

func (l Link) String() string {
	return string(l)
}

func ParseLink(s string) (Link, error) {
	uri, err := url.ParseRequestURI(s)
	if err != nil || !uri.IsAbs() {
		return "", fmt.Errorf("failed to parse link %v", s)
	}
	return Link(s), nil
}

// PaymentUser is identified by one of the following fields.
//   - ID
type PaymentUser struct {
	ID          idtype.PaymentUserID
	UserID      idtype.UserID
	DisplayName mo.Option[shareddomain.RequiredString]
	Link        mo.Option[Link]

	BeforeUpdated mo.Option[*PaymentUser]
}

func (pu *PaymentUser) Clone() *PaymentUser {
	ret := lo.ToPtr(*pu)
	return ret
}

func (pu *PaymentUser) BeforeUpdateHook() {
	pu.BeforeUpdated = mo.Some(pu.Clone())
}

type PaymentUserDataToUpdate struct {
	DisplayName mo.Option[mo.Option[shareddomain.RequiredString]]
	Link        mo.Option[mo.Option[Link]]
}

func (pu *PaymentUser) Update(data PaymentUserDataToUpdate) {
	if displayName, ok := data.DisplayName.Get(); ok {
		pu.DisplayName = displayName
	}

	if link, ok := data.Link.Get(); ok {
		pu.Link = link
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
