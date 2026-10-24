<script setup lang="ts">
withDefaults(
  defineProps<{
    color?: "ghost" | "primary" | "danger";
    iconPosition?: "left" | "right";
  }>(),
  {
    color: "primary",
    iconPosition: "left",
  }
);
defineEmits<{ click: [] }>();
defineSlots<{
  icon: () => unknown;
  text: () => unknown;
}>();
</script>

<template>
  <div
    :class="{ 'tertiary-button': true, [`tertiary-button--${color}`]: true }"
    :style="{ flexDirection: iconPosition === 'left' ? 'row' : 'row-reverse' }"
    @click="$emit('click')"
  >
    <div class="tertiary-button__icon material-icons">
      <slot name="icon"></slot>
    </div>
    <div class="tertiary-button__text">
      <slot name="text"></slot>
    </div>
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles" as *;

.tertiary-button {
  display: flex;
  gap: $spacing-1;
  width: max-content;
  padding: $spacing-2 $spacing-1;
  border-radius: $radius-1;
  font-size: $font-small;
  transition: $transition-box-shadow;

  &__icon {
    font-weight: 400;
  }

  &__text {
    font-weight: 500;
  }

  &--ghost {
    color: getColor(--color-button-gray);
  }

  &--primary {
    color: getColor(--color-primary);
  }

  &--danger {
    color: getColor(--color-danger);
  }

  &:hover {
    box-shadow: $shadow-convex-hover;
  }

  &:active {
    box-shadow: $shadow-concave;
  }

  @include button-cursor;
}
</style>
