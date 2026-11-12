<script setup lang="ts">
import { computed } from "vue";
import { TagColor } from "~/domain/tag";
import { tagColorToCss } from "~/presentation/presenters/tag";

const props = defineProps<{
  tags: { id: string; name: string; color: TagColor | null }[];
  selectedId?: string;
}>();

defineEmits<{
  click: [string | undefined];
}>();

const allTags = computed(() => [
  ...props.tags,

  {
    id: "all-courses",
    name: "すべての授業",
    color: null,
  },
]);
</script>

<template>
  <ul class="tag-list">
    <li
      v-for="tag in allTags"
      :key="tag.id"
      :class="['tag-list-item', tag.id === selectedId && 'selected']"
      @click="$emit('click', tag.id)"
    >
      <div
        class="tag-list-item__color-chip"
        :style="{
          backgroundColor: tagColorToCss(tag.color),
        }"
      />
      <span>{{ tag.name }}</span>
    </li>
  </ul>
</template>

<style scoped lang="scss">
@use "sass:color";
@use "sass:map";
@use "~/ui/styles/variable";

.tag-list {
  display: flex;
  flex-direction: column;
  gap: variable.$spacing-1;
}

.tag-list-item {
  display: flex;
  width: 16rem;
  border-radius: 4px;
  align-items: center;
  justify-content: flex-start;
  cursor: pointer;
  padding: variable.$spacing-2 variable.$spacing-3;

  &:hover {
    background-color: rgba(var(--color-primary-light));
  }

  &.selected {
    color: rgba(var(--color-white));
    background-color: rgba(var(--color-primary-dull));
  }

  &__color-chip {
    width: 1.3rem;
    height: 1.3rem;
    border-radius: 50%;
    background-color: rgba(var(--color-primary-light));
    margin-inline-end: 0.6rem;
  }
}
</style>
