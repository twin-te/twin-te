<script lang="ts">
//Notifcation Settings
declare global {
  // eslint-disable-next-line no-unused-vars
  interface Window {
    android?: {
      openSettings: () => void;
      //shareが無いとTypeError...どこまで弄っていいか分からなかったのでこのまま
      share: (message: string) => void;
    };
    webkit?: {
      messageHandlers?: {
        iPhoneSettings?: {
          postMessage: (hoge: string) => void;
        };
        share?: {
          postMessage: (message: string) => void;
        };
      };
    };
  }
}
</script>
<template>
  <div class="settings">
    <PageHeader>
      <template #left-button-icon>
        <IconButton
          size="large"
          color="normal"
          icon-name="arrow_back"
          @click="$router.back()"
        ></IconButton>
      </template>
      <template #title>設定</template>
    </PageHeader>
    <div class="main">
      <div class="main__contents">
        <div class="main__content">
          ダークテーマ
          <ToggleSwitch
            class="switch"
            :isChecked="setting.darkMode"
            @click-toggle-switch="
              updateSetting({ darkMode: !setting.darkMode })
            "
          />
        </div>
        <div class="main__content">
          土曜授業を表示する
          <ToggleSwitch
            class="switch"
            :isChecked="setting.saturdayCourseMode"
            @click-toggle-switch="
              updateSetting({ saturdayCourseMode: !setting.saturdayCourseMode })
            "
          />
        </div>
        <div class="main__content">
          8限まで表示する(大学院生用)
          <ToggleSwitch
            class="switch"
            :isChecked="setting.nightPeriodMode"
            @click-toggle-switch="
              updateSetting({ nightPeriodMode: !setting.nightPeriodMode })
            "
          />
        </div>
        <div class="main__content">
          各時限の開始・終了時刻を表示する
          <ToggleSwitch
            class="switch"
            :isChecked="setting.timeLabelMode"
            @click-toggle-switch="
              updateSetting({ timeLabelMode: !setting.timeLabelMode })
            "
          />
        </div>
        <div class="main__content--dropdown">
          <p>時間割の表示、授業の検索に適用する年度</p>
          <Dropdown
            :selectedOption="selectedYearOption"
            :options="yearOptions"
            @update:selectedOption="updateSelectedYearOption"
          ></Dropdown>
        </div>
        <div v-show="isMobile()" class="main__content">
          <p>通知</p>
          <Button
            class="button"
            size="small"
            color="base"
            :pauseActiveStyle="false"
            @click="openNotificationSetting()"
            >通知設定を開く</Button
          >
        </div>
        <div v-if="isAuthenticated" class="main__content--ical">
          <p>カレンダー連携（ベータ版）</p>
          <ToggleSwitch
            class="switch"
            :isChecked="icalUrl !== null"
            @click-toggle-switch="onIcalToggle"
          />
          <div v-if="icalUrl" class="ical-detail">
            <p class="ical-description">
              以下のURLをGoogleカレンダーやAppleのカレンダーアプリなどに登録すると、Twin:teの時間割が自動的に同期されます。
            </p>
            <div class="ical-url-row">
              <input
                v-model="icalUrl"
                type="text"
                readonly
                class="ical-url-input"
              />
              <Button
                class="button"
                size="small"
                color="base"
                :pauseActiveStyle="false"
                @click="copyIcalUrl"
                >コピー</Button
              >
            </div>
            <div class="ical-tag-entry">
              <TertiaryButton
                class="ical-tag-entry__link"
                color="primary"
                @click="onClickIcalTagCustomize"
              >
                <template #icon>sell</template>
                <template #text>対象のタグをカスタマイズ</template>
              </TertiaryButton>
              <span
                v-if="issuedIcalTags.length > 0"
                class="ical-tag-entry__label"
                >対象:
                {{ issuedIcalTags.map(({ name }) => name).join("・") }}</span
              >
            </div>
            <div>
              <h5>注意事項</h5>
              <ul class="ical-cautions">
                <li>
                  このURLを知っている人は誰でもあなたの時間割を閲覧できます。取り扱いにご注意ください。
                </li>
                <li>カレンダーへの反映には時間がかかる場合があります。</li>
                <li>一度機能を無効にするとURLが変更されます。</li>
              </ul>
            </div>
          </div>
        </div>
        <div v-if="isAuthenticated" class="main__content--account">
          <p>アカウント情報</p>
          <div class="account-btns">
            <Button
              class="button"
              size="small"
              color="primary"
              :pauseActiveStyle="false"
              @click="logout"
              >ログアウトする</Button
            >
            <Button
              class="button"
              size="small"
              color="danger"
              :pauseActiveStyle="false"
              @click="onClickAccountDeleteModel()"
              >アカウントを削除する</Button
            >
          </div>
        </div>
      </div>
    </div>
    <Modal
      v-if="icalUrl && isIcalTagModalVisible"
      class="ical-tag-modal"
      @click="closeIcalTagModal"
    >
      <template #title>{{
        icalTagModalStep === "select"
          ? "対象のタグをカスタマイズ"
          : "URLを発行しました"
      }}</template>
      <template #contents>
        <div
          v-if="icalTagModalStep === 'select'"
          class="ical-tag-modal__contents"
        >
          <p class="ical-tag-modal__description">
            選択したタグが付いた授業だけに限定したURLを発行します。
          </p>
          <template v-if="displayIcalTags.length > 0">
            <div class="ical-tag-modal__toolbar">
              <span class="ical-tag-modal__selected-label">{{
                selectedIcalTagIds.length > 0
                  ? `${selectedIcalTagIds.length}個のタグを選択中`
                  : "タグを選択してください"
              }}</span>
              <TertiaryButton color="ghost" @click="toggleAllIcalTags">
                <template #icon>{{
                  isAllIcalTagsSelected ? "remove_done" : "done_all"
                }}</template>
                <template #text>{{
                  isAllIcalTagsSelected ? "すべて解除" : "すべて選択"
                }}</template>
              </TertiaryButton>
            </div>
            <div class="ical-tag-modal__tags">
              <TagListContent
                v-for="tag in displayIcalTags"
                :key="tag.id"
                class="ical-tag-modal__tag"
                :name="tag.name"
                :credit="`${tag.courseCount}件`"
                @click="toggleIcalTag(tag.id)"
              >
                <template #btns>
                  <Checkbox :isChecked="selectedIcalTagIds.includes(tag.id)" />
                </template>
              </TagListContent>
            </div>
          </template>
          <p v-else class="ical-tag-modal__empty">
            タグがまだありません。授業の詳細画面からタグを作成して、授業に設定してください。
          </p>
        </div>
        <div v-else class="ical-tag-modal__contents">
          <p class="ical-tag-modal__description">
            以下のタグが付いた授業だけがカレンダーに同期されます。
          </p>
          <div class="ical-tag-modal__issued-tags">
            <Tag v-for="tag in issuedIcalTags" :key="tag.id" :assign="true">
              {{ tag.name }}
            </Tag>
          </div>
          <div class="ical-tag-modal__url">
            <input
              :value="issuedIcalUrl"
              type="text"
              readonly
              class="ical-tag-modal__url-input"
            />
            <Button
              size="small"
              :color="isIssuedIcalUrlCopied ? 'primary' : 'base'"
              :pauseActiveStyle="false"
              @click="copyIssuedIcalUrl"
              >{{ isIssuedIcalUrlCopied ? "コピー済" : "コピー" }}</Button
            >
          </div>
          <ul class="ical-tag-modal__notes">
            <li>※ 元のURL（すべての授業）はそのまま使えます。</li>
            <li>※ 対象のタグを変更するには、新しいURLを発行してください。</li>
          </ul>
        </div>
      </template>
      <template #button>
        <template v-if="icalTagModalStep === 'select'">
          <Button
            class="ical-tag-modal__button"
            size="medium"
            layout="fill"
            color="base"
            @click="closeIcalTagModal"
            >キャンセル</Button
          >
          <Button
            class="ical-tag-modal__button"
            size="medium"
            layout="fill"
            color="primary"
            :state="selectedIcalTagIds.length > 0 ? 'default' : 'disabled'"
            @click="issueIcalUrl"
            >発行する</Button
          >
        </template>
        <template v-else>
          <Button
            class="ical-tag-modal__button"
            size="medium"
            layout="fill"
            color="base"
            @click="backToIcalTagSelect"
            >選び直す</Button
          >
          <Button
            class="ical-tag-modal__button"
            size="medium"
            layout="fill"
            color="primary"
            @click="closeIcalTagModal"
            >閉じる</Button
          >
        </template>
      </template>
    </Modal>
    <Modal
      v-if="isAccountDeletionModalVisible"
      class="account-delete-modal"
      @click="closeAccountDeletionModal"
    >
      <template #title>アカウントを消去しますか？</template>
      <template #contents>
        <p class="modal__text">
          Twin:teに登録した、すべてのデータも消去されます。これには時間割やメモ等を含み、消去後は復元することができません。
        </p>
      </template>
      <template #button>
        <Button
          size="medium"
          layout="fill"
          color="base"
          @click="closeAccountDeletionModal"
          >キャンセル</Button
        >
        <Button
          size="medium"
          layout="fill"
          color="danger"
          @click="confirmDeleteAccount"
          >消去</Button
        >
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { useHead } from "@vueuse/head";
import { computed, onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import {
  InternalServerError,
  isResultError,
  NetworkError,
  UnauthenticatedError,
} from "~/domain/error";
import { academicYears, currentAcademicYear } from "~/domain/year";
import { getIcalUrlWithTags } from "~/presentation/presenters/calendar";
import { getDisplayIcalTags } from "~/presentation/presenters/tag";
import Checkbox from "~/ui/components/Checkbox.vue";
import Dropdown from "~/ui/components/Dropdown.vue";
import IconButton from "~/ui/components/IconButton.vue";
import Modal from "~/ui/components/Modal.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import Tag from "~/ui/components/Tag.vue";
import TagListContent from "~/ui/components/TagListContent.vue";
import TertiaryButton from "~/ui/components/TertiaryButton.vue";
import ToggleSwitch from "~/ui/components/ToggleSwitch.vue";
import { useSwitch } from "~/ui/hooks/useSwitch";
import { isiOS, isMobile } from "~/ui/ua";
import { authUseCase, calendarUseCase, timetableUseCase } from "~/usecases";
import Button from "../components/Button.vue";
import { useAuth, useSetting, useToast } from "../store";
import { getLogoutUrl, redirectToUrl } from "../url";
import type { DisplayIcalTag } from "~/presentation/viewmodels/tag";

const router = useRouter();
const { displayToast } = useToast();

useHead({
  title: "Twin:te | 設定",
});

const { setting, updateSetting } = useSetting();

const { isAuthenticated } = useAuth();

/** ical subscription */
const icalUrl = ref<string | null>(null);

onMounted(async () => {
  if (!isAuthenticated.value) {
    return;
  }
  const result = await calendarUseCase.getIcalSubscriptionUrl();
  if (!isResultError(result) && "url" in result && result.url) {
    icalUrl.value = result.url;
  } else {
    icalUrl.value = null;
  }
});

const onIcalToggle = async () => {
  if (icalUrl.value !== null) {
    const result = await calendarUseCase.disableIcalSubscription();
    if (!isResultError(result)) {
      icalUrl.value = null;
      // The URLs issued with tags are no longer valid because the token has been revoked.
      issuedIcalTagIds.value = [];
    } else if (result instanceof NetworkError) {
      displayToast(
        "ネットワークエラーが発生しました。お使いの端末がインターネットに接続されているか、今一度確認ください。",
        { type: "danger" }
      );
    } else if (result instanceof InternalServerError) {
      displayToast("サーバーエラーが発生しました。", { type: "danger" });
    }
  } else {
    const result = await calendarUseCase.enableIcalSubscription();
    if (!isResultError(result) && "url" in result) {
      icalUrl.value = result.url;
    } else if (result instanceof NetworkError) {
      displayToast(
        "ネットワークエラーが発生しました。お使いの端末がインターネットに接続されているか、今一度確認ください。",
        { type: "danger" }
      );
    } else if (result instanceof InternalServerError) {
      displayToast("サーバーエラーが発生しました。", { type: "danger" });
    }
  }
};

const copyIcalUrl = async () => {
  if (!icalUrl.value) return;
  try {
    await navigator.clipboard.writeText(icalUrl.value);
    displayToast("URLをコピーしました", { type: "primary" });
  } catch {
    displayToast("コピーに失敗しました", { type: "danger" });
  }
};

/** ical subscription filtered by tags */
const [isIcalTagModalVisible, openIcalTagModal, closeIcalTagModal] = useSwitch(
  false
);
const icalTagModalStep = ref<"select" | "issued">("select");
const displayIcalTags = ref<DisplayIcalTag[]>([]);
const selectedIcalTagIds = ref<string[]>([]);
const issuedIcalTagIds = ref<string[]>([]);
const isIssuedIcalUrlCopied = ref(false);

const isAllIcalTagsSelected = computed(
  () =>
    displayIcalTags.value.length > 0 &&
    selectedIcalTagIds.value.length === displayIcalTags.value.length
);

// Keep the display order so that the same selection always gives the same URL.
const issuedIcalTags = computed(() =>
  displayIcalTags.value.filter(({ id }) => issuedIcalTagIds.value.includes(id))
);

const issuedIcalUrl = computed(() =>
  icalUrl.value
    ? getIcalUrlWithTags(
        icalUrl.value,
        issuedIcalTags.value.map(({ id }) => id)
      )
    : ""
);

const onClickIcalTagCustomize = async () => {
  const [tags, registeredCourses] = await Promise.all([
    timetableUseCase.listTags(),
    // The subscription URL has no year parameter, so the server exports the courses of the current academic year.
    timetableUseCase.listRegisteredCourses(currentAcademicYear),
  ]);

  if (isResultError(tags) || isResultError(registeredCourses)) {
    const error = isResultError(tags) ? tags : registeredCourses;
    if (error instanceof UnauthenticatedError) {
      displayToast(
        "ログインの確認に失敗しました。お手数ですが、再度ログインした上でお試しいただけますと幸いです。",
        { type: "danger" }
      );
      router.push("/login");
    } else if (error instanceof NetworkError) {
      displayToast(
        "ネットワークエラーが発生しました。お使いの端末がインターネットに接続されているか、今一度確認ください。",
        { type: "danger" }
      );
    } else if (error instanceof InternalServerError) {
      displayToast("サーバーエラーが発生しました。", { type: "danger" });
    }
    return;
  }

  displayIcalTags.value = getDisplayIcalTags(
    registeredCourses,
    tags.sort((tagA, tagB) => tagA.order - tagB.order)
  );
  selectedIcalTagIds.value = issuedIcalTags.value.map(({ id }) => id);
  icalTagModalStep.value = "select";
  openIcalTagModal();
};

const toggleIcalTag = (tagId: string) => {
  selectedIcalTagIds.value = selectedIcalTagIds.value.includes(tagId)
    ? selectedIcalTagIds.value.filter((id) => id !== tagId)
    : [...selectedIcalTagIds.value, tagId];
};

const toggleAllIcalTags = () => {
  selectedIcalTagIds.value = isAllIcalTagsSelected.value
    ? []
    : displayIcalTags.value.map(({ id }) => id);
};

const issueIcalUrl = () => {
  if (selectedIcalTagIds.value.length === 0) return;
  issuedIcalTagIds.value = displayIcalTags.value
    .filter(({ id }) => selectedIcalTagIds.value.includes(id))
    .map(({ id }) => id);
  isIssuedIcalUrlCopied.value = false;
  icalTagModalStep.value = "issued";
};

const copyIssuedIcalUrl = async () => {
  try {
    await navigator.clipboard.writeText(issuedIcalUrl.value);
    isIssuedIcalUrlCopied.value = true;
  } catch {
    displayToast("コピーに失敗しました", { type: "danger" });
  }
};

const backToIcalTagSelect = () => {
  icalTagModalStep.value = "select";
};

/** logout */
const logout = () => {
  redirectToUrl(getLogoutUrl());
};

/** display year */
const autoOption = "自動(現在の年度)";

const yearOptions: string[] = [
  autoOption,
  ...academicYears.map((year) => `${year}年度`).reverse(),
];

const selectedYearOption = computed<string>(() =>
  setting.value.displayYear === 0
    ? autoOption
    : `${setting.value.displayYear}年度`
);

const updateSelectedYearOption = async (option: string) => {
  const year: number = option === autoOption ? 0 : Number(option.slice(0, 4));
  await updateSetting({ displayYear: year });
};

const openNotificationSetting = () => {
  // apply setTimeout for animation
  setTimeout(() => {
    if (isiOS())
      window.webkit?.messageHandlers?.iPhoneSettings?.postMessage("");
    else window.android?.openSettings();
  }, 300);
};

/** Account Delete modal */
const [
  isAccountDeletionModalVisible,
  openAccountDeletionModal,
  closeAccountDeletionModal,
] = useSwitch(false);

const onClickAccountDeleteModel = () => {
  openAccountDeletionModal();
};
const confirmDeleteAccount = async () => {
  const deleteUserResult = await authUseCase.deleteUser();
  if (!isResultError(deleteUserResult)) {
    closeAccountDeletionModal();
    displayToast(
      "アカウントの削除に成功しました。今までのご利用、誠にありがとうございました。",
      {
        type: "primary",
      }
    );
    router.push("/login");
  } else {
    const error = deleteUserResult;
    console.log(error);
    if (error instanceof UnauthenticatedError) {
      displayToast(
        "ログインの確認に失敗しました。お手数ですが、再度ログインした上でお試しいただけますと幸いです。",
        { type: "danger" }
      );
      router.push("/login");
    } else if (error instanceof NetworkError) {
      displayToast(
        "ネットワークエラーが発生しました。お使いの端末がインターネットに接続されているか、今一度確認ください。",
        { type: "danger" }
      );
    } else if (error instanceof InternalServerError) {
      displayToast("サーバーエラーが発生しました。", { type: "danger" });
    }
  }
};
</script>

<style scoped lang="scss">
@import "~/ui/styles";
.settings {
  @include max-width;
}

.main {
  margin-top: $spacing-5;
  &__contents {
    height: calc(#{$vh} - 8rem);
    overflow-y: auto;
    padding: 0 1.2rem;
    margin: 0 -1.2rem;
  }
  &__content {
    display: flex;
    align-items: center;
    padding: 1.2rem 0;
    color: getColor(--color-text-main);
    font-weight: 500;
    & .switch,
    & .button {
      margin: 0 0 0 auto;
    }
    &--dropdown {
      display: grid;
      gap: 0.8rem;
      padding: 2rem 0;
      p {
        line-height: $single-line;
        font-weight: 500;
        color: getColor(--color-text-main);
      }
    }
    &--ical {
      display: flex;
      align-items: center;
      padding: 1.2rem 0;
      flex-wrap: wrap;
      p {
        font-weight: 500;
        color: getColor(--color-text-main);
      }
      & .switch {
        margin: 0 0 0 auto;
      }
      .ical-detail {
        display: flex;
        flex-direction: column;
        gap: 0.8rem;
        margin-top: 0.8rem;
      }
      .ical-description {
        line-height: $single-line;
        color: getColor(--color-text-sub);
        font-weight: 400;
      }
      .ical-url-row {
        display: flex;
        gap: 1.6rem;
        align-items: center;
      }
      .ical-url-input {
        flex: 1;
        width: 0;
        color: getColor(--color-text-main);
        background: getColor(--color-background-sub);
        text-overflow: ellipsis;
      }
      .ical-tag-entry {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        gap: $spacing-2;
        &__link {
          &:hover,
          &:active {
            box-shadow: none;
          }
          &:hover :deep(.tertiary-button__text) {
            text-decoration: underline;
            text-underline-offset: 0.3rem;
          }
          :deep(.tertiary-button__icon) {
            font-size: $font-large;
          }
        }
        &__label {
          font-size: $font-small;
          font-weight: 400;
          color: getColor(--color-text-sub);
        }
      }
      .ical-cautions {
        margin-top: 0.8rem;
        li {
          list-style: disc inside;
          margin-bottom: 0.8rem;
          font-weight: 400;
        }
      }
    }
    &--account {
      display: flex;
      padding: 1.2rem 0;
      color: getColor(--color-text-main);
      font-weight: 500;
      .account-btns {
        margin: 0 0 0 auto;
        display: flex;
        flex-direction: column;
        gap: 2rem;
        & .button {
          margin: 0 0 0 auto;
        }
      }
    }
  }
}

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
  &__tag {
    @include button-cursor;
    :deep(.tag-list-content__credit) {
      font-size: $font-small;
      font-weight: 400;
      color: getColor(--color-text-sub);
    }
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
</style>
