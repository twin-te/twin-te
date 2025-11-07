<script setup lang="ts">
import { captureException, startSpan } from "@sentry/vue";
import dayjs from "dayjs";
import { computed, ref } from "vue";
import { useRouter } from "vue-router";
import { isResultError } from "~/domain/error";
import { removeDuplicateSchedules, sortSchedules } from "~/domain/schedule";
import { getKdbClassroom } from "~/infrastructure/local/courseLocationExcel";
import { LocalStorage } from "~/infrastructure/localstorage";
import { registeredCourseToDisplay } from "~/presentation/presenters/course";
import Button from "~/ui/components/Button.vue";
import CardAdd from "~/ui/components/CardAdd.vue";
import Checkbox from "~/ui/components/Checkbox.vue";
import CourseDetailMini from "~/ui/components/CourseDetailMini.vue";
import IconButton from "~/ui/components/IconButton.vue";
import InputButtonFile from "~/ui/components/InputButtonFile.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import { useSetting, useToast } from "~/ui/store";
import { timetableUseCase } from "~/usecases";

const { displayToast } = useToast();

const router = useRouter();

function goBack() {
  if (currentStep.value === "description") {
    router.back();
  } else {
    currentStep.value = "description";
  }
}

const steps = ["description", "upload", "apply"] as const;
const currentStep = ref<typeof steps[number]>("description");

const localStorage = LocalStorage.getInstance();
const latestData = ref(localStorage.get("courseLocationInfo"));

const dataLength = computed(() =>
  latestData.value ? Object.keys(latestData.value.courseLocations).length : 0
);

/* upload */
const loadState = ref<"ready" | "loading" | "error" | "ok">("ready");
async function load(file: File) {
  loadState.value = "loading";
  startSpan(
    {
      name: "Excel Load",
      op: "excel.load",
      attributes: {
        name: file.name,
        size: file.size,
        type: file.type,
      },
    },
    async (span) => {
      try {
        const data = await getKdbClassroom(file);
        latestData.value = data;
        localStorage.set("courseLocationInfo", data);
        loadState.value = "ok";
        span?.setAttribute("excel.load.success", true);
      } catch (error) {
        loadState.value = "error";
        span?.setAttribute("excel.load.success", false);
        captureException(error);
      }
    }
  );
}

/* apply */
const dayjsFormat = "YYYY/MM/DD HH:mm:ss";

const { appliedYear } = useSetting();

const tags = await timetableUseCase.listTags().then((result) => {
  if (isResultError(result)) throw result;
  return result;
});

const registered = await timetableUseCase
  .listRegisteredCourses(appliedYear.value)
  .then((result) => {
    if (isResultError(result)) throw result;
    return result;
  });
const registeredMap = new Map(registered.map((course) => [course.id, course]));

if (isResultError(registered)) throw registered;

type CourseWithLocation = ReturnType<typeof registeredCourseToDisplay> & {
  newLocation: string;
  selected: boolean;
};

const coursesWithChange = ref<CourseWithLocation[]>([]);
const coursesWithoutChange = ref<CourseWithLocation[]>([]);

const selectedCourses = computed(() =>
  coursesWithChange.value.filter((course) => course.selected)
);

const initializeCourseSelection = () => {
  const allCourses = registered
    .map((course) => registeredCourseToDisplay(course, tags))
    .map((course) => {
      const newLocation = latestData.value?.courseLocations[course.code] ?? "";
      return {
        ...course,
        newLocation,
        selected: course.room !== newLocation,
      };
    })
    .filter((course) => course.newLocation);

  coursesWithChange.value = allCourses.filter(
    (course) => course.room !== course.newLocation
  );
  coursesWithoutChange.value = allCourses
    .filter((course) => course.room === course.newLocation)
    .sort((a, b) => (a.room === undefined ? -1 : b.room === undefined ? 1 : 0)); // 未登録のものを先に
};

const toggleCourseSelection = (courseId: string) => {
  const course = coursesWithChange.value.find((c) => c.id === courseId);
  if (!course) return;
  course.selected = !course.selected;
};

const uploadLoading = ref(false);
async function upload() {
  uploadLoading.value = true;

  await Promise.all(
    selectedCourses.value.map(async (course) => {
      const schedules = sortSchedules(
        removeDuplicateSchedules(registeredMap.get(course.id)?.schedules ?? [])
      );
      await timetableUseCase
        .updateRegisteredCourse(course.id, {
          rooms: [
            {
              name: course.newLocation,
              schedules,
            },
          ],
          schedules,
        })
        .then((result) => {
          if (isResultError(result)) throw result;
          return result;
        });
    })
  );

  uploadLoading.value = false;
  displayToast("授業場所の登録が完了しました。", {
    displayPeriod: 5000,
    type: "primary",
  });
  await router.push("/");
}
</script>

