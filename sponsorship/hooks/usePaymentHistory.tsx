import { useEffect, useState } from 'react';
import { Payment } from '../domain';
import { useCase } from '../usecases';

export const usePaymentHistory = () => {
	// undefined: API呼び出しが完了していない状態
	// null: API呼び出し中にエラーが発生した状態
	const [paymentHistory, setPaymentHistory] = useState<undefined | null | [Payment]>(undefined);

	useEffect(() => {
		(async () => {
			try {
				const payments = await useCase.getPayments();
				setPaymentHistory(payments as [Payment]);
			} catch (error) {
				console.error(error);
				setPaymentHistory(null);
			}
		})();
	}, []);
	return paymentHistory;
};
