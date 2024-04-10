import { useEffect, useState } from 'react';
import { Plan, Subscription } from '../domain';
import { useCase } from '../usecases';
import { isNotFoundError } from '../usecases/error';

// export const useSubscription = () => {
// 	// undefined: API呼び出しが完了していない状態
// 	// null: API呼び出し中にエラーが発生した状態
// 	const [subscription, setSubscription] = useState<undefined | null | Subscription>(undefined);

// 	useEffect(() => {
// 		useCase
// 			.getSubscription()
// 			.then((subscription) => setSubscription(subscription))
// 			.catch((error) => {
// 				console.error(error);
// 				setSubscription(null);
// 			});
// 	}, []);

// 	return subscription;
// };

type State = {
	plans: Plan[];
	selectedPlanID: string;
	isLoading: boolean;
};

export const usePlans = () => {
	const [plans, setPlans] = useState<State>({ plans: [], selectedPlanID: "", isLoading: true });

	useEffect(() => {
		useCase
			.getSubscriptionPlans()
			.then((plans) => {
				if (plans.length == 0) {
					throw new Error("not found subscription plan")
				}
				setPlans({ plans: plans, selectedPlanID: plans[0].id, isLoading: false, })
			}).catch((error) => {
				setPlans({plans: [], selectedPlanID: "", isLoading: false})
				throw error
			})
	}, []);

	const selectPlanID = (id: string) => {
		setPlans({ plans: plans.plans, selectedPlanID: id, isLoading: false, })
	}

	return { plans, selectPlanID };
};
