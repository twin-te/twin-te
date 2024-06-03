<script setup lang="ts">

import { useHead } from "@vueuse/head";
import { computed, reactive, ref, watch } from "vue";
import { isResultError } from "~/domain/error";
import { academicYears } from "~/domain/year";
import { creditToDisplay } from "~/presentation/presenters/credit";
import { getDisplayCreditTags } from "~/presentation/presenters/tag";
import { DisplayCreditTag } from "~/presentation/viewmodels/tag";
import Button from "~/ui/components/Button.vue";
import Dropdown from "~/ui/components/Dropdown.vue";
import IconButton from "~/ui/components/IconButton.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import TagListContent from "~/ui/components/TagListContent.vue";
import { useCreditYear } from "~/ui/store";
import { timetableUseCase } from "~/usecases";

useHead({
  title: 'Twin:te | タグ'
})

const {
  creditYear: year,
  setCreditYear,
  setCreditYearToAll,
} = useCreditYear()

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

/** display credit tag */
const displayCreditTags = ref<DisplayCreditTag[]>([]);
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

const createTag = async () => {
  try {
    await timetableUseCase.createTag('新しいタグ').then((result) => {
      if (isResultError(result)) throw result;
      return result;
    })
  } finally {
    await updateDisplayCreditTags();
  }
};

watch(year, () => updateDisplayCreditTags(), {
  immediate: true,
});
</script>

<template>
  <div class="tags">
    <PageHeader>
      <template #left-button-icon>
        <IconButton
          size="large"
          color="normal"
          icon-name="arrow_back"
          @click="$router.push('/')"
        ></IconButton>
      </template>
      <template #title>タグ</template>
    </PageHeader>
    <section class="tag-list">
      <Dropdown
        :selectedOption="selectedYearOption"
        :options="yearOptions"
        label="対象年度"
        @update:selected-option="updateCreditYearOption"
      />
      <div class="tag-list__header">
        <Button size="medium" @click="createTag">新しいタグを作成</Button>
      </div>
      <div class="tag-list__contents">
        <TagListContent
          name="すべての授業"
          :credit="totalCredits"
        >
          <template #btns>
            <IconButton
              size="small"
              color="normal"
              icon-name="chevron_right"
              @click="$router.push('/credit/all-courses')"
            />
          </template>
        </TagListContent>
        <TagListContent
          v-for="tag in displayCreditTags"
          :key="tag.id"
          :name="tag.name"
          :credit="tag.credit"
        >
          <template #btns>
            <IconButton
              size="small"
              color="normal"
              icon-name="chevron_right"
              @click="$router.push(`/tags/${tag.id}`)"
            ></IconButton>
          </template>
        </TagListContent>
      </div>
    </section>
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles/mixin" as *;
@use "~/ui/styles/variable" as *;

.container {
  //display: grid;
  height: 100dvh;
  
  grid-template-columns: 1fr;
  @include large-screen {
    grid-template-columns: 400px 1fr;
  }
}

.tag-list {
  display: flex;
  gap: $spacing-2;
  flex-direction: column;
  
  &__header {
    flex: none;
    display: flex;
    justify-content: flex-end;
    align-items: center;
  }
  &__contents {
    
  }
}

.tags-list {
  background-color: #c0c0f0;
  min-height: 100%;
}
</style>
