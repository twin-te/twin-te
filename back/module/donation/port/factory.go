package donationport

import (
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type Factory interface {
	NewPaymentUser(userID idtype.UserID, displayName *shareddomain.RequiredString, link *donationdomain.Link) (*donationdomain.PaymentUser, error)
}
