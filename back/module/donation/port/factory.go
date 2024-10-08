package donationport

import (
	"github.com/samber/mo"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type Factory interface {
	NewPaymentUser(userID idtype.UserID, displayName mo.Option[shareddomain.RequiredString], link mo.Option[donationdomain.Link]) (*donationdomain.PaymentUser, error)
}
