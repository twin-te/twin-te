import { useEffect, useState } from 'react';
import { Subscription } from '../domain';
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
	value?: Subscription;
	isLoading: boolean;
	failed: boolean;
};

export const useActiveSubscription = () => {
	const [activeSubscription, setActiveSubscription] = useState<State>({ isLoading: true, failed: false });

	useEffect(() => {
		useCase
			.getActiveSubscription()
			.then((subscription) => setActiveSubscription({ value: subscription, isLoading: false, failed: false }))
			.catch((error) => {
				if (isNotFoundError(error)) {
					setActiveSubscription({ isLoading: false, failed: false });
				} else {
					setActiveSubscription({ isLoading: false, failed: true });
				}
			});
	}, []);

	return activeSubscription;
};
