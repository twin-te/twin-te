import { useEffect, useState } from 'react';
import { apiClient } from '../api/apiClient';
import { Subscription } from '../types';

export const useSubscriptions = () => {
	// undefined: API呼び出しが完了していない状態
	// null: API呼び出し中にエラーが発生した状態
	const [subscriptions, setSubscriptions] = useState<undefined | null | [Subscription]>(undefined);

	useEffect(() => {
		(async () => {
			try {
				const res = await apiClient.get('/donation/subscriptions');
				setSubscriptions(res.data);
			} catch (error) {
				console.error(error);
				setSubscriptions(null);
			}
		})();
	}, []);
	return subscriptions;
};
