import * as DonationV1PB from '../../api/gen/donation/v1/type_pb';
import { Subscription, User, Plan, Payment } from '../../domain';
import { PaymentStatus, PaymentType } from '../../domain/Payment';
import { fromPBRFC3339DateTime, fromPBUUID } from './shared';
import { assurePresence } from './utils';

export const fromPBPaymentUser = (pbPaymentUser: DonationV1PB.PaymentUser): User => {
	return {
		twinteUserId: pbPaymentUser.id,
		paymentUserId: fromPBUUID(assurePresence(pbPaymentUser.userId)),
		displayName: pbPaymentUser.displayName,
		link: pbPaymentUser.link
	};
};

export const fromPBSubscription = (pbSubscription: DonationV1PB.Subscription): Subscription => {
	return {
		id: pbSubscription.id,
		plan: fromPBPlan(assurePresence(pbSubscription.plan)),
		status: pbSubscription.isActive ? 'Active' : 'Canceled',
		createdAt: fromPBRFC3339DateTime(assurePresence(pbSubscription.createdAt))
	};
};

export const fromPBPlan = (pbPlan: DonationV1PB.SubscriptionPlan): Plan => {
	return {
		id: pbPlan.id,
		name: pbPlan.name,
		amount: pbPlan.amount
	};
};

export const fromPBPaymentType = (pbPaymentType: DonationV1PB.PaymentType): PaymentType => {
	switch (pbPaymentType) {
		case DonationV1PB.PaymentType.ONE_TIME:
			return 'OneTime';
		case DonationV1PB.PaymentType.SUBSCRIPTION:
			return 'Subscription';
	}
	throw new Error(`invalid enum ${pbPaymentType}`);
};

export const fromPBPaymentStatus = (pbPaymentStatus: DonationV1PB.PaymentStatus): PaymentStatus => {
	switch (pbPaymentStatus) {
		case DonationV1PB.PaymentStatus.PENDING:
			return 'Pending';
		case DonationV1PB.PaymentStatus.CANCELED:
			return 'Canceled';
		case DonationV1PB.PaymentStatus.SUCCEEDED:
			return 'Succeeded';
	}
	throw new Error(`invalid enum ${pbPaymentStatus}`);
};

export const fromPBPayment = (pbPayment: DonationV1PB.PaymentHistory): Payment => {
	return {
		id: pbPayment.id,
		type: fromPBPaymentType(pbPayment.type),
		status: fromPBPaymentStatus(pbPayment.status),
		amount: pbPayment.amount,
		createdAt: fromPBRFC3339DateTime(assurePresence(pbPayment.createdAt))
	};
};
