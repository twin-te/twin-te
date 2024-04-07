package authv3

import (
	"context"
	"fmt"

	"github.com/twin-te/twin-te/back/appenv"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleapisidtoken "google.golang.org/api/idtoken"
	googleapisoauth2 "google.golang.org/api/oauth2/v2"
	googleapisoption "google.golang.org/api/option"
)

var googleOAuth2Config = &oauth2.Config{
	ClientID:     appenv.AUTH_GOOGLE_CLIENT_ID,
	ClientSecret: appenv.AUTH_GOOGLE_CLIENT_SECRET,
	Endpoint:     google.Endpoint,
	RedirectURL:  appenv.AUTH_GOOGLE_CALLBACK_URL,
	Scopes:       []string{googleapisoauth2.OpenIDScope},
}

func getGoogleSocialID(ctx context.Context, code string) (socialID authdomain.SocialID, err error) {
	token, err := googleOAuth2Config.Exchange(ctx, code)
	if err != nil {
		return
	}

	if !token.Valid() {
		return "", fmt.Errorf("invalid token retrieved from google, %+v", token)
	}

	svc, err := googleapisoauth2.NewService(
		ctx,
		googleapisoption.WithTokenSource(googleOAuth2Config.TokenSource(ctx, token)),
	)
	if err != nil {
		return
	}

	tokenInfo, err := svc.Tokeninfo().Do()
	if err != nil {
		return
	}

	return authdomain.ParseSocialID(tokenInfo.UserId)
}

func verifyGoogleIDToken(ctx context.Context, idToken string) (socialID authdomain.SocialID, err error) {
	payload, err := googleapisidtoken.Validate(ctx, idToken, googleOAuth2Config.ClientID)
	if err != nil {
		return
	}
	return authdomain.ParseSocialID(payload.Subject)
}
