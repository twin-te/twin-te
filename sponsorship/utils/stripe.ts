export const stripePublicKey = 'pk_live_wnnqlGjrN71cV9uwBXDvNQJP009Chm8Nt0';

const stripeSubscription200yenID = 'plan_H9D4eZ0Vohpqpy';
const stripeSubscription500yenID = 'plan_H9D4AJchCmsejL';
const stripeSubscription1000yenID = 'plan_H9D48FqtiALjlL';
export const subscriptions = [
	{ planId: stripeSubscription200yenID, amount: '200' },
	{ planId: stripeSubscription500yenID, amount: '500' },
	{ planId: stripeSubscription1000yenID, amount: '1000' }
];
