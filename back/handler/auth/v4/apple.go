/*
	[References]
		- https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api/authenticating_users_with_sign_in_with_apple
*/

package authv4

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
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

	// cf. https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api/verifying_a_user#3383769
	idToken, err := jwt.Parse(
		idTokenString,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %+v", t.Header["alg"])
			}
			return getAppleRSAPublicKey(ctx, t.Header["kid"].(string))
		},
		jwt.WithAudience(appenv.AUTH_APPLE_CLIENT_ID),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer("https://appleid.apple.com"),
		jwt.WithValidMethods([]string{"RS256"}),
	)
	if err != nil {
		return
	}

	sub, err := idToken.Claims.GetSubject()
	if err != nil {
		return
	}

	return authdomain.ParseSocialID(sub)
}

// getAppleRSAPublicKey returns apple's public key to verify the ID token signature.
//
// cf. https://developer.apple.com/documentation/sign_in_with_apple/fetch_apple_s_public_key_for_verifying_token_signature
func getAppleRSAPublicKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://appleid.apple.com/auth/keys", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	decodedBody := &struct {
		Keys []*appleJWK `json:"keys"`
	}{}

	err = json.Unmarshal(body, decodedBody)
	if err != nil {
		return nil, err
	}

	for _, key := range decodedBody.Keys {
		if key.Kid == kid {
			return key.RSAPublicKey()
		}
	}

	return nil, fmt.Errorf("not found apple public key whose kid is %s", kid)
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
	publicExponent := int(binary.BigEndian.Uint32(publicExponentBytes))

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
	now := time.Now()
	iat := int(now.Unix())
	exp := int(now.Add(35 * 24 * time.Hour).Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": appenv.AUTH_APPLE_TEAM_ID,
		"iat": iat,
		"exp": exp,
		"aud": "https://appleid.apple.com",
		"sub": appenv.AUTH_APPLE_CLIENT_ID,
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
