import axios from 'axios';

export const apiClient = axios.create({
	baseURL: 'https://app.twinte.net/api/v3',
	withCredentials: true,
	responseType: 'json',
	headers: {
		'Content-Type': 'application/json'
	}
});
