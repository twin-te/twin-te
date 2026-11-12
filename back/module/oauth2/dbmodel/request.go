package oauth2dbmodel

import (
	"encoding/json"
	"net/url"
	"time"

	oauth2domain "github.com/twin-te/twin-te/back/module/oauth2/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type Request struct {
	Kind            string `gorm:"primaryKey"`
	Signature       string `gorm:"primaryKey"`
	RequestID       string
	UserID          string
	ClientID        string
	RequestedAt     time.Time
	RequestedScopes []string        `gorm:"column:requested_scopes;serializer:json"`
	GrantedScopes   []string        `gorm:"column:granted_scopes;serializer:json"`
	Form            url.Values      `gorm:"serializer:json"`
	Session         json.RawMessage `gorm:"serializer:json"`
	Active          bool
	ExpiresAt       time.Time
}

func (Request) TableName() string { return "oauth2_requests" }

func FromDBRequest(dbRequest *Request) (*oauth2domain.Request, error) {
	return oauth2domain.ConstructRequest(func(r *oauth2domain.Request) (err error) {
		r.Kind = oauth2domain.RequestKind(dbRequest.Kind)
		r.Signature = dbRequest.Signature
		r.RequestID = dbRequest.RequestID
		r.UserID, err = idtype.ParseUserID(dbRequest.UserID)
		if err != nil {
			return err
		}
		r.ClientID = dbRequest.ClientID
		r.RequestedAt = dbRequest.RequestedAt
		r.RequestedScopes = dbRequest.RequestedScopes
		r.GrantedScopes = dbRequest.GrantedScopes
		r.Form = dbRequest.Form
		r.Session = dbRequest.Session
		r.Active = dbRequest.Active
		r.ExpiresAt = dbRequest.ExpiresAt
		return nil
	})
}

func ToDBRequest(request *oauth2domain.Request) *Request {
	return &Request{
		Kind:            string(request.Kind),
		Signature:       request.Signature,
		RequestID:       request.RequestID,
		UserID:          request.UserID.String(),
		ClientID:        request.ClientID,
		RequestedAt:     request.RequestedAt,
		RequestedScopes: request.RequestedScopes,
		GrantedScopes:   request.GrantedScopes,
		Form:            request.Form,
		Session:         request.Session,
		Active:          request.Active,
		ExpiresAt:       request.ExpiresAt,
	}
}
