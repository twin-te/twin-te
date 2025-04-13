<template>
  <div class="courses">
    <PageHeader class="header">
      <template #left-button-icon>
        <IconButton
          size="large"
          color="normal"
          icon-name="arrow_back"
          @click="$router.push('/credit')"
        ></IconButton>
      </template>
      <template #title>単位数</template>
    </PageHeader>
    <nav class="tags">
      <TagListSmall
        :tags="allTags"
        :selected-id="id"
        @click="(id_) => $router.push(`/credit/${id_}`)"
      />
    </nav>
    <div class="main">
      <div v-if="selectedTag" class="main__edit edit">
        <LabeledTextField label="タグ名">
          <TextFieldSingleLine v-model="selectedTag.name" @change="updateTag" />
        </LabeledTextField>
        <TagColorSelect v-model="selectedTag.color" @change="updateTag" />
      </div>
      <div class="main__info info">
        <div class="info__tag">{{ info.tag }}</div>
        <div class="info__year">({{ info.year }})</div>
        <div class="info__credit">{{ info.credit }}</div>
      </div>
      <div class="main__mask">
        <div class="main__courses">
          <CreditCourseListContent
            v-for="course in courses"
            :key="course.id"
            :state="courseIdToState[course.id]"
            :code="course.code"
            :name="course.name"
            :credit="course.credit"
            :tags="course.tags"
            @click="toggleState(course.id)"
            @create-tag="(tagName) => onCreateTag(course, tagName)"
            @click-tag="(tag) => onClickTag(course, tag)"
          ></CreditCourseListContent>
          <div v-if="courses.length === 0" class="main__no-course">
            {{ noCourseMessage }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref, watch } from "vue";
import { useRoute } from "vue-router";
import { NotFoundError, isResultError } from "~/domain/error";
import { Tag } from "~/domain/tag";
import { creditToDisplay } from "~/presentation/presenters/credit";
import { getDisplayCourseTags } from "~/presentation/presenters/tag";
import CreditCourseListContent from "~/ui/components/CreditCourseListContent.vue";
import IconButton from "~/ui/components/IconButton.vue";
import LabeledTextField from "~/ui/components/LabeledTextField.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import TagColorSelect from "~/ui/components/TagColorSelect.vue";
import TagListSmall from "~/ui/components/TagListSmall.vue";
import TextFieldSingleLine from "~/ui/components/TextFieldSingleLine.vue";
import { createNewTagId } from "~/ui/shared";
import { useCreditYear } from "~/ui/store";
import { timetableUseCase } from "~/usecases";
import type { DisplayCourseTag } from "~/presentation/viewmodels/tag";
import type { CreditCourseListContentState } from "~/ui/components/CreditCourseListContent.vue";

const route = useRoute();

const id = computed(() => route.params.id as string);
const { creditYear: year } = useCreditYear();
const allTags = ref<Tag[]>([]);

const selectedTag = ref<Tag>();

watch(
  id,
  async (newId) => {
    selectedTag.value =
      newId === "all-courses"
        ? undefined
        : await timetableUseCase.getTagById(newId).then((result) => {
            if (result instanceof NotFoundError) return undefined;
            if (isResultError(result)) throw result;
            return result;
          });
    await updateView(true);
  },
  { immediate: true }
);

const courseIdToState = reactive<Record<string, CreditCourseListContentState>>(
  {}
);

const totalCredits = ref<string>("");

type VMCourse = {
  id: string;
  name: string;
  code: string;
  credit: string;
  tags: DisplayCourseTag[];
};

const courses = ref<VMCourse[]>([]);

const noCourseMessage = ref<string>("");

const info = computed(() => ({
  year: year.value === 0 ? "すべての年度" : `${year.value}年度`,
  tag: selectedTag.value
    ? `タグ「${selectedTag.value.name}」`
    : "すべての授業 ",
  credit: `${totalCredits.value}単位`,
}));

const toggleState = (id: string) => {
  courseIdToState[id] =
    courseIdToState[id] === "default" ? "selected" : "default";
};

