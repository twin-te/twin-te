package authv4

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestVerifyAppleIDTokenRejectsMissingKeyID(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate test key: %v", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": "https://appleid.apple.com",
		"aud": appAudienceForTest(),
		"sub": "apple-user",
		"exp": time.Now().Add(time.Minute).Unix(),
	})
	signed, err := token.SignedString(privateKey)
	if err != nil {
		t.Fatalf("sign token: %v", err)
	}

	_, err = verifyAppleIDToken(context.Background(), signed, "")
	if err == nil || !strings.Contains(err.Error(), "key id is missing") {
		t.Fatalf("expected missing key id error, got %v", err)
	}
}

func appAudienceForTest() string {
	if len(appleOAuth2Config.ClientID) != 0 {
		return appleOAuth2Config.ClientID
	}
	return "test-client"
}
