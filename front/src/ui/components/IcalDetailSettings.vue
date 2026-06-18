<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import { IcalSubscriptionMode } from "~/domain/calendar";
import Tag from "./Tag.vue";

export type IcalDetailSettingsValue = {
  mode: IcalSubscriptionMode;
  targetTagIds: string[];
};

type Behavior = {
  mode: IcalSubscriptionMode;
  title: string;
  description: string;
};

const behaviors: Behavior[] = [
  { mode: "sync", title: "そのまま", description: "通常の予定として表示されます" },
  { mode: "exclude", title: "入れない", description: "予定を表示しません" },
  { mode: "transparent", title: "予定なし", description: "空き時間として表示されます" },
];

export default defineComponent({
  name: "IcalDetailSettings",
  components: { Tag },
  props: {
    modelValue: {
      type: Object as PropType<IcalDetailSettingsValue>,
      required: true,
    },
    tags: {
      type: Array as PropType<{ id: string; name: string }[]>,
      required: true,
    },
  },
  emits: ["update:modelValue"],
  setup(props, { emit }) {
    const selectedTagIds = computed(() => props.modelValue.targetTagIds);
    const mode = computed(() => props.modelValue.mode);

    const isTagSelected = (id: string) => selectedTagIds.value.includes(id);

    const onClickTag = (id: string) => {
      const next = isTagSelected(id)
        ? selectedTagIds.value.filter((tagId) => tagId !== id)
        : [...selectedTagIds.value, id];
      emit("update:modelValue", { mode: mode.value, targetTagIds: next });
    };

    const onSelectMode = (next: IcalSubscriptionMode) => {
      emit("update:modelValue", {
        mode: next,
        targetTagIds: selectedTagIds.value,
      });
    };

    return { behaviors, mode, isTagSelected, onClickTag, onSelectMode };
  },
});
</script>

<template>
  <div class="ical-detail-settings">
    <p class="ical-detail-settings__description">
      タグのついた授業を、連携カレンダーにどう出すか設定できます。
    </p>

    <section class="ical-detail-settings__section">
      <h3 class="ical-detail-settings__heading">対象のタグ</h3>
      <p v-if="tags.length === 0" class="ical-detail-settings__empty">
        タグがまだありません。授業にタグを付けると、ここで選べるようになります。
      </p>
      <div v-else class="ical-detail-settings__tags">
        <Tag
          v-for="tag in tags"
          :key="tag.id"
          :assign="isTagSelected(tag.id)"
          @click="() => onClickTag(tag.id)"
          >{{ tag.name }}</Tag
        >
      </div>
    </section>

    <section class="ical-detail-settings__section">
      <h3 class="ical-detail-settings__heading">対象タグの挙動</h3>
      <div class="ical-detail-settings__cards">
        <button
          v-for="behavior in behaviors"
          :key="behavior.mode"
          type="button"
          :class="{
            card: true,
            '--selected': mode === behavior.mode,
          }"
          @click="() => onSelectMode(behavior.mode)"
        >
          <span v-if="mode === behavior.mode" class="material-icons card__check"
            >check_circle</span
          >
          <div class="card__preview">
            <span class="card__cell"></span>
            <span
              :class="`card__cell card__cell--main card__cell--${behavior.mode}`"
            >
              <template v-if="behavior.mode === 'sync'">授業</template>
              <template v-else-if="behavior.mode === 'transparent'"
                >空き</template
              >
            </span>
            <span class="card__cell"></span>
            <span class="card__cell"></span>
            <span class="card__cell"></span>
            <span class="card__cell"></span>
            <span class="card__cell"></span>
            <span class="card__cell"></span>
          </div>
          <div class="card__labels">
            <span class="card__title">{{ behavior.title }}</span>
            <span class="card__description">{{ behavior.description }}</span>
          </div>
        </button>
      </div>
    </section>
  </div>
</template>

<style scoped lang="scss">
@import "~/ui/styles";

.ical-detail-settings {
  display: flex;
  flex-direction: column;
  gap: $spacing-6;

  &__description {
    @include text-discription;
    color: getColor(--color-text-sub);
    line-height: $single-line;
  }

  &__section {
    display: flex;
    flex-direction: column;
    gap: $spacing-3;
  }

  &__heading {
    font-size: $font-small;
    font-weight: 700;
    letter-spacing: 0.04em;
    color: getColor(--color-text-sub);
  }

  &__empty {
    @include text-discription;
    color: getColor(--color-text-sub);
    line-height: $single-line;
  }

  &__tags {
    display: flex;
    flex-wrap: wrap;
    gap: $spacing-2;
  }

  &__cards {
    display: flex;
    flex-direction: column;
    gap: $spacing-3;
    @include large-screen {
      flex-direction: row;
    }
  }
}

.card {
  position: relative;
  flex: 1;
  display: flex;
  align-items: center;
  gap: $spacing-3;
  padding: $spacing-3 $spacing-4;
  border: none;
  border-radius: $radius-3;
  background: getColor(--color-base);
  box-shadow: $shadow-convex;
  text-align: left;
  @include button-cursor;

  @include large-screen {
    flex-direction: column;
    text-align: center;
    padding: $spacing-4 $spacing-3;
  }

  &.--selected {
    box-shadow: $shadow-primary-concave;
    outline: 0.2rem solid getColor(--color-primary);
  }

  &__check {
    position: absolute;
    top: -0.9rem;
    right: -0.9rem;
    font-size: 2.2rem;
    color: getColor(--color-white);
    background: var(--primary-liner);
    border-radius: 50%;
  }

  &__preview {
    flex-shrink: 0;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: repeat(3, 1.6rem);
    gap: 0.3rem;
    width: 10.8rem;
    padding: 0.6rem;
    border-radius: $radius-2;
    background: getColor(--color-base);
  }

  &__cell {
    border-radius: 0.3rem;
    background: #e2e6ed;

    &--main {
      grid-row: span 2;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 0.7rem;
      font-weight: 700;
    }

    &--sync {
      background: var(--primary-liner);
      color: getColor(--color-white);
    }

    &--exclude {
      background: transparent;
      border: 0.15rem dashed getColor(--color-unselected);
    }

    &--transparent {
      color: getColor(--color-primary-dull);
      background: repeating-linear-gradient(
        45deg,
        #c9eef0,
        #c9eef0 0.3rem,
        #eef7f8 0.3rem,
        #eef7f8 0.6rem
      );
    }
  }

  &__labels {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  &__title {
    font-size: $font-small;
    font-weight: 700;
    color: getColor(--color-text-main);
  }

  &.--selected &__title {
    color: getColor(--color-primary);
  }

  &__description {
    font-size: 1.1rem;
    font-weight: 400;
    color: getColor(--color-text-sub);
  }
}
</style>
