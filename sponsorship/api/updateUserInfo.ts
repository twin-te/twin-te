import { User } from '../types';
import { apiClient } from './apiClient';

export const updateUserInfo = async (newDisplayName: string, newLink: string) => {
	const res = await apiClient.patch<User>(`/donation/users/me`, {
		displayName: newDisplayName,
		link: newLink
	});
	return res.data;
};
