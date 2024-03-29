import { apiClient } from './apiClient';
import { loadStripe } from '@stripe/stripe-js/pure';
import { stripePublicKey } from '../utils/stripe';

interface stripeResponse {
	sessionId: string;
}

export const registOneTime = async (price: number) => {
	const stripe = await loadStripe(stripePublicKey);
	apiClient
		.post<stripeResponse>('/donation/session/onetime', {
			amount: price
		})
		.then((response) => {
			const sessionId = response.data.sessionId;
			stripe
				?.redirectToCheckout({
					sessionId
				})
				.then(function (err) {
					console.error(err);
				});
		});
};

export const registSubscription = async (plan: string) => {
	const stripe = await loadStripe(stripePublicKey);
	apiClient
		.post<stripeResponse>('/donation/session/subscription', {
			planId: plan
		})
		.then((response) => {
			const sessionId = response.data.sessionId;
			stripe
				?.redirectToCheckout({
					sessionId
				})
				.then(function (err) {
					console.error(err);
				});
		});
};

export const cancelSubscription = async (id: string) => {
	await apiClient.delete(`/donation/subscriptions/${id}`);
};
