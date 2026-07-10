package authfactory

import (
	"testing"
	"time"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func TestNewAuthChallengeCreatesUniqueExpiringNonce(t *testing.T) {
	now := time.Date(2026, time.July, 10, 12, 0, 0, 0, time.UTC)
	factory := New(func() time.Time { return now })

	first, err := factory.NewAuthChallenge(authdomain.ProviderGoogle)
	if err != nil {
		t.Fatalf("create first challenge: %v", err)
	}
	second, err := factory.NewAuthChallenge(authdomain.ProviderGoogle)
	if err != nil {
		t.Fatalf("create second challenge: %v", err)
	}

	if first.ID == second.ID || first.Nonce == second.Nonce {
		t.Fatal("authentication challenges must be unique")
	}
	if first.Provider != authdomain.ProviderGoogle {
		t.Fatalf("provider = %v, want Google", first.Provider)
	}
	if want := now.Add(authdomain.AuthChallengeLifeTime); !first.ExpiredAt.Equal(want) {
		t.Fatalf("expiration = %v, want %v", first.ExpiredAt, want)
	}
	if len(first.Nonce) < 32 {
		t.Fatalf("nonce is unexpectedly short: %q", first.Nonce)
	}
}
