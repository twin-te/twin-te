type Provider = 'google' | 'apple' | 'twitter';

export const getLoginUrl = (provider: Provider) => {
	return `${process.env.NEXT_PUBLIC_AUTH_BASE_URL}/${provider}?redirect_url=${process.env.NEXT_PUBLIC_AUTH_REDIRECT_URL}`;
};

export const getLogoutUrl = () => {
	return `${process.env.NEXT_PUBLIC_AUTH_BASE_URL}/logout?redirect_url=${process.env.NEXT_PUBLIC_AUTH_REDIRECT_URL}`;
};
