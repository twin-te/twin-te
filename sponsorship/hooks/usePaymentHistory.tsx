import { useEffect, useState } from 'react';
import { apiClient } from '../api/apiClient';
import { Payment } from '../types';

export const usePaymentHistory = () => {
	// undefined: API呼び出しが完了していない状態
	// null: API呼び出し中にエラーが発生した状態
	const [paymentHistory, setPaymentHistory] = useState<undefined | null | [Payment]>(undefined);

	useEffect(() => {
		(async () => {
			try {
				const res = await apiClient.get('/donation/payment');
				setPaymentHistory(res.data);
			} catch (error) {
				console.error(error);
				setPaymentHistory(null);
			}
		})();
	}, []);
	return paymentHistory;
};
