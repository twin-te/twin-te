package donationport

import (
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type Factory interface {
	NewPaymentUser(userID idtype.UserID, displayName *string, link *string) (*donationdomain.PaymentUser, error)
}
