/*
	[References]
		- https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api/authenticating_users_with_sign_in_with_apple
*/

package authv4

import (
	"context"
	"crypto/rsa"
	"crypto/subtle"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/twin-te/twin-te/back/appenv"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"golang.org/x/oauth2"
)

var appleOAuth2Config = &oauth2.Config{
	ClientID:     appenv.AUTH_APPLE_CLIENT_ID,
	ClientSecret: "",
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://appleid.apple.com/auth/authorize",
		TokenURL:  "https://appleid.apple.com/auth/token",
		AuthStyle: oauth2.AuthStyleInParams,
	},
	RedirectURL: appenv.AUTH_APPLE_CALLBACK_URL,
	Scopes:      []string{""},
}

var appleKeys = &applePublicKeyCache{
	httpClient: &http.Client{Timeout: 10 * time.Second},
	keys:       map[string]*rsa.PublicKey{},
}

func getAppleSocialID(ctx context.Context, code string) (socialID authdomain.SocialID, err error) {
	token, err := appleOAuth2Config.Exchange(ctx, code)
	if err != nil {
		return
	}

	if !token.Valid() {
		return "", fmt.Errorf("invalid token retrieved from apple, %+v", token)
	}

	idTokenString, ok := token.Extra("id_token").(string)
	if !ok {
		return "", errors.New("failed to retrieve id token for apple")
	}

	identity, err := verifyAppleIDToken(ctx, idTokenString, "")
	if err != nil {
		return "", err
	}
	return identity.SocialID, nil
}

type verifiedAppleIdentity struct {
	SocialID authdomain.SocialID
	Audience string
}

func verifyAppleIDToken(ctx context.Context, idTokenString string, expectedNonce string) (verifiedAppleIdentity, error) {
	// cf. https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api/verifying_a_user#3383769
	idToken, err := jwt.Parse(
		idTokenString,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %+v", t.Header["alg"])
			}
			kid, ok := t.Header["kid"].(string)
			if !ok || kid == "" {
				return nil, errors.New("apple token key id is missing")
			}
			return getAppleRSAPublicKey(ctx, kid)
		},
		jwt.WithExpirationRequired(),
		jwt.WithIssuer("https://appleid.apple.com"),
		jwt.WithValidMethods([]string{"RS256"}),
	)
	if err != nil {
		return verifiedAppleIdentity{}, err
	}
	audience, err := validateAppleAudience(idToken.Claims)
	if err != nil {
		return verifiedAppleIdentity{}, err
	}
	if expectedNonce != "" {
		claims, ok := idToken.Claims.(jwt.MapClaims)
		nonce, nonceOK := claims["nonce"].(string)
		if !ok || !nonceOK || subtle.ConstantTimeCompare([]byte(nonce), []byte(expectedNonce)) != 1 {
			return verifiedAppleIdentity{}, errors.New("invalid apple token nonce")
		}
	}

	sub, err := idToken.Claims.GetSubject()
	if err != nil {
		return verifiedAppleIdentity{}, err
	}
	socialID, err := authdomain.ParseSocialID(sub)
	if err != nil {
		return verifiedAppleIdentity{}, err
	}
	return verifiedAppleIdentity{SocialID: socialID, Audience: audience}, nil
}

func validateAppleAudience(claims jwt.Claims) (string, error) {
	audiences, err := claims.GetAudience()
	if err != nil {
		return "", err
	}
	for _, audience := range audiences {
		if slices.Contains(appenv.AUTH_APPLE_AUDIENCES, audience) {
			return audience, nil
		}
	}
	return "", fmt.Errorf("invalid apple token audience: %s", strings.Join(audiences, ","))
}

func exchangeAppleAuthorizationCode(
	ctx context.Context,
	code string,
	expectedIdentity verifiedAppleIdentity,
	expectedNonce string,
) (string, error) {
	config := *appleOAuth2Config
	config.ClientID = expectedIdentity.Audience
	config.RedirectURL = ""
	clientSecret, err := generateAppleClientSecretForClient(expectedIdentity.Audience)
	if err != nil {
		return "", err
	}
	config.ClientSecret = clientSecret
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return "", err
	}
	if token.RefreshToken == "" {
		return "", errors.New("apple refresh token is missing")
	}
	idTokenString, ok := token.Extra("id_token").(string)
	if !ok || idTokenString == "" {
		return "", errors.New("apple identity token is missing from token response")
	}
	exchangedIdentity, err := verifyAppleIDToken(ctx, idTokenString, expectedNonce)
	if err != nil {
		return "", err
	}
	if exchangedIdentity != expectedIdentity {
		return "", errors.New("apple authorization code does not match identity token")
	}
	return token.RefreshToken, nil
}