<template>
  <div class="import-excel">
    <PageHeader>
      <template #left-button-icon>
        <IconButton
          size="large"
          color="normal"
          icon-name="arrow_back"
          @click="goBack"
        ></IconButton>
      </template>
      <template #title>Excel ファイルから登録</template>
    </PageHeader>
    <main class="main">
      <div v-if="currentStep === 'description'" class="page page-description">
        <p>
          このページでは、授業場所の Excel ファイルを用いて、
          現在の年度に登録されている授業に教室を登録することができます。
        </p>

        <div class="cards">
          <CardAdd
            iconName="upload_file"
            heading="Excel ファイルを新しくアップロードする"
            text=""
            @click-next-button="currentStep = 'upload'"
          />
          <CardAdd
            heading="以前アップロードした Excel ファイルの情報で登録する"
            text="アップロードしたデータはデバイスのみに保存されています。"
            icon-name="file_open"
            :disabled="latestData == null"
            @click-next-button="
              () => {
                currentStep = 'apply';
                initializeCourseSelection();
              }
            "
          />
        </div>

        <section class="info">
          <h5 class="header">注意事項</h5>
          <ul class="list">
            <li>
              登録した授業場所の情報は、このアカウントにのみ保存されます。
            </li>
            <li>
              新たに授業を追加した際には、再度データの登録操作が必要です。
            </li>
            <li>
              授業場所を登録した Twin:te
              のスクリーンショット等は、取り扱いに十分注意していただき、
              大学による注意事項を守って利用してください。
            </li>
          </ul>
        </section>
      </div>
      <div v-else-if="currentStep === 'upload'" class="page page-upload">
        <p class="upload__header">Excel ファイルを選択してください</p>
        <InputButtonFile name="excel-file" accept=".xlsx" @change-file="load">
          アップロードする
        </InputButtonFile>

        <div v-if="loadState === 'loading'" class="loading">読み込み中...</div>
        <div v-else-if="loadState === 'error'" class="load-error">
          アップロードされたファイルを解析することができませんでした。<br />
          正しい Excel ファイルを選択したかを確認していただき、
          解決しない場合は運営までお問い合わせください。
        </div>
        <div v-else-if="loadState === 'ok'" class="load-ok">
          <p>
            ファイルを読み込みました（データ数:
            {{ dataLength.toLocaleString() }}
            件）
          </p>
        </div>

        <div class="download-info">
          <h6>Excel ファイルのダウンロード方法</h6>
          Excel ファイルは筑波大学関係者限定で配布されています。
          以下のリンクから、<code>kdb_(年度)--ja.xlsx</code>
          ファイルをダウンロードしてください。<br />
          <a href="https://bit.ly/UT-classroominfo" target="_blank">
            授業場所 Excel ファイルのダウンロードはこちら（要認証）
          </a>
        </div>

        <Button
          class="next-button"
          color="primary"
          size="medium"
          layout="fill"
          :state="loadState === 'ok' ? 'default' : 'disabled'"
          @click="
            () => {
              currentStep = 'apply';
              initializeCourseSelection();
            }
          "
        >
          次へ
        </Button>
      </div>
      <div v-else-if="currentStep === 'apply'" class="page page-apply">
        <div class="label">このデータを使用して登録します:</div>
        <div class="data-info">
          <div class="data-info__content">
            <div class="title">アップロード日</div>
            <div class="content">
              {{ dayjs(latestData!.uploadAt).format(dayjsFormat) }}
            </div>
          </div>
          <div class="data-info__content">
            <div class="title">授業データ数</div>
            <div class="content">
              {{ dataLength.toLocaleString() }}
              件
            </div>
          </div>
        </div>
        <div v-if="coursesWithChange.length === 0" class="no-data">
          登録できる授業が存在しません
        </div>
        <div v-else class="cards__mask">
          <div v-if="coursesWithChange.length > 0" class="group">
            <div class="label">以下の授業に授業場所が登録されます:</div>
            <div class="cards">
              <div
                v-for="course in coursesWithChange"
                :key="course.id"
                class="course-card"
              >
                <Checkbox
                  :isChecked="course.selected"
                  @clickCheckbox.stop="toggleCourseSelection(course.id)"
                />
                <div>
                  <div class="name">{{ course.name }}</div>
                  <div class="details">
                    <CourseDetailMini
                      icon-name="schedule"
                      :text="course.schedule.full"
                    />
                    <CourseDetailMini
                      icon-name="room"
                      :text="`${course.room || '未登録'} → ${
                        course.newLocation
                      }`"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div v-if="coursesWithoutChange.length > 0" class="group">
            <div class="label">
              以下の授業は既に同じ授業場所が登録されています:
            </div>
            <div class="cards">
              <div
                v-for="course in coursesWithoutChange"
                :key="course.id"
                class="course-card"
              >
                <div>
                  <div class="name">{{ course.name }}</div>
                  <div class="details">
                    <CourseDetailMini
                      icon-name="schedule"
                      :text="course.schedule.full"
                    />
                    <CourseDetailMini
                      icon-name="room"
                      :text="course.newLocation"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <Button
          class="btn-upload"
          color="primary"
          size="medium"
          layout="fill"
          :state="
            uploadLoading || selectedCourses.length === 0
              ? 'disabled'
              : 'default'
          "
          @click="upload"
        >
          このデータを登録する</Button
        >
      </div>
    </main>
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles/variable";
@use "~/ui/styles/mixin";

.import-excel {
  @include mixin.max-width;
  height: 100vh;
  display: flex;
  flex-direction: column;
  gap: variable.$spacing-8;
}

.main {
  flex-grow: 1;
}

.page {
  height: 100%;
}

.page-description {
  .cards {
    margin-top: variable.$spacing-8;
    display: flex;
    flex-direction: column;
    gap: variable.$spacing-4;
  }
}

.page-upload {
  display: flex;
  flex-direction: column;

  &__header {
    font-weight: 500;
  }

  .loading,
  .load-error,
  .load-ok {
    margin-top: variable.$spacing-4;
  }

  .loading {
    color: darkgray;
  }

  .load-error {
    border: solid red 2px;
    color: red;
    border-radius: 4px;
    padding: 12px;
  }

  .load-ok {
    display: flex;
    flex-direction: column;
  }

  .download-info {
    margin-top: variable.$spacing-8;
    line-height: 2.5rem;

    h6 {
      font-weight: bold;
    }

    code {
      font-family: monospace;
    }

    a {
      text-decoration: underline;
    }
  }

  .next-button {
    margin: variable.$spacing-12 auto 0;
  }
}

.page-apply {
  display: flex;
  flex-direction: column;

  .no-data {
    flex: 1 1 0;
    margin: variable.$spacing-4 0 variable.$spacing-6;
    font-size: 1.4rem;
    line-height: 1.4;

    .small {
      margin-top: variable.$spacing-1;
      font-size: 1.2rem;
      color: variable.getColor(--color-text-sub);
    }
  }

  .data-info {
    width: max-content;
    margin-top: variable.$spacing-2;

    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    gap: variable.$spacing-12;

    &__content {
      .title {
        font-weight: 500;
        font-size: 1.2rem;
      }
      .content {
        margin-top: variable.$spacing-1;
        font-size: 1.5rem;
      }
    }
  }

  .cards {
    display: flex;
    flex-direction: column;
    margin: variable.$spacing-2 0 variable.$spacing-4;
    padding: variable.$spacing-3 variable.$spacing-2 variable.$spacing-3
      variable.$spacing-0;
    gap: variable.$spacing-4;

    &__mask {
      margin-top: variable.$spacing-2;
      flex: 1 1 0;
      overflow-y: auto;
      @include mixin.scroll-mask;
    }
  }

  .group {
    margin-top: variable.$spacing-8;
  }

  .course-card {
    display: flex;
    margin-inline-start: variable.$spacing-2;
    gap: variable.$spacing-4;
    align-items: center;

    .name {
      @include mixin.text-main;
    }
    .details {
      display: flex;
      flex-wrap: wrap;
      margin-top: variable.$spacing-1;
      gap: variable.$spacing-1 variable.$spacing-4;
    }
  }

  .btn-upload {
    margin: variable.$spacing-2 auto variable.$spacing-4;
  }
}

.info {
  margin-top: variable.$spacing-8;

  .header {
    font-size: 1.4rem;
    font-weight: 500;
  }

  .list {
    margin-top: variable.$spacing-2;
    li {
      list-style: disc inside;
      line-height: 130%;
      margin-bottom: variable.$spacing-2;
    }
  }
}
</style>
