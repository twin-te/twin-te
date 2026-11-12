<script setup lang="ts">
import { computed } from "vue";
import { DisplayCourseTag } from "~/presentation/viewmodels/tag";
import Tag from "./Tag.vue";
import TagList from "./TagList.vue";

export type CreditCourseListContentState = "default" | "selected";

const props = defineProps<{
  selected: boolean;
  code: string;
  name: string;
  credit: string;
  tags: DisplayCourseTag[];
}>();

defineEmits<{
  click: [];
  "click-tag": [DisplayCourseTag];
}>();

const assignedTags = computed(() => props.tags.filter((tag) => tag.assign));
</script>

<template>
  <div
    :class="{
      'credit-course-list-content': true,
      '--selected': selected,
    }"
  >
    <div
      class="credit-course-list-content__course-info course-info"
      @click="$emit('click')"
    >
      <div class="course-info__container">
        <div class="course-info__code">{{ code }}</div>
        <div class="course-info__name">{{ name }}</div>
      </div>
      <div v-show="!selected" class="course-info__tags">
        <Tag v-for="tag in assignedTags" :key="tag.id" :tag="tag" />
      </div>
      <div class="course-info__credit">{{ credit }}</div>
      <div
        :class="{
          'course-info__expand': true,
          '--turned': selected,
          'material-icons': true,
        }"
      >
        expand_more
      </div>
    </div>
    <TagList
      v-show="selected"
      heading="タグ"
      :tags="tags"
      @click:tag="$emit('click-tag', $event)"
    />
    <div class="credit-course-list-content__border" />
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles" as *;
@use "~/ui/styles/variable" as *;

.credit-course-list-content {
  width: 100%;

  &__border {
    width: 100%;
    height: 0.4rem;

    background: getColor(--color-base);
    box-shadow: $shadow-input-concave;
    border-radius: 0.2rem;
  }
}

.course-info {
  display: flex;
  align-items: center;
  gap: $spacing-2;

  padding: $spacing-3 $spacing-2 $spacing-2;

  @include button-cursor;

  &__container {
    flex-grow: 1;
  }

  &__code {
    color: getColor(--color-text-sub);
    font-size: $font-small;
  }

  &__name {
    margin-top: 0.2rem;
    font-weight: 500;
    line-height: $multi-line;
  }

  &__tags {
    display: flex;
    flex-direction: column;
    gap: $spacing-2;
    align-items: flex-end;
  }

  &__credit {
    font-weight: 500;
  }

  &__expand {
    color: getColor(--color-text-sub);
    font-size: $font-large;

    transition: $transition-transform;
    &.--turned {
      transform: rotate(-180deg);
    }
  }
}
</style>
