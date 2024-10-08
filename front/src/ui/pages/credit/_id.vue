<template>
  <div class="courses">
    <PageHeader>
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
    <div class="main">
      <div class="main__info info">
        <div class="info__year">{{ info.year }}</div>
        <div class="info__tag">{{ info.tag }}</div>
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
import { computed, reactive, ref } from "vue";
import { useRoute } from "vue-router";
import { NotFoundError, isResultError } from "~/domain/error";
import { Tag } from "~/domain/tag";
import { creditToDisplay } from "~/presentation/presenters/credit";
import { getDisplayCourseTags } from "~/presentation/presenters/tag";
import CreditCourseListContent from "~/ui/components/CreditCourseListContent.vue";
import IconButton from "~/ui/components/IconButton.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import { createNewTagId } from "~/ui/shared";
import { useCreditYear } from "~/ui/store";
import { timetableUseCase } from "~/usecases";
import type { DisplayCourseTag } from "~/presentation/viewmodels/tag";
import type { CreditCourseListContentState } from "~/ui/components/CreditCourseListContent.vue";

const route = useRoute();

const { id } = route.params as { id: string };
const { creditYear: year } = useCreditYear();

const selectedTag: Tag | undefined = await timetableUseCase
  .getTagById(id)
  .then((result) => {
    if (result instanceof NotFoundError) return undefined;
    if (isResultError(result)) throw result;
    return result;
  });

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
  tag: selectedTag ? `タグ「${selectedTag.name}」` : "すべての授業 ",
  credit: `${totalCredits.value}単位`,
}));

const toggleState = (id: string) => {
  courseIdToState[id] =
    courseIdToState[id] === "default" ? "selected" : "default";
};

const updateView = async (init = false) => {
  const [registeredCourses, tags] = await Promise.all([
    timetableUseCase
      .listRegisteredCourses(
        year.value === 0 ? undefined : year.value,
        selectedTag?.id
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
};

await updateView(true);

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
  updateView();
};

const onClickTag = async (course: VMCourse, tag: DisplayCourseTag) => {
  tag.assign = !tag.assign;
  await timetableUseCase.updateRegisteredCourse(course.id, {
    tagIds: course.tags.filter(({ assign }) => assign).map(({ id }) => id),
  });
  updateView();
};
</script>

<style lang="scss" scoped>
@import "~/ui/styles";

.courses {
  @include max-width;
  height: 100vh;

  display: flex;
  flex-direction: column;
  gap: $spacing-6;

  padding-bottom: $spacing-4;
}

.main {
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

.info {
  display: flex;
  justify-content: left;
  gap: $spacing-3;

  user-select: none;
}
</style>
