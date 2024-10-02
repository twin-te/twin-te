import {
	ENV_NEXT_PUBLIC_AUTH_BASE_URL,
	ENV_NEXT_PUBLIC_AUTH_REDIRECT_URL,
} from "@/env";

type Provider = "google" | "apple" | "twitter";

export const getLoginUrl = (provider: Provider) => {
	return `${ENV_NEXT_PUBLIC_AUTH_BASE_URL}/${provider}?redirect_url=${ENV_NEXT_PUBLIC_AUTH_REDIRECT_URL}`;
};

export const getLogoutUrl = () => {
	return `${ENV_NEXT_PUBLIC_AUTH_BASE_URL}/logout?redirect_url=${ENV_NEXT_PUBLIC_AUTH_REDIRECT_URL}`;
};
