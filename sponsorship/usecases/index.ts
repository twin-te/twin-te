import { createPromiseClient, PromiseClient, Transport } from '@connectrpc/connect';
import { Payment, Plan, Subscription, User } from '../domain';
import { DonationService } from '../api/gen/donation/v1/service_connect';
import { AuthService } from '../api/gen/auth/v1/service_connect';
import { fromPBPayment, fromPBPaymentUser, fromPBPlan, fromPBSubscription } from '../api/converters/donationv1';
import { assurePresence } from '../api/converters/utils';
import { createConnectTransport } from '@connectrpc/connect-web';
import { ConvertAPIError, isUnauthenticatedError } from './error';

class UseCase {
	#authClient: PromiseClient<typeof AuthService>;
	#donationClient: PromiseClient<typeof DonationService>;

	constructor(transport: Transport) {
		this.#authClient = createPromiseClient(AuthService, transport);
		this.#donationClient = createPromiseClient(DonationService, transport);
	}

	async getUser(): Promise<User> {
		return this.#donationClient
			.getPaymentUser({})
			.then((res) => fromPBPaymentUser(assurePresence(res.paymentUser)))
			.catch(ConvertAPIError);
	}

	async updateUserInfo(newDisplayName: string, newLink: string): Promise<User> {
		return this.#donationClient
			.updatePaymentUser({ displayName: newDisplayName, link: newLink })
			.then((res) => fromPBPaymentUser(assurePresence(res.paymentUser)))
			.catch(ConvertAPIError);
	}

	async makeOneTimeDonation(price: number): Promise<string> {
		return this.#donationClient
			.createOneTimeCheckoutSession({ amount: price })
			.then((res) => res.checkoutSessionId)
			.catch(ConvertAPIError);
	}

	async registerSubscription(planId: string): Promise<string> {
		return this.#donationClient
			.createSubscriptionCheckoutSession({ planId })
			.then((res) => res.checkoutSessionId)
			.catch(ConvertAPIError);
	}

	async getSubscriptionPlans(): Promise<Plan[]> {
		return this.#donationClient
			.getSubscriptionPlans({})
			.then((res) => res.subscriptionPlans.map(fromPBPlan))
			.catch(ConvertAPIError);
	}

	async getActiveSubscription(): Promise<Subscription> {
		return this.#donationClient
			.getActiveSubscription({})
			.then((res) => fromPBSubscription(assurePresence(res.subscription)))
			.catch(ConvertAPIError);
	}

	async cancelSubscription(id: string): Promise<void> {
		return this.#donationClient
			.unsubscribe({ id })
			.then(() => {
				return;
			})
			.catch(ConvertAPIError);
	}

	async getPayments(): Promise<Payment[]> {
		return this.#donationClient
			.getPaymentHistories({})
			.then((res) => res.paymentHistories.map(fromPBPayment))
			.catch(ConvertAPIError);
	}

	async checkAuthentication(): Promise<boolean> {
		return this.#authClient
			.getMe({})
			.then(() => true)
			.catch(ConvertAPIError)
			.catch((error) => {
				if (isUnauthenticatedError(error)) {
					return false;
				}
				throw error;
			});
	}
}

const transport = createConnectTransport({
	baseUrl: process.env.NEXT_PUBLIC_API_BASE_URL as string,
	credentials: 'include',
	useHttpGet: true
});

export const useCase = new UseCase(transport);
