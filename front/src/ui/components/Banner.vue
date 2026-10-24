<template>
  <div class="banner">
    <span class="material-icons banner__icon">{{ icon }}</span>
    <span class="banner__text"><slot /></span>
  </div>
</template>

<script setup lang="ts">
withDefaults(defineProps<{ icon?: string }>(), { icon: "info" });
</script>

<style lang="scss" scoped>
@use "~/ui/styles/variable" as *;
@use "~/ui/styles/mixin" as *;

.banner {
  display: flex;
  align-items: center;
  gap: $spacing-1;
  background: getColor(--color-primary-light);
  color: getColor(--color-text-main);
  padding: $spacing-2;
  border-radius: $radius-3;
  font-size: $font-small;
  margin: $spacing-2 0 0;

  @include portrait {
    margin-left: (-$spacing-4);
    margin-right: (-$spacing-4);
    border-radius: 0;
    padding-left: $spacing-4;
    padding-right: $spacing-4;
  }

  &__icon {
    font-size: 1.6rem;
    line-height: 1;
    flex-shrink: 0;
  }

  &__text {
    text-align: left;
  }
}
</style>
