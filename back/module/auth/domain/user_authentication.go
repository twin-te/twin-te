package authdomain

import (
	"fmt"
	"strings"

	"github.com/twin-te/twin-te/back/base"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Provider -trimprefix=Provider -output=provider_string.gen.go
type Provider int

const (
	ProviderGoogle Provider = iota + 1
	ProviderApple
	ProviderTwitter
)

var AllProviders = []Provider{
	ProviderGoogle,
	ProviderApple,
	ProviderTwitter,
}

func ParseProvider(s string) (Provider, error) {
	ret, ok := base.FindByString(AllProviders, s)
	if ok {
		return ret, nil
	}
	return 0, fmt.Errorf("failed to parse provider %v", s)
}

// SocialID represents provider's user id
type SocialID string

func (sid SocialID) String() string {
	return string(sid)
}

func (sid SocialID) IsZero() bool {
	return sid == ""
}

func ParseSocialID(s string) (SocialID, error) {
	v := strings.TrimSpace(s)
	if v == "" {
		return "", fmt.Errorf("social id must not be empty")
	}
	return SocialID(v), nil
}

type UserAuthentication struct {
	Provider Provider
	SocialID SocialID
}

func NewUserAuthentication(provider Provider, socialID SocialID) UserAuthentication {
	return UserAuthentication{Provider: provider, SocialID: socialID}
}
