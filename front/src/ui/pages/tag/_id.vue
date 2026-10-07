<template>
  <div class="wrapper">
    <PageHeader>
      <template #title>タグの編集</template>
    </PageHeader>
    <Banner class="banner">
      タグの並び替えは
      <RouterLink class="link" to="/credit">「単位数」画面</RouterLink>
      から行えます。
    </Banner>
    <nav class="tags">
      <IconButton
        class="add-icon"
        size="small"
        icon-name="add"
        @click="addTag"
      />
      <TagListSmall
        :tags="allTags"
        :selected-id="id"
        @click="(id_) => $router.push(`/tag/${id_}`)"
      />
    </nav>
    <div class="main">
      <div v-if="selectedTag" class="main__edit edit">
        <LabeledTextField label="タグ名">
          <TextFieldSingleLine v-model="selectedTag.name" @change="updateTag" />
        </LabeledTextField>
        <LabeledTextField label="タグの色">
          <TagColorSelect v-model="selectedTag.color" @change="updateTag" />
        </LabeledTextField>
      </div>
      <div v-if="selectedTag">
        <TertiaryButton color="danger" @click="deleteTag">
          <template #icon>
            <span class="material-symbols-outlined">delete</span>
          </template>
          <template #text>タグを削除する</template>
        </TertiaryButton>
      </div>
      <div class="main__info info">
        <div class="info__tag">授業一覧</div>
        <div class="info__credit">({{ totalCredits }} 単位)</div>
      </div>
      <div class="main__mask">
        <div class="main__courses">
          <ul v-if="courses.length > 0" class="main__courses-list">
            <li v-for="course in courses" :key="course.id" class="main__course">
              <div class="course-code">{{ course.code }}</div>
              <div class="course-info">
                <div class="course-name">{{ course.name }}</div>
                <div class="course-credit">
                  ({{ course.year }}年度 {{ course.credit }}単位)
                </div>
              </div>
            </li>
          </ul>
          <div v-else class="main__no-course">
            {{ noCourseMessage }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { NotFoundError, isResultError } from "~/domain/error";
import { Tag } from "~/domain/tag";
import { creditToDisplay } from "~/presentation/presenters/credit";
import { getDisplayCourseTags } from "~/presentation/presenters/tag";
import Banner from "~/ui/components/Banner.vue";
import IconButton from "~/ui/components/IconButton.vue";
import LabeledTextField from "~/ui/components/LabeledTextField.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import TagColorSelect from "~/ui/components/TagColorSelect.vue";
import TagListSmall from "~/ui/components/TagListSmall.vue";
import TertiaryButton from "~/ui/components/TertiaryButton.vue";
import TextFieldSingleLine from "~/ui/components/TextFieldSingleLine.vue";
import { timetableUseCase } from "~/usecases";
import type { DisplayCourseTag } from "~/presentation/viewmodels/tag";

const router = useRouter();
const route = useRoute();

const id = computed(() => route.params.id as string);
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

const openCourses = reactive(new Set<string>());

const totalCredits = ref<string>("");

type VMCourse = {
  id: string;
  name: string;
  code: string;
  credit: string;
  tags: DisplayCourseTag[];
  year: number;
};

const courses = ref<VMCourse[]>([]);

const noCourseMessage = ref<string>("");

async function updateView(init = false) {
  const [registeredCourses, tags] = await Promise.all([
    timetableUseCase
      .listRegisteredCourses(undefined, selectedTag.value?.id)
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
        year: registeredCourse.year,
      };
    })
    .sort((courseA, courseB) => {
      if (courseA.year !== courseB.year) {
        return courseA.year < courseB.year ? -1 : 1;
      }

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

    openCourses.clear();
  }
}

const addTag = async () => {
  const result = await timetableUseCase.createTag(null);
  if (isResultError(result)) throw result;
  allTags.value.push(result);
  await router.push(`/tag/${result.id}`);
};

const updateTag = async () => {
  if (!selectedTag.value) return;
  const { id, name, color } = selectedTag.value;
  await timetableUseCase.updateTag(id, name, color);
  await updateView();
};

const deleteTag = async () => {
  if (!selectedTag.value) return;

  if (
    courses.value.length > 0 &&
    !confirm(
      `このタグには${courses.value.length}件の授業が登録されています。\n本当に削除しますか？`
    )
  ) {
    return;
  }

  const tagId = selectedTag.value.id;
  await timetableUseCase.deleteTag(tagId);
  allTags.value = allTags.value.filter((tag) => tag.id !== tagId);
  await router.push("/tag/all-courses");
};
</script>

<style lang="scss" scoped>
@use "~/ui/styles/variable" as *;
@use "~/ui/styles/mixin" as *;

.wrapper {
  @include max-width;
  height: 100vh;

  display: flex;
  flex-direction: column;
  gap: $spacing-6;

  padding-bottom: $spacing-4;

  @include pc {
    display: grid;
    grid-template: "header header" "banner banner" auto "tags main" 1fr / auto 1fr;
  }
}

.header {
  grid-area: header;
}

.banner {
  .link {
    text-decoration: underline;
  }

  grid-area: banner;
}

.main {
  grid-area: main;
  flex-grow: 1;

  display: flex;
  flex-direction: column;
  gap: $spacing-3;

  &__mask {
    flex: 1 1 0;

    overflow-y: auto;
    @include scroll-mask;
  }

  &__courses {
    padding: $spacing-3 $spacing-2 $spacing-6 $spacing-0; // padding of scrollable element
  }

  &__no-course {
    color: getColor(--color-text-sub);
    font-size: $font-small;
    line-height: $single-line;
  }

  &__courses-list {
    display: flex;
    flex-direction: column;
    gap: $spacing-2;
  }

  &__course {
    display: flex;
    gap: $spacing-1;
    align-items: baseline;

    .course-code {
      width: 4.3em;
    }
    .course-info {
      display: flex;
      gap: 0 $spacing-1;
      align-items: baseline;
      flex: 1;
      flex-wrap: wrap;
    }
    .course-code,
    .course-credit {
      color: getColor(--color-text-sub);
      font-size: $font-small;
      line-height: $single-line;
    }
    .course-name {
      color: getColor(--color-text-main);
      font-size: $font-medium;
      line-height: $single-line;
    }
  }
}

.tags {
  grid-area: tags;
  display: none;

  padding-right: $spacing-3;
  border-right: 1px solid lightgray;

  .add-icon {
    margin-left: auto;
    margin-bottom: $spacing-3;
  }

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
