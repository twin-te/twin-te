<script setup lang="ts">
import { computed, ref } from "vue";
import { icalUrlWithTags } from "~/presentation/presenters/calendar";
import { useToast } from "../store";
import Button from "./Button.vue";
import Checkbox from "./Checkbox.vue";
import Modal from "./Modal.vue";
import Tag from "./Tag.vue";
import TertiaryButton from "./TertiaryButton.vue";

export type IcalTagOption = {
  id: string;
  name: string;
  courseCount: number;
};

const props = defineProps<{
  /** The iCal subscription URL, which exports all the registered courses. */
  url: string;
  /** Sorted in display order. */
  tags: IcalTagOption[];
  initialSelectedTagIds: string[];
}>();

const emit = defineEmits<{
  close: [];
  issue: [tagIds: string[]];
}>();

const { displayToast } = useToast();

const step = ref<"select" | "issued">("select");

/** select step */
const selectedTagIds = ref<string[]>(
  props.initialSelectedTagIds.filter((id) =>
    props.tags.some((tag) => tag.id === id)
  )
);

const isSelected = (tagId: string) => selectedTagIds.value.includes(tagId);

const toggleTag = (tagId: string) => {
  selectedTagIds.value = isSelected(tagId)
    ? selectedTagIds.value.filter((id) => id !== tagId)
    : [...selectedTagIds.value, tagId];
};

const isAllSelected = computed(
  () =>
    props.tags.length > 0 && selectedTagIds.value.length === props.tags.length
);

const toggleAll = () => {
  selectedTagIds.value = isAllSelected.value
    ? []
    : props.tags.map((tag) => tag.id);
};

/** issued step */
// Keep the display order so that the same selection always gives the same URL.
const issuedTags = computed(() =>
  props.tags.filter((tag) => isSelected(tag.id))
);

const issuedUrl = computed(() =>
  icalUrlWithTags(
    props.url,
    issuedTags.value.map((tag) => tag.id)
  )
);

const copied = ref(false);

const issue = () => {
  if (selectedTagIds.value.length === 0) return;
  copied.value = false;
  step.value = "issued";
  emit(
    "issue",
    issuedTags.value.map((tag) => tag.id)
  );
};

const copyIssuedUrl = async () => {
  try {
    await navigator.clipboard.writeText(issuedUrl.value);
    copied.value = true;
  } catch {
    displayToast("コピーに失敗しました", { type: "danger" });
  }
};

const backToSelect = () => {
  step.value = "select";
};
</script>

