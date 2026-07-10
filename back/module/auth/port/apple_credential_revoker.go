package authport

import "context"

type AppleCredentialRevoker interface {
	Revoke(ctx context.Context, clientID string, refreshToken string) error
}
