import { computed, ref } from "vue";
import { UnauthenticatedError, isResultError } from "~/domain/error";
import { authUseCase } from "~/usecases";

const isAuthenticated = ref<boolean>(false);

const capturedUnauthenticatedError = () => {
	isAuthenticated.value = false;
};

const initializeAuth = async () => {
	return authUseCase.getMe().then((result) => {
		if (!isResultError(result)) {
			isAuthenticated.value = true;
			return;
		}

		if (result instanceof UnauthenticatedError) {
			isAuthenticated.value = false;
			return;
		}

		throw result;
	});
};

const useAuth = () => {
	return {
		isAuthenticated: computed(() => isAuthenticated.value),
		capturedUnauthenticatedError,
		initializeAuth,
	};
};

export default useAuth;
