<script setup lang="ts">
import { computed } from "vue";
import { TagColor } from "~/domain/tag";
import { tagColorToCss } from "~/presentation/presenters/tag";

const props = defineProps<{
  tag: {
    id: string;
    name: string;
    assign: boolean;
    color: TagColor | null;
  };
}>();

defineEmits<{ click: [] }>();

const style = computed(() => ({
  "--color": tagColorToCss(props.tag.color),
  borderColor: props.tag.color == null ? undefined : "var(--color)",
}));
</script>

<template>
  <div
    :class="['tag', tag.assign && '--assigned']"
    :style="style"
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
    background-color: oklch(from var(--color) 0.91 calc(c * 0.4) h);
  }
}
:global(.dark .tag.--assigned) {
  background: oklch(from var(--color) 0.33 calc(c * 0.4) h) !important;
}
</style>
