<script setup lang="ts">
import { DisplayCourseTag } from "~/presentation/viewmodels/tag";
import Tag from "~/ui/components/Tag.vue";

defineProps<{
  heading: string;
  tags: DisplayCourseTag[];
}>();

defineEmits<{
  "click:tag": [DisplayCourseTag];
}>();
</script>

<template>
  <div ref="tagEditorRef" class="tag-editor">
    <p class="tag-editor__heading">{{ heading }}</p>
    <section class="tag-editor__contents">
      <div class="tag-editor__tags">
        <template v-if="tags.length === 0">
          作成済みのタグがありません。<br />
          タグを作成すると授業を分類することができます。
        </template>
        <Tag
          v-for="tag in tags"
          v-else
          :key="tag.id"
          :tag="tag"
          @click="() => $emit('click:tag', tag)"
        />
      </div>
    </section>
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles" as *;
@use "~/ui/styles/variable" as *;

.tag-editor {
  width: 100%;
  display: grid;
  row-gap: $spacing-3;

  padding: $spacing-2;

  &__heading {
    font-size: $font-small;
    color: getColor(--color-text-sub);
  }
  &__contents {
    display: flex;
    flex-direction: column;
    gap: $spacing-2;
  }
  &__tags {
    width: 100%;

    display: flex;
    flex-wrap: wrap;
    justify-content: left;
    gap: $spacing-3 $spacing-2;

    color: getColor(--color-disabled);
    font-size: $font-small;
    line-height: $multi-line;
  }
}
</style>
