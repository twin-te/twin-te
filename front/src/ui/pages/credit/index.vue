<template>
  <div class="credit">
    <PageHeader>
      <template #left-button-icon>
        <IconButton
          size="large"
          color="normal"
          icon-name="arrow_back"
          @click="$router.push('/')"
        ></IconButton>
      </template>
      <template #title>単位数</template>
    </PageHeader>
    <Dropdown
      :selectedOption="selectedYearOption"
      :options="yearOptions"
      label="対象年度"
      :state="mode === 'default' ? 'default' : 'disabled'"
      @update:selected-option="updateCreditYearOption"
    ></Dropdown>
    <section class="tags">
      <div class="tags__header">
        <div class="tags__label">分類</div>
        <Button
          size="small"
          :state="editingTagId || dragging ? 'disabled' : 'default'"
          @click="toggleMode"
          >{{
            mode === "default" ? "タグの作成・編集" : "タグの作成・編集を終わる"
          }}
        </Button>
      </div>
      <div class="tags__mask">
        <div ref="tagsContentsRef" class="tags__contents">
          <TagListContent
            v-show="mode === 'default'"
            name="すべての授業"
            :credit="totalCredits"
          >
            <template #btns>
              <IconButton
                v-show="mode === 'default'"
                size="small"
                color="normal"
                icon-name="chevron_right"
                @click="$router.push('/credit/all-courses')"
              ></IconButton>
            </template>
          </TagListContent>
          <draggable
            :model-value="displayCreditTags"
            item-key="id"
            :animation="250"
            handle=".tag-list-content__drag-icon"
            :disabled="editingTagId != undefined"
            @update:model-value="onChangeOrder"
            @start="dragging = true"
            @end="dragging = false"
          >
            <template #item="{ element }">
              <TagListContent
                :name="`タグ「${element.name}」`"
                :credit="element.credit"
                :mode="mode"
                :textfield="editingTagId === element.id"
                :drag-handle="editingTagId ? 'disabled' : 'show'"
              >
                <template #textfiled>
                  <TextFieldSingleLine
                    :id="`text-field-single-line--${element.id}`"
                    v-model.trim="element.name"
                    placeholder="タグ名"
                    @enter-text-field="() => onClickNormalBtn(element)"
                  ></TextFieldSingleLine>
                </template>
                <template #btns>
                  <IconButton
                    v-show="mode === 'default'"
                    size="small"
                    color="normal"
                    icon-name="chevron_right"
                    @click="$router.push(`/credit/${element.id}`)"
                  ></IconButton>
                  <IconButton
                    v-show="mode === 'edit'"
                    size="small"
                    color="normal"
                    :icon-name="editingTagId === element.id ? 'check' : 'edit'"
                    :state="
                      !editingTagId ||
                      (editingTagId === element.id && element.name !== '')
                        ? 'default'
                        : 'disabled'
                    "
                    @click="() => onClickNormalBtn(element)"
                  ></IconButton>
                  <IconButton
                    v-show="mode === 'edit'"
                    size="small"
                    color="danger"
                    :icon-name="isNewTagId(element.id) ? ' clear' : 'delete'"
                    :state="
                      editingTagId && !isNewTagId(element.id)
                        ? 'disabled'
                        : 'default'
                    "
                    @click="() => onClickDangerBtn(element)"
                  ></IconButton>
                </template>
              </TagListContent>
            </template>
          </draggable>
          <div v-if="displayCreditTags.length === 0" class="tags__no-tag">
            作成済みのタグがありません。<br />
            タグを作成すると授業を分類することができます。
          </div>
        </div>
      </div>
      <div
        v-show="mode === 'edit'"
        :class="{
          'tags__add-btn': true,
          'add-btn': true,
          '--disabled': editingTagId || dragging,
        }"
        @click="onClickAddBtn"
      >
        <div class="add-btn__icon material-icons">add</div>
        <div class="add-btn__value">タグを新たに作成する</div>
      </div>
    </section>
    <Modal
      v-if="tagToBeDeleted"
      class="delete-tag-modal"
      size="small"
      @click="tagToBeDeleted = undefined"
    >
      <template #title>タグを削除しますか？</template>
      <template #contents>
        タグ「{{ tagToBeDeleted.name }}」を削除しますか？<br />
        現在このタグを{{
          numCoursesAssociatedWithTagToBeDeleted
        }}件の授業に割り当てています。<br />
        タグを削除すると、割り当てた全ての授業との紐付けが解除されます。
      </template>
      <template #button>
        <Button
          size="medium"
          layout="fill"
          color="base"
          @click="tagToBeDeleted = undefined"
        >
          キャンセル
        </Button>
        <Button
          size="medium"
          layout="fill"
          color="danger"
          @click="onClickDeleteModal"
        >
          削除
        </Button>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { useHead } from "@vueuse/head";
