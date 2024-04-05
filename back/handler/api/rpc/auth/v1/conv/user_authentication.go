package authv1conv

import (
	authv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/auth/v1"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func FromPBUserAuthentication(pbUserAuthentication *authv1.UserAuthentication) (authdomain.UserAuthentication, error) {
	provider, err := FromPBProvider(pbUserAuthentication.Provider)
	if err != nil {
		return authdomain.UserAuthentication{}, err
	}

	socialID, err := authdomain.ParseSocialID(pbUserAuthentication.SocialId)
	if err != nil {
		return authdomain.UserAuthentication{}, err
	}

	userAuthentication := authdomain.NewUserAuthentication(provider, socialID)

	return userAuthentication, nil
}

func ToPBUserAuthentication(userAuthentication authdomain.UserAuthentication) (*authv1.UserAuthentication, error) {
	pbProvider, err := ToPBProvider(userAuthentication.Provider)
	if err != nil {
		return nil, err
	}

	pbUserAuthentications := &authv1.UserAuthentication{
		SocialId: userAuthentication.SocialID.String(),
		Provider: pbProvider,
	}

	return pbUserAuthentications, nil
}
