import { ENV_NEXT_PUBLIC_STRIPE_PUBLIC_KEY } from '@/env';
import { loadStripe } from '@stripe/stripe-js/pure';

export const redirectToCheckout = async (sessionId: string) => {
	const stripe = await loadStripe(ENV_NEXT_PUBLIC_STRIPE_PUBLIC_KEY);
	stripe
		?.redirectToCheckout({
			sessionId
		})
		.then(({ error }) => {
			console.error(error);
		});
};
