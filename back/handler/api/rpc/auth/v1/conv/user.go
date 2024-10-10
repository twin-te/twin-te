package authv1conv

import (
	"github.com/twin-te/twin-te/back/base"
	sharedconv "github.com/twin-te/twin-te/back/handler/api/rpc/shared/conv"
	authv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/auth/v1"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func ToPBUser(user *authdomain.User) (*authv1.User, error) {
	pbAuthentications, err := base.MapWithErr(user.Authentications, ToPBUserAuthentication)
	if err != nil {
		return nil, err
	}

	pbUser := &authv1.User{
		Id:              sharedconv.ToPBUUID(user.ID),
		Authentications: pbAuthentications,
	}

	return pbUser, nil
}
