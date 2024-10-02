import type { Dayjs } from "dayjs";

export type PaymentHistory = {
	id: string;
	type: PaymentType;
	status: PaymentStatus;
	amount: number;
	createdAt: Dayjs;
};

export type PaymentStatus = "Succeeded" | "Canceled" | "Pending";

export type PaymentType = "OneTime" | "Subscription";

export const PaymentTypeMap: { [key in PaymentType]: string } = {
	OneTime: "単発",
	Subscription: "サブスク",
};