<template>
  <Modal
    v-if="step === 'select'"
    class="ical-tag-modal"
    size="large"
    @click="emit('close')"
  >
    <template #title>対象のタグをカスタマイズ</template>
    <template #contents>
      <div class="ical-tag-modal__contents">
        <p class="ical-tag-modal__description">
          選択したタグが付いた授業だけに限定したURLを発行します。
        </p>
        <template v-if="tags.length > 0">
          <div class="ical-tag-modal__toolbar">
            <span class="ical-tag-modal__selected-label">
              {{
                selectedTagIds.length > 0
                  ? `${selectedTagIds.length}個のタグを選択中`
                  : "タグを選択してください"
              }}
            </span>
            <TertiaryButton color="ghost" @click="toggleAll">
              <template #icon>
                {{ isAllSelected ? "remove_done" : "done_all" }}
              </template>
              <template #text>
                {{ isAllSelected ? "すべて解除" : "すべて選択" }}
              </template>
            </TertiaryButton>
          </div>
          <div class="ical-tag-modal__tags">
            <div v-for="tag in tags" :key="tag.id" class="ical-tag-row">
              <div class="ical-tag-row__container" @click="toggleTag(tag.id)">
                <div class="ical-tag-row__name">{{ tag.name }}</div>
                <span class="ical-tag-row__count">
                  {{ tag.courseCount }}件
                </span>
                <Checkbox :isChecked="isSelected(tag.id)" />
              </div>
              <div class="ical-tag-row__border" />
            </div>
          </div>
        </template>
        <p v-else class="ical-tag-modal__empty">
          タグがまだありません。授業の詳細画面からタグを作成して、授業に設定してください。
        </p>
      </div>
    </template>
    <template #button>
      <Button
        class="ical-tag-modal__button"
        size="medium"
        layout="fill"
        color="base"
        @click="emit('close')"
        >キャンセル</Button
      >
      <Button
        class="ical-tag-modal__button"
        size="medium"
        layout="fill"
        color="primary"
        :state="selectedTagIds.length > 0 ? 'default' : 'disabled'"
        @click="issue"
        >発行する</Button
      >
    </template>
  </Modal>
  <Modal v-else class="ical-tag-modal" size="large" @click="emit('close')">
    <template #title>URLを発行しました</template>
    <template #contents>
      <div class="ical-tag-modal__contents">
        <p class="ical-tag-modal__description">
          以下のタグが付いた授業だけがカレンダーに同期されます。
        </p>
        <div class="ical-tag-modal__issued-tags">
          <Tag v-for="tag in issuedTags" :key="tag.id" :assign="true">
            {{ tag.name }}
          </Tag>
        </div>
        <div class="ical-tag-modal__url">
          <input
            :value="issuedUrl"
            type="text"
            readonly
            class="ical-tag-modal__url-input"
          />
          <Button
            size="small"
            :color="copied ? 'primary' : 'base'"
            :pauseActiveStyle="false"
            @click="copyIssuedUrl"
            >{{ copied ? "コピー済" : "コピー" }}</Button
          >
        </div>
        <ul class="ical-tag-modal__notes">
          <li>※ 元のURL（すべての授業）はそのまま使えます。</li>
          <li>※ 対象のタグを変更するには、新しいURLを発行してください。</li>
        </ul>
      </div>
    </template>
    <template #button>
      <Button
        class="ical-tag-modal__button"
        size="medium"
        layout="fill"
        color="base"
        @click="backToSelect"
        >選び直す</Button
      >
      <Button
        class="ical-tag-modal__button"
        size="medium"
        layout="fill"
        color="primary"
        @click="emit('close')"
        >閉じる</Button
      >
    </template>
  </Modal>
</template>

<style scoped lang="scss">
@import "~/ui/styles";

.ical-tag-modal {
  &__contents {
    display: flex;
    flex-direction: column;
    height: 100%;
    padding: 0 0.6rem;
  }

  &__description,
  &__empty {
    @include text-description-sub;
  }

  &__empty {
    margin-top: $spacing-4;
  }

  &__toolbar {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: $spacing-4;
    padding: 0 $spacing-2;
  }

  &__selected-label {
    font-size: $font-small;
    color: getColor(--color-text-sub);
  }

  &__tags {
    flex: 1;
    min-height: 0;
    margin-top: $spacing-1;
    padding: $spacing-2 0;
    overflow-x: hidden;
    overflow-y: auto;
    @include scroll-mask;
  }

  &__issued-tags {
    display: flex;
    flex-wrap: wrap;
    gap: $spacing-3 $spacing-2;
    margin-top: $spacing-3;
  }

  &__url {
    display: flex;
    align-items: center;
    gap: $spacing-2;
    height: $spacing-10;
    margin-top: $spacing-6;
    padding: 0 0.6rem 0 1.4rem;
    border-radius: $radius-input;
    background: getColor(--color-base);
    box-shadow: $shadow-input-concave;
  }

  &__url-input {
    flex: 1;
    width: 0;
    font-size: $font-medium;
    font-weight: 500;
    color: getColor(--color-text-main);
    background: transparent;
    text-overflow: ellipsis;
  }

  &__notes {
    margin-top: $spacing-4;
    @include text-description-sub;
    li {
      margin-bottom: $spacing-1;
    }
  }

  &__button + &__button {
    margin-left: $spacing-3;
  }
}

.ical-tag-row {
  width: 100%;

  &__container {
    display: flex;
    align-items: center;
    height: $spacing-11;
    padding: 0 $spacing-2;
    user-select: none;
    @include button-cursor;
  }

  &__name {
    flex-grow: 1;
    font-size: $font-small;
    color: getColor(--color-text-main);
  }

  &__count {
    margin-right: $spacing-4;
    font-size: $font-small;
    color: getColor(--color-text-sub);
  }

  &__border {
    width: 100%;
    height: 0.4rem;
    background: getColor(--color-base);
    box-shadow: $shadow-input-concave;
    border-radius: 0.2rem;
  }
}
</style>
