<template>
  <Suspense>
    <template #default>
      <slot />
    </template>
    <template #fallback>
      <div class="loading">
        <p class="loading__text">Now loading...</p>
      </div>
    </template>
  </Suspense>
</template>

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
