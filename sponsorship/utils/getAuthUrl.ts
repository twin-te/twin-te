type Provider = 'google' | 'apple' | 'twitter';

export const getLoginUrl = (provider: Provider, redirectUrl = 'https://sponsorship.twinte.net') => {
	return `https://app.twinte.net/auth/v3/${provider}?redirect_url=${redirectUrl}`;
};

export const getLogoutUrl = (redirectUrl = 'https://sponsorship.twinte.net') => {
	return `https://app.twinte.net/auth/v3/logout?redirect_url=${redirectUrl}`;
};
