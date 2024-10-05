import * as DonationV1PB from '../../api/gen/donation/v1/type_pb';
import { Subscription, User, SubscriptionPlan, PaymentHistory } from '../../domain';
import { PaymentStatus, PaymentType } from '../../domain/payment_history';
import { fromPBRFC3339DateTime, fromPBUUID } from './shared';
import { assurePresence } from './utils';

export const fromPBPaymentUser = (pbPaymentUser: DonationV1PB.PaymentUser): User => {
	return {
		twinteUserId: fromPBUUID(assurePresence(pbPaymentUser.userId)),
		paymentUserId: pbPaymentUser.id,
		displayName: pbPaymentUser.displayName,
		link: pbPaymentUser.link
	};
};

export const fromPBSubscription = (pbSubscription: DonationV1PB.Subscription): Subscription => {
	return {
		id: pbSubscription.id,
		plan: fromPBPlan(assurePresence(pbSubscription.plan)),
		isActive: pbSubscription.isActive,
		createdAt: fromPBRFC3339DateTime(assurePresence(pbSubscription.createdAt))
	};
};

export const fromPBPlan = (pbPlan: DonationV1PB.SubscriptionPlan): SubscriptionPlan => {
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

export const fromPBPaymentHistory = (pbPaymentHistory: DonationV1PB.PaymentHistory): PaymentHistory => {
	return {
		id: pbPaymentHistory.id,
		type: fromPBPaymentType(pbPaymentHistory.type),
		status: fromPBPaymentStatus(pbPaymentHistory.status),
		amount: pbPaymentHistory.amount,
		createdAt: fromPBRFC3339DateTime(assurePresence(pbPaymentHistory.createdAt))
	};
};
