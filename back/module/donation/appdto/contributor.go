package donationappdto

import (
	"github.com/samber/mo"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

type Contributor struct {
	DisplayName shareddomain.RequiredString
	Link        mo.Option[donationdomain.Link]
}