import { computed, reactive, ref, watch, watchEffect } from "vue";
import draggable from "vuedraggable";
import { isResultError } from "~/domain/error";
import { academicYears } from "~/domain/year";
import { creditToDisplay } from "~/presentation/presenters/credit";
import { getDisplayCreditTags } from "~/presentation/presenters/tag";
import Button from "~/ui/components/Button.vue";
import Dropdown from "~/ui/components/Dropdown.vue";
import IconButton from "~/ui/components/IconButton.vue";
import Modal from "~/ui/components/Modal.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import TagListContent from "~/ui/components/TagListContent.vue";
import TextFieldSingleLine from "~/ui/components/TextFieldSingleLine.vue";
import { useFocus } from "~/ui/hooks/useFocus";
import { useStringToggle } from "~/ui/hooks/useStringToggle";
import { isNewTagId, createNewTagId } from "~/ui/shared";
import { useCreditYear } from "~/ui/store";
import { timetableUseCase } from "~/usecases";
import type { DisplayCreditTag } from "~/presentation/viewmodels/tag";

console.log("/credit");

useHead({
  title: "Twin:te | 単位数",
});

/** year */
const {
  creditYear: year,
  setCreditYear,
  setCreditYearToAll,
} = useCreditYear();

/** mode */
const [mode, toggleMode] = useStringToggle("default", "edit");

/** year dropdown */
const allYearOption = "すべての年度";
const selectedYearOption = computed(() =>
  year.value === 0 ? allYearOption : `${year.value}年度`
);
const yearOptions: string[] = [
  allYearOption,
  ...academicYears.map((year) => `${year}年度`).reverse(),
];
const updateCreditYearOption = (option: string) => {
  if (option === allYearOption) setCreditYearToAll();
  else setCreditYear(Number(option.slice(0, 4)));
};

/** focus */
const { targetRef: tagsContentsRef, focus } = useFocus();

/** display credit tag */
const displayCreditTags = ref<DisplayCreditTag[]>([]);
const editingTagId = ref<string>();
const dragging = ref(false);
const totalCredits = ref<string>("0.0");

const updateDisplayCreditTags = async () => {
  const [registeredCourses, tags] = await Promise.all([
    timetableUseCase
      .getRegisteredCourses(year.value == 0 ? undefined : year.value, undefined)
      .then((result) => {
        if (isResultError(result)) throw result;
        return result;
      }),
    timetableUseCase.getTags().then((result) => {
      if (isResultError(result)) throw result;
      return result;
    }),
  ]);

  displayCreditTags.value = reactive(
    getDisplayCreditTags(
      registeredCourses,
      tags.sort((tagA, tagB) => tagA.order - tagB.order),
      year.value === 0 ? academicYears : [year.value]
    )
  );

  totalCredits.value = creditToDisplay(
    registeredCourses.reduce(
      (totalCredits, registeredCourse) =>
        totalCredits + registeredCourse.credit,
      0
    )
  );
};

watch(year, () => updateDisplayCreditTags(), {
  immediate: true,
});