// getAppleRSAPublicKey returns apple's public key to verify the ID token signature.
//
// cf. https://developer.apple.com/documentation/sign_in_with_apple/fetch_apple_s_public_key_for_verifying_token_signature
func getAppleRSAPublicKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	return appleKeys.get(ctx, kid)
}

type applePublicKeyCache struct {
	mu          sync.Mutex
	httpClient  *http.Client
	keys        map[string]*rsa.PublicKey
	lastRefresh time.Time
}

func (cache *applePublicKeyCache) get(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	now := time.Now()
	if key := cache.keys[kid]; key != nil && now.Sub(cache.lastRefresh) < 6*time.Hour {
		return key, nil
	}
	if !cache.lastRefresh.IsZero() && now.Sub(cache.lastRefresh) < time.Minute {
		return nil, fmt.Errorf("apple public key %q is not cached", kid)
	}
	if err := cache.refresh(ctx, now); err != nil {
		return nil, err
	}
	if key := cache.keys[kid]; key != nil {
		return key, nil
	}
	return nil, fmt.Errorf("apple public key %q was not found", kid)
}

func (cache *applePublicKeyCache) refresh(ctx context.Context, now time.Time) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://appleid.apple.com/auth/keys", nil)
	if err != nil {
		return err
	}

	resp, err := cache.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("apple public key request failed: status=%d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return err
	}

	decodedBody := &struct {
		Keys []*appleJWK `json:"keys"`
	}{}

	err = json.Unmarshal(body, decodedBody)
	if err != nil {
		return err
	}

	keys := make(map[string]*rsa.PublicKey, len(decodedBody.Keys))
	for _, key := range decodedBody.Keys {
		if key.Kid == "" || key.Kty != "RSA" || key.Alg != "RS256" || key.Use != "sig" {
			continue
		}
		publicKey, err := key.RSAPublicKey()
		if err != nil {
			return err
		}
		keys[key.Kid] = publicKey
	}
	if len(keys) == 0 {
		return errors.New("apple public key response contains no usable keys")
	}
	cache.keys = keys
	cache.lastRefresh = now
	return nil
}

type appleJWK struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	Use string `json:"use"`
}

// RSAPublicKey returns the RSA public key.
//
// cf. https://stackoverflow.com/questions/66067321/marshal-appleids-public-key-to-rsa-publickey
func (jwk *appleJWK) RSAPublicKey() (*rsa.PublicKey, error) {
	modulesBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(jwk.N)
	if err != nil {
		return nil, err
	}
	modules := big.NewInt(0).SetBytes(modulesBytes)

	publicExponentBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(jwk.E)
	if err != nil {
		return nil, err
	}
	publicExponent := int(big.NewInt(0).SetBytes(publicExponentBytes).Int64())

	publicKey := &rsa.PublicKey{
		N: modules,
		E: publicExponent,
	}

	return publicKey, nil
}

// generateAppleClientSecret generates apple client secret.
// The generated client secret will expire in 35 days.
//
// [References]
//   - https://developer.apple.com/documentation/accountorganizationaldatasharing/creating-a-client-secret
//   - https://github.com/markbates/goth/blob/v1.79.0/providers/apple/apple.go#L71
func generateAppleClientSecret() (string, error) {
	return generateAppleClientSecretForClient(appenv.AUTH_APPLE_CLIENT_ID)
}

func generateAppleClientSecretForClient(clientID string) (string, error) {
	now := time.Now()
	iat := int(now.Unix())
	exp := int(now.Add(35 * 24 * time.Hour).Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": appenv.AUTH_APPLE_TEAM_ID,
		"iat": iat,
		"exp": exp,
		"aud": "https://appleid.apple.com",
		"sub": clientID,
	})
	token.Header["alg"] = "ES256"
	token.Header["kid"] = appenv.AUTH_APPLE_KEY_ID

	block, _ := pem.Decode([]byte(appenv.AUTH_APPLE_PRIVATE_KEY))
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("invalid apple private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	signedTokenString, err := token.SignedString(privateKey)
	return signedTokenString, err
}
