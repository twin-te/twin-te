<template>
  <slot />
</template>

<script setup lang="ts">
import * as Sentry from "@sentry/vue";
import { onErrorCaptured, watch } from "vue";
import { UnauthenticatedError, isResultError } from "~/domain/error";
import { authUseCase } from "~/usecases";
import { useAuth, useNews, useSetting } from "./store";

const {
  isAuthenticated,
  capturedUnauthenticatedError,
  initializeAuth,
} = useAuth();
const { initializeNews } = useNews();
const { initializeSetting } = useSetting();

await initializeAuth();
await Promise.all([initializeNews(), initializeSetting()]);

watch(isAuthenticated, () => {
  if (isAuthenticated.value) {
    authUseCase
      .getUser()
      .then((user) => {
        if (isResultError(user)) return;
        Sentry.setUser(user);
      })
      .catch(() => Sentry.setUser(null));
  } else {
    Sentry.setUser(null);
  }
});

onErrorCaptured((error) => {
  if (error instanceof UnauthenticatedError) {
    capturedUnauthenticatedError();
  }
  return true;
});
</script>
