<template>
  <div v-if="isFutureYear(appliedYear)" class="future-year-banner">
    <span class="material-icons future-year-banner__icon">info</span>
    <span class="future-year-banner__text">
      【ベータ版】{{ currentAcademicYear }}年度のシラバスを使用しています。
    </span>
  </div>
</template>

<script setup lang="ts">
import { currentAcademicYear, isFutureYear } from "~/domain/year";
import { useSetting } from "~/ui/store";

const { appliedYear } = useSetting();
</script>

<style lang="scss" scoped>
@import "~/ui/styles";

.future-year-banner {
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
