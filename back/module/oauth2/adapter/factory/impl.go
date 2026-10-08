package oauth2factory

import (
	"time"

	oauth2domain "github.com/twin-te/twin-te/back/module/oauth2/domain"
	oauth2port "github.com/twin-te/twin-te/back/module/oauth2/port"
)

var _ oauth2port.Factory = (*impl)(nil)

const fallbackRequestLifetime = 30 * 24 * time.Hour

type impl struct {
	nowFunc func() time.Time
}

func (f *impl) NewClient(id, name, clientURI string, redirectURIs []string) (*oauth2domain.Client, error) {
	return oauth2domain.ConstructClient(func(c *oauth2domain.Client) error {
		c.ID = id
		c.Name = name
		c.ClientURI = clientURI
		c.RedirectURIs = redirectURIs
		return nil
	})
}

func (f *impl) NewRequest(params oauth2port.NewRequestParams) (*oauth2domain.Request, error) {
	return oauth2domain.ConstructRequest(func(r *oauth2domain.Request) error {
		r.Kind = params.Kind
		r.Signature = params.Signature
		r.RequestID = params.RequestID
		r.UserID = params.UserID
		r.ClientID = params.ClientID
		r.RequestedAt = params.RequestedAt.UTC().Truncate(time.Microsecond)
		r.RequestedScopes = params.RequestedScopes
		r.GrantedScopes = params.GrantedScopes
		r.Form = params.Form
		r.Session = params.Session
		r.Active = true
		r.ExpiresAt = params.ExpiresAt.OrElse(f.nowFunc().Add(fallbackRequestLifetime)).UTC().Truncate(time.Microsecond)
		return nil
	})
}

func New(nowFunc func() time.Time) *impl {
	return &impl{nowFunc: nowFunc}
}
