package authv1conv

import (
	"fmt"

	authv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/auth/v1"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func FromPBProvider(pbProvider authv1.Provider) (authdomain.Provider, error) {
	switch pbProvider {
	case authv1.Provider_PROVIDER_GOOGLE:
		return authdomain.ProviderGoogle, nil
	case authv1.Provider_PROVIDER_APPLE:
		return authdomain.ProviderApple, nil
	case authv1.Provider_PROVIDER_TWITTER:
		return authdomain.ProviderTwitter, nil
	}
	return 0, fmt.Errorf("invalid %v", pbProvider)
}

func ToPBProvider(provider authdomain.Provider) (authv1.Provider, error) {
	switch provider {
	case authdomain.ProviderGoogle:
		return authv1.Provider_PROVIDER_GOOGLE, nil
	case authdomain.ProviderApple:
		return authv1.Provider_PROVIDER_APPLE, nil
	case authdomain.ProviderTwitter:
		return authv1.Provider_PROVIDER_TWITTER, nil
	}
	return 0, fmt.Errorf("invalid %v", provider)
}
