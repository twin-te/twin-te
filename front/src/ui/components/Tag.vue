<script setup lang="ts">
import { computed } from "vue";
import { TagColor } from "~/domain/tag";

const props = defineProps<{
  tag: {
    id: string;
    name: string;
    assign: boolean;
    color: TagColor | null;
  };
}>();

defineEmits<{ click: [] }>();

const color = computed(() =>
  props.tag.color ? `var(--tag-${props.tag.color})` : null
);
</script>

<template>
  <div
    :class="['tag', tag.assign && '--assigned']"
    :style="{
      backgroundColor: tag.assign ? color : null,
      borderColor: color,
    }"
    @click="$emit('click')"
  >
    {{ tag.name }}
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles" as *;
.tag {
  width: max-content;
  height: 2rem;

  padding: 0.3rem 0.7rem;
  @include center-flex;

  font-size: $font-small;
  color: getColor(--color-text-main);
  background-color: getColor(--base);

  border: 0.2rem solid getColor(--color-unselected);
  border-radius: $radius-1;

  @include button-cursor;

  &.--assigned {
    background-color: var(--primary-liner);
  }
}
</style>
