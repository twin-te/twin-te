package oauth2port

import (
	"net/url"
	"time"

	"github.com/samber/mo"
	oauth2domain "github.com/twin-te/twin-te/back/module/oauth2/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type Factory interface {
	NewClient(id, name, clientURI string, redirectURIs []string) (*oauth2domain.Client, error)
	NewRequest(
		kind oauth2domain.RequestKind,
		signature, requestID string,
		userID idtype.UserID,
		clientID string,
		requestedAt time.Time,
		requestedScopes, grantedScopes []string,
		form url.Values,
		session []byte,
		expiresAt mo.Option[time.Time],
	) (*oauth2domain.Request, error)
}
