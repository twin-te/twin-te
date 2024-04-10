import { loadStripe } from '@stripe/stripe-js/pure';

const stripePublicKey = process.env.NEXT_PUBLIC_STRIPE_PUBLIC_KEY as string;

const stripeSubscription200yenID = 'plan_H9D4eZ0Vohpqpy';
const stripeSubscription500yenID = 'plan_H9D4AJchCmsejL';
const stripeSubscription1000yenID = 'plan_H9D48FqtiALjlL';

export const subscriptions = [
	{ planId: stripeSubscription200yenID, amount: '200' },
	{ planId: stripeSubscription500yenID, amount: '500' },
	{ planId: stripeSubscription1000yenID, amount: '1000' }
];

export const redirectToCheckout = async (sessionId: string) => {
	const stripe = await loadStripe(stripePublicKey);
	stripe
		?.redirectToCheckout({
			sessionId
		})
		.then(function (err) {
			console.error(err);
		});
};
