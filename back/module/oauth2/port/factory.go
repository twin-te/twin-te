package oauth2port

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/samber/mo"
	oauth2domain "github.com/twin-te/twin-te/back/module/oauth2/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type NewRequestParams struct {
	Kind            oauth2domain.RequestKind
	Signature       string
	RequestID       string
	UserID          idtype.UserID
	ClientID        string
	RequestedAt     time.Time
	RequestedScopes []string
	GrantedScopes   []string
	Form            url.Values
	Session         json.RawMessage
	ExpiresAt       mo.Option[time.Time]
}

type Factory interface {
	NewClient(id, name, clientURI string, redirectURIs []string) (*oauth2domain.Client, error)
	NewRequest(params NewRequestParams) (*oauth2domain.Request, error)
}
