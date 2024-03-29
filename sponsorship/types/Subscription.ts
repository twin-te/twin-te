type Plan = {
	id: string;
	name: string;
	amount: number;
};

export type Subscription = {
	id: string;
	status: 'Active' | 'Canceled';
	plans: [Plan];
	created: string;
};
