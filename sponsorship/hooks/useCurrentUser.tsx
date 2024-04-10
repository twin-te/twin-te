import { Dispatch, SetStateAction, useEffect, useState } from 'react';
import { User } from '../domain';
import { useCase } from '../usecases';

export const useCurrentUser = (): [User | null | undefined, Dispatch<SetStateAction<User | null | undefined>>] => {
	// undefined: API呼び出しが完了していない状態
	// null: API呼び出し中にエラーが発生した状態
	const [currentUser, setCurrentUser] = useState<undefined | null | User>(undefined);

	useEffect(() => {
		(async () => {
			try {
				const user = await useCase.getUser();
				setCurrentUser(user);
			} catch (error) {
				console.error(error);
				setCurrentUser(null);
			}
		})();
	}, []);
	return [currentUser, setCurrentUser];
};
