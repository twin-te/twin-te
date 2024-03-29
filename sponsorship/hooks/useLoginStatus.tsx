import { useEffect, useState } from 'react';
import { apiClient } from '../api/apiClient';

export const useLoginStatus = () => {
	// undefined: ログイン状態を確認できていなdい
	// boolean: ログイン / ログアウト状態
	const [isLogin, setIsLogin] = useState<undefined | boolean>(undefined);

	useEffect(() => {
		(async () => {
			try {
				await apiClient.get('/users/me');
				setIsLogin(true);
			} catch (error) {
				setIsLogin(false);
			}
		})();
	}, []);
	return isLogin;
};
