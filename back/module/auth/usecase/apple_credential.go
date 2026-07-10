package authusecase

import (
	"context"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func (uc *impl) SaveAppleCredential(ctx context.Context, credential *authdomain.AppleCredential) error {
	return uc.r.SaveAppleCredential(ctx, credential)
}
