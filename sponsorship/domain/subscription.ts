import type { Dayjs } from "dayjs";

export type SubscriptionPlan = {
	id: string;
	name: string;
	amount: number;
};

export type Subscription = {
	id: string;
	plan: SubscriptionPlan;
	isActive: boolean;
	createdAt: Dayjs;
};
