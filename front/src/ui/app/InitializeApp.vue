<template>
  <slot />
</template>

<script setup lang="ts">
import * as Sentry from "@sentry/vue";
import { onErrorCaptured, watch } from "vue";
import { UnauthenticatedError, isResultError } from "~/domain/error";
import { authUseCase } from "~/usecases";
import { useAuth, useSetting } from "../store";

const {
  isAuthenticated,
  capturedUnauthenticatedError,
  initializeAuth,
} = useAuth();
const { initializeSetting } = useSetting();

await initializeAuth();
await initializeSetting();

watch(isAuthenticated, () => {
  authUseCase.getMe().then((result) => {
    if (!isResultError(result)) Sentry.setUser(result);
    else if (result instanceof UnauthenticatedError) Sentry.setUser(null);
    else throw result;
  });
});

onErrorCaptured((error) => {
  if (error instanceof UnauthenticatedError) {
    capturedUnauthenticatedError();
  }
  return true;
});
</script>
