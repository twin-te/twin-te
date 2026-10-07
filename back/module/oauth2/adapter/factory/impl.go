package oauth2factory

import (
	"net/url"
	"time"

	"github.com/samber/mo"
	oauth2domain "github.com/twin-te/twin-te/back/module/oauth2/domain"
	oauth2port "github.com/twin-te/twin-te/back/module/oauth2/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var _ oauth2port.Factory = (*impl)(nil)

// fallbackRequestLifetime is used only when the expiry is not given.
const fallbackRequestLifetime = 30 * 24 * time.Hour

type impl struct {
	nowFunc func() time.Time
}

func (f *impl) NewClient(id, name, clientURI string, redirectURIs []string) (*oauth2domain.Client, error) {
	return oauth2domain.ConstructClient(func(c *oauth2domain.Client) (err error) {
		c.ID = id
		c.Name = name
		c.ClientURI = clientURI
		c.RedirectURIs = redirectURIs
		return
	})
}

func (f *impl) NewRequest(
	kind oauth2domain.RequestKind,
	signature, requestID string,
	userID idtype.UserID,
	clientID string,
	requestedAt time.Time,
	requestedScopes, grantedScopes []string,
	form url.Values,
	session []byte,
	expiresAt mo.Option[time.Time],
) (*oauth2domain.Request, error) {
	return oauth2domain.ConstructRequest(func(r *oauth2domain.Request) (err error) {
		r.Kind = kind
		r.Signature = signature
		r.RequestID = requestID
		r.UserID = userID
		r.ClientID = clientID
		r.RequestedAt = requestedAt.UTC().Truncate(time.Microsecond)
		r.RequestedScopes = requestedScopes
		r.GrantedScopes = grantedScopes
		r.Form = form
		r.Session = session
		r.Active = true
		r.ExpiresAt = expiresAt.OrElse(f.nowFunc().Add(fallbackRequestLifetime)).UTC().Truncate(time.Microsecond)
		return
	})
}

func New(nowFunc func() time.Time) *impl {
	return &impl{nowFunc: nowFunc}
}
