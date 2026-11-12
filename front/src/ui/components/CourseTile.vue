<script setup lang="ts">
import { computed, PropType } from "vue";
import { TagColor } from "~/domain/tag";
import { tagColorToCss } from "~/presentation/presenters/tag";

export type State = "default" | "none";

const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  room: {
    type: String,
    required: true,
  },
  color: {
    type: String as PropType<TagColor | null>,
    default: null,
  },
  state: {
    type: String as PropType<State>,
    required: true,
    validator: function (value: string) {
      return ["default", "none"].includes(value);
    },
  },
  caution: {
    type: String,
    default: "", // 空欄の場合 caution は表示されない
  },
});

const emit = defineEmits<{
  click: [MouseEvent];
}>();

const handleClick = (e: MouseEvent) => {
  emit("click", e);
};

const hasCaution = computed(() => {
  return props.caution !== "";
});

const style = computed(() =>
  props.state === "none"
    ? {}
    : {
        "--color": tagColorToCss(props.color),
      }
);
</script>

<template>
  <div
    :class="{
      tile: true,
      [`--${state}`]: true,
      [`--under-filter`]: hasCaution,
      'default-color': color == null,
    }"
    :style="style"
    @click="handleClick"
  >
    <div class="tile__course-name">{{ name }}</div>
    <div class="tile__course-room">{{ room }}</div>
    <div
      v-show="hasCaution"
      :class="{
        tile__caution: true,
      }"
    >
      {{ caution }}
    </div>
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles/variable" as *;
@use "~/ui/styles/mixin" as *;

.tile {
  @include button-cursor;
  position: relative;
  border: solid transparent 0.2rem;
  padding: 0.3rem 0.4rem;
  border-radius: $radius-1;
  text-align: left;
  transition: $transition-box-shadow;
  overflow: hidden;
  &.--default {
    background-color: oklch(from var(--color) 0.91 calc(c * 0.4) h);

    &.default-color {
      background-color: getColor(--color-primary-light) !important;
    }

    &:active {
      box-shadow: $shadow-tile-concave;
    }
  }
  &.--none {
    background-color: getColor(--color-undefined);
  }
  &.--under-filter::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border-radius: $radius-1;
    background-color: getColor(--color-filter-darken);
  }
  &__course-name {
    line-height: $single-line;
    font-size: $font-minimum;
    font-weight: 500;
    color: getColor(--color-text-main);
  }
  &__course-room {
    @include text-course-tile-id;
    overflow: hidden;
    white-space: nowrap;
    width: 100%;
    text-overflow: ellipsis;
  }
  &__caution {
    position: absolute;
    z-index: 10;
    left: $spacing-1;
    bottom: $spacing-1;
    width: calc(100% - 0.8rem);
    padding: $spacing-1 0;
    border-radius: $radius-1;
    background-color: getColor(--color-base);
    text-align: center;
    line-height: $fit;
    font-size: $font-minimum;
    font-weight: 500;
    color: getColor(--color-primary-dull);
  }
}
:global(.dark .tile.--default) {
  background-color: oklch(from var(--color) 0.33 calc(c * 0.4) h) !important;
}
:global(.dark .tile.--default.default-color) {
  background-color: getColor(--color-primary-dark) !important;
}
</style>
