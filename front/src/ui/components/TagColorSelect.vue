<script setup lang="ts">
import {PresetColor, TagColor} from "~/domain/tag";

const model = defineModel<TagColor | null>();
const colors = [
  null,
  "pink",
  "sky",
  "mint",
  "peach",
  "lilac",
  "ivory",
] as const satisfies (PresetColor | null)[];

const emits = defineEmits<{
  change: [TagColor | null];
}>();

const onClick = (color: TagColor | null) => {
  model.value = color;
  if (color !== model.value) emits('change', color);
};
</script>

<template>
  <div class="color-select">
    <div
      v-for="color in colors"
      :key="color ?? 'default'"
      class="color-rect"
      :style="{ background: `var(--tag-${color ?? 'default'})` }"
      @click="onClick(color)"
    >
      <div v-if="color === model" class="material-icons check">check</div>
    </div>
  </div>
</template>

<style scoped lang="scss">
@use '~/ui/styles/variable';

.color-select {
  --tag-default: rgba(var(--color-primary-light));
  display: flex;
  gap: variable.$spacing-3;
  align-items: center;
}

.color-rect {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 3rem;
  height: 3rem;
  cursor: pointer;

  border: solid 1px rgba(var(--color-text-sub-light));
  border-radius: 0.4rem;
  user-select: none;

  .check {
    color: rgba(var(--color-text-main));
    font-size: 2rem;
  }
}
</style>
