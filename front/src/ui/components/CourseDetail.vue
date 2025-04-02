<script setup lang="ts">
defineProps<{
  item: string
  value: string
}>()
defineSlots<{
  default(): unknown
  value(props: { value: string}): unknown
}>()
</script>

<template>
  <div class="course-detail">
    <div class="course-detail__icon">
      <slot></slot>
    </div>
    <div class="course-detail__item">{{ item }}</div>
    <div class="course-detail__value">
      <slot name="value">{{ value }}</slot>
    </div>
  </div>
</template>

<style scoped lang="scss">
@import "~/ui/styles";

.course-detail {
  display: grid;
  grid-template:
    "icon ... item " auto
    "icon ... ...  " 1fr
    "icon ... value" auto
    / auto $spacing-4 1fr;
  text-align: left;
  &__icon {
    grid-area: icon;
  }
  &__item {
    grid-area: item;
    line-height: $single-line;
    font-size: $font-small;
    color: getColor(--color-text-sub);
  }
  &__value {
    grid-area: value;
    line-height: $single-line;
    font-size: $font-medium;
    font-weight: 500;
    color: getColor(--color-text-main);
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }
}
</style>
