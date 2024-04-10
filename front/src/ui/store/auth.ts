import { computed, ref } from "vue";
import { isResultError } from "~/domain/error";
import { authUseCase } from "~/usecases";

const isAuthenticated = ref<boolean>(false);

const capturedUnauthenticatedError = () => {
  isAuthenticated.value = false;
};

const initializeAuth = () => {
  return authUseCase.checkAuthentication().then((result) => {
    if (isResultError(result)) throw result;
    isAuthenticated.value = result;
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
