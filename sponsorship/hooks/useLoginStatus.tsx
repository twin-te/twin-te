import { useEffect, useState } from 'react';
import { useCase } from '../usecases';

export const useLoginStatus = () => {
	// undefined: ログイン状態を確認できていなdい
	// boolean: ログイン / ログアウト状態
	const [isLogin, setIsLogin] = useState<undefined | boolean>(undefined);

	useEffect(() => {
		(async () => {
			const isAuthenticated = await useCase.checkAuthentication();
			setIsLogin(isAuthenticated);
		})();
	}, []);
	return isLogin;
};
