package authintegrator

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/twin-te/twin-te/back/appenv"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
)

type AppleCredentialRevoker struct {
	httpClient *http.Client
}

var _ authport.AppleCredentialRevoker = (*AppleCredentialRevoker)(nil)

func NewAppleCredentialRevoker() *AppleCredentialRevoker {
	return &AppleCredentialRevoker{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (r *AppleCredentialRevoker) Revoke(ctx context.Context, clientID string, refreshToken string) error {
	clientSecret, err := generateAppleClientSecret(clientID)
	if err != nil {
		return err
	}
	form := url.Values{
		"client_id":       []string{clientID},
		"client_secret":   []string{clientSecret},
		"token":           []string{refreshToken},
		"token_type_hint": []string{"refresh_token"},
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://appleid.apple.com/auth/revoke",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
	return fmt.Errorf("apple token revocation failed: status=%d body=%s", resp.StatusCode, string(body))
}

func generateAppleClientSecret(clientID string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": appenv.AUTH_APPLE_TEAM_ID,
		"iat": now.Unix(),
		"exp": now.Add(5 * time.Minute).Unix(),
		"aud": "https://appleid.apple.com",
		"sub": clientID,
	})
	token.Header["kid"] = appenv.AUTH_APPLE_KEY_ID

	block, _ := pem.Decode([]byte(appenv.AUTH_APPLE_PRIVATE_KEY))
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", fmt.Errorf("invalid apple private key")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	ecdsaKey, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("apple private key is not ECDSA")
	}
	return token.SignedString(ecdsaKey)
}
