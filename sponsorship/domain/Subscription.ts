import { Dayjs } from 'dayjs';

export type Plan = {
	id: string;
	name: string;
	amount: number;
};

export type Subscription = {
	id: string;
	plan: Plan;
	status: 'Active' | 'Canceled';
	createdAt: Dayjs;
};
