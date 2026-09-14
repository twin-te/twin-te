package authdomain_test

import (
	"testing"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func TestProvider_String(t *testing.T) {
	tests := []struct {
		name string
		p    authdomain.Provider
		want string
	}{
		{"google", authdomain.ProviderGoogle, "Google"},
		{"apple", authdomain.ProviderApple, "Apple"},
		{"twitter", authdomain.ProviderTwitter, "Twitter"},
		{"unknown", authdomain.Provider(999), "Provider(999)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestParseProvider(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    authdomain.Provider
		wantErr bool
	}{
		{"google", "Google", authdomain.ProviderGoogle, false},
		{"apple", "Apple", authdomain.ProviderApple, false},
		{"twitter", "Twitter", authdomain.ProviderTwitter, false},
		{"invalid", "invalid", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := authdomain.ParseProvider(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSocialID_String(t *testing.T) {
	sid := authdomain.SocialID("abc")
	if sid.String() != "abc" {
		t.Errorf("String() = %q, want %q", sid.String(), "abc")
	}
}

func TestSocialID_IsZero(t *testing.T) {
	tests := []struct {
		name string
		sid  authdomain.SocialID
		want bool
	}{
		{"empty", authdomain.SocialID(""), true},
		{"non-empty", authdomain.SocialID("abc"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sid.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSocialID(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    authdomain.SocialID
		wantErr bool
	}{
		{"valid", "abc", authdomain.SocialID("abc"), false},
		{"trims whitespace", "  abc  ", authdomain.SocialID("abc"), false},
		{"empty", "", "", true},
		{"whitespace only", "   ", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := authdomain.ParseSocialID(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserAuthentication(t *testing.T) {
	ua := authdomain.NewUserAuthentication(authdomain.ProviderGoogle, authdomain.SocialID("abc"))
	if ua.Provider != authdomain.ProviderGoogle || ua.SocialID != authdomain.SocialID("abc") {
		t.Errorf("got %+v", ua)
	}
}
