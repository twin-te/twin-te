<template>
  <!-- <ErrorBoundary> -->
    <SuspenseBoundary>
      <InitializeApp>
        <Layout>
          <router-view />
        </Layout>
      </InitializeApp>
    </SuspenseBoundary>
  <!-- </ErrorBoundary> -->
</template>

<script setup lang="ts">
import Layout from "~/ui/templates/Layout.vue";
import ErrorBoundary from "./ErrorBoundary.vue";
import InitializeApp from "./InitializeApp.vue";
import SuspenseBoundary from "./SuspenseBoundary.vue";
</script>

<style scoped lang="scss">
@import "~/ui/styles";

.loading {
  @include center-flex;
  width: 100%;
  height: 100vh;
  background: var(--base-liner);
  color: var(--text-sub-light);
}
</style>
