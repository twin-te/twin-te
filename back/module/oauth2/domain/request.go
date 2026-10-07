package oauth2domain

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type RequestKind string

const (
	RequestKindAuthorizeCode RequestKind = "authorize_code"
	RequestKindAccessToken   RequestKind = "access_token"
	RequestKindRefreshToken  RequestKind = "refresh_token"
	RequestKindPKCE          RequestKind = "pkce"
)

// Request is identified by the following fields.
//   - Kind
//   - Signature
type Request struct {
	Kind            RequestKind
	Signature       string
	RequestID       string
	UserID          idtype.UserID
	ClientID        string
	RequestedAt     time.Time
	RequestedScopes []string
	GrantedScopes   []string
	Form            url.Values
	Session         json.RawMessage
	Active          bool
	ExpiresAt       time.Time
}

func ConstructRequest(fn func(r *Request) (err error)) (*Request, error) {
	r := new(Request)
	if err := fn(r); err != nil {
		return nil, err
	}

	if r.Kind == "" ||
		r.Signature == "" ||
		r.RequestID == "" ||
		r.UserID.IsZero() ||
		r.ClientID == "" ||
		r.RequestedAt.IsZero() ||
		r.ExpiresAt.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", r)
	}

	return r, nil
}