const onClickNormalBtn = async (tag: DisplayCreditTag) => {
  if (editingTagId.value === tag.id) {
    editingTagId.value = undefined;
    try {
      if (isNewTagId(tag.id)) {
        await timetableUseCase.createTag(tag.name).then((result) => {
          if (isResultError(result)) throw result;
          return result;
        });
      } else {
        await timetableUseCase
          .updateTagName(tag.id, tag.name)
          .then((result) => {
            if (isResultError(result)) throw result;
            return result;
          });
      }
    } finally {
      updateDisplayCreditTags();
    }
  } else {
    editingTagId.value = tag.id;
    focus([`#text-field-single-line--${tag.id}`, "input"]);
  }
};

const onClickDangerBtn = async (tag: DisplayCreditTag) => {
  if (isNewTagId(tag.id)) {
    updateDisplayCreditTags();
    editingTagId.value = undefined;
  } else tagToBeDeleted.value = tag;
};

const onChangeOrder = async (newTags: DisplayCreditTag[]) => {
  displayCreditTags.value = reactive(newTags);
  await timetableUseCase
    .updateTagOrders(displayCreditTags.value.map(({ id }) => id))
    .then((result) => {
      if (isResultError(result)) throw result;
      return result;
    });
  updateDisplayCreditTags();
};

const onClickAddBtn = () => {
  const id = createNewTagId();
  displayCreditTags.value.push({ id, name: "", credit: "0.0" });
  editingTagId.value = id;
  focus([`#text-field-single-line--${id}`, "input"]);
};

/** delete tag modal */
const tagToBeDeleted = ref<DisplayCreditTag>();
const numCoursesAssociatedWithTagToBeDeleted = ref<number>(0);

watchEffect(() => {
  if (tagToBeDeleted.value == undefined) {
    numCoursesAssociatedWithTagToBeDeleted.value = 0;
    return;
  }

  timetableUseCase
    .getRegisteredCourses(undefined, tagToBeDeleted.value.id)
    .then((result) => {
      if (isResultError(result)) throw result;
      numCoursesAssociatedWithTagToBeDeleted.value = result.length;
    });
});

const onClickDeleteModal = async () => {
  if (tagToBeDeleted.value == undefined) return;
  await timetableUseCase.deleteTag(tagToBeDeleted.value.id).then((result) => {
    if (isResultError(result)) throw result;
    return result;
  });
  tagToBeDeleted.value = undefined;
  updateDisplayCreditTags();
};
</script>

<style lang="scss" scoped>
@import "~/ui/styles";

.credit {
  @include max-width;

  display: flex;
  flex-direction: column;
  gap: $spacing-6;

  padding-bottom: $spacing-6; // spacing at the bottom of the page
}

.tags {
  height: calc(
    100vh - env(safe-area-inset-top) - env(safe-area-inset-bottom) - 19.6rem
  ); // tags 以外の height を引いた分

  display: flex;
  flex-direction: column;
  gap: $spacing-2;

  &__header {
    flex: none;

    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  &__label {
    font-weight: 500;
  }

  &__mask {
    flex: initial;

    overflow-y: auto;
    @include scroll-mask;
  }

  &__contents {
    padding: $spacing-3 $spacing-2 $spacing-3 $spacing-0; // padding of scrollable element
  }

  &__no-tag {
    padding: $spacing-2 $spacing-0 $spacing-0 $spacing-2;

    font-size: $font-small;
    color: getColor(--color-disabled);
    line-height: $multi-line;
  }

  &__add-btn {
    flex: none;
  }
}

.add-btn {
  width: max-content;

  display: flex;
  gap: $spacing-1;

  padding: $spacing-2 $spacing-1;

  @include button-cursor;

  font-size: $font-small;
  color: getColor(--color-button-gray);

  &.--disabled {
    opacity: 0.3;
  }

  &__value {
    font-weight: 500;
  }
}

.delete-tag-modal {
  .button {
    width: calc(50% - $spacing-3);
    &:first-child {
      margin-right: $spacing-3;
    }
    &:last-child {
      margin-left: $spacing-3;
    }
  }
}
</style>
