import { loadStripe } from '@stripe/stripe-js/pure';

const stripePublicKey = process.env.NEXT_PUBLIC_STRIPE_PUBLIC_KEY as string;

export const redirectToCheckout = async (sessionId: string) => {
	const stripe = await loadStripe(stripePublicKey);
	stripe
		?.redirectToCheckout({
			sessionId
		})
		.then(({ error }) => {
			console.error(error);
		});
};