async function updateView(init = false) {
  const [registeredCourses, tags] = await Promise.all([
    timetableUseCase
      .listRegisteredCourses(
        year.value === 0 ? undefined : year.value,
        selectedTag.value?.id
      )
      .then((result) => {
        if (isResultError(result)) throw result;
        return result;
      }),
    timetableUseCase
      .listTags()
      .then((result) => {
        if (isResultError(result)) throw result;
        return result;
      })
      .then((tags) => {
        return tags.sort((tagA, tagB) => tagA.order - tagB.order);
      })
      .then((tags) => {
        allTags.value = tags;
        return tags;
      }),
  ]);

  courses.value = registeredCourses
    .map((registeredCourse) => {
      return {
        id: registeredCourse.id,
        name: registeredCourse.name,
        code: registeredCourse.code ?? "-",
        credit: creditToDisplay(registeredCourse.credit),
        tags: getDisplayCourseTags(registeredCourse, tags),
      };
    })
    .sort((courseA, courseB) => {
      if (courseA.code !== courseB.code) {
        return courseA.code < courseB.code ? -1 : 1;
      }

      return courseA.name < courseB.name ? -1 : 1;
    });

  totalCredits.value = creditToDisplay(
    registeredCourses.reduce((totalCredits, registeredCourse) => {
      return totalCredits + registeredCourse.credit;
    }, 0)
  );

  if (init) {
    noCourseMessage.value = await timetableUseCase
      .listRegisteredCourses()
      .then((result) => {
        if (isResultError(result)) throw result;
        return result;
      })
      .then((allRegisteredCourses) => {
        if (allRegisteredCourses.length === 0) {
          return "登録済みの授業がありません。";
        }
        return "該当する授業がありません。";
      });

    registeredCourses.forEach(({ id }) => {
      courseIdToState[id] = "default";
    });
  }
}

const updateTag = async () => {
  if (!selectedTag.value) return;
  const { id, name, color } = selectedTag.value;
  await timetableUseCase.updateTag(id, name, color);
  await updateView();
};

const onCreateTag = async (course: VMCourse, tagName: string) => {
  const tagIds = course.tags.filter(({ assign }) => assign).map(({ id }) => id);
  course.tags.push({ id: createNewTagId(), name: tagName, assign: true });
  const newTag = await timetableUseCase.createTag(tagName).then((result) => {
    if (isResultError(result)) throw result;
    return result;
  });
  await timetableUseCase.updateRegisteredCourse(course.id, {
    tagIds: [...tagIds, newTag.id],
  });
  await updateView();
};

const onClickTag = async (course: VMCourse, tag: DisplayCourseTag) => {
  tag.assign = !tag.assign;
  await timetableUseCase.updateRegisteredCourse(course.id, {
    tagIds: course.tags.filter(({ assign }) => assign).map(({ id }) => id),
  });
  await updateView();
};
</script>

<style lang="scss" scoped>
@use "~/ui/styles" as *;

.courses {
  @include max-width;
  height: 100vh;

  display: flex;
  flex-direction: column;
  gap: $spacing-6;

  padding-bottom: $spacing-4;

  @include pc {
    display: grid;
    grid-template: "header header" auto "tags main" 1fr / auto 1fr;
  }
}

.header {
  grid-area: header;
}

.main {
  grid-area: main;
  flex-grow: 1;

  display: flex;
  flex-direction: column;
  gap: $spacing-5;

  &__mask {
    flex: 1 1 0px;

    overflow-y: auto;
    @include scroll-mask;
  }

  &__courses {
    padding: $spacing-3 $spacing-2 $spacing-3 $spacing-0; // padding of scrollable element
  }

  &__no-course {
    color: getColor(--color-text-sub);
    font-size: $font-small;
    line-height: $single-line;
  }
}

.tags {
  grid-area: tags;
  display: none;

  @include pc {
    display: block;
  }
}

.info {
  display: flex;
  justify-content: left;
  align-items: baseline;
  gap: $spacing-1;

  user-select: none;

  &__year {
    font-size: 0.9em;
  }

  &__credit {
    margin-inline-start: $spacing-1;
  }
}

.edit {
  display: flex;
  flex-direction: column;
  gap: $spacing-5;
}
</style>
