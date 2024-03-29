export type Payment = {
	id: string;
	type: PaymentType;
	status: 'Succeeded' | 'Canceled' | 'Pending';
	amount: number;
	created: string;
};

type PaymentType = 'OneTime' | 'Subscription';

export const PaymentTypeMap: { [key in PaymentType]: string } = {
	OneTime: '単発',
	Subscription: 'サブスク'
};
