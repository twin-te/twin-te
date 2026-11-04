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
            <input
              v-model="icalUrl"
              type="text"
              readonly
              class="ical-url-text"
            />
            <div class="ical-actions">
              <Button
                class="ical-register-button"
                size="small"
                color="primary"
                :pauseActiveStyle="false"
                @click="openCalendarRegisterModal"
              >
                <span class="material-icons ical-register-button__icon"
                  >calendar_today</span
                >
                <span class="ical-register-button__label">登録</span>
              </Button>
              <Button
                class="button"
                size="small"
                color="base"
                :pauseActiveStyle="false"
                @click="copyIcalUrl"
                >コピー</Button
              >
            </div>
            <div class="ical-cautions">
              <p class="ical-cautions__title">注意事項</p>
              <ul class="ical-cautions__list">
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
      v-if="isCalendarRegisterModalVisible"
      size="auto"
      class="calendar-register-modal"
      @click="closeCalendarRegisterModal"
    >
      <template #title>カレンダーに登録</template>
      <template #contents>
        <ul class="register-targets">
          <li v-for="target in registerTargets" :key="target.id">
            <button
              class="register-targets__item"
              :class="{
                'register-targets__item--disabled': !target.isAvailable,
              }"
              :disabled="!target.isAvailable"
              @click="target.open"
            >
              <span class="register-targets__title">{{ target.title }}</span>
              <span v-if="target.note" class="register-targets__note">
                <span class="material-icons register-targets__note-icon"
                  >desktop_windows</span
                >{{ target.note }}
              </span>
              <span
                v-else-if="target.description"
                class="register-targets__desc"
                >{{ target.description }}</span
              >
            </button>
          </li>
        </ul>
      </template>
      <template #button>
        <Button
          size="medium"
          layout="fill"
          color="base"
          @click="closeCalendarRegisterModal"
          >閉じる</Button
        >
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
import { academicYears } from "~/domain/year";
import Dropdown from "~/ui/components/Dropdown.vue";
import IconButton from "~/ui/components/IconButton.vue";
import Modal from "~/ui/components/Modal.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import ToggleSwitch from "~/ui/components/ToggleSwitch.vue";
import { useSwitch } from "~/ui/hooks/useSwitch";
import { isAndroid, isiOS, isMobile } from "~/ui/ua";
import { authUseCase, calendarUseCase } from "~/usecases";
import Button from "../components/Button.vue";
import { useAuth, useSetting, useToast } from "../store";
import { getLogoutUrl, redirectToUrl } from "../url";

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

const [
  isCalendarRegisterModalVisible,
  openCalendarRegisterModal,
  closeCalendarRegisterModal,
] = useSwitch(false);

/** Google Calendar cannot subscribe to a URL from the iOS app, so hide it there. */
const canRegisterGoogle = computed(() => !isiOS());

/** Apple Calendar (webcal) cannot be used on the Android app, so hide it there. */
const canRegisterApple = computed(() => !isAndroid());

const toWebcalUrl = (url: string) => url.replace(/^https?:\/\//, "webcal://");

/**
 * Build a click handler that opens a calendar registration URL in a new tab.
 * `buildUrl` receives the current ical URL and returns the URL to open;
 * `enabled` optionally guards whether the registration is allowed.
 */
const createRegisterHandler = (
  buildUrl: (icalUrl: string) => string,
  enabled?: () => boolean
) => () => {
  if (!icalUrl.value || (enabled && !enabled())) return;
  window.open(buildUrl(icalUrl.value), "_blank");
  closeCalendarRegisterModal();
};

const openGoogleCalendar = createRegisterHandler(
  (url) =>
    `https://calendar.google.com/calendar/r?cid=${encodeURIComponent(
      toWebcalUrl(url)
    )}`,
  () => canRegisterGoogle.value
);

const openAppleCalendar = createRegisterHandler((url) => toWebcalUrl(url));

const openOutlookCalendar = createRegisterHandler(
  (url) =>
    `https://outlook.office.com/calendar/addcalendar?name=${encodeURIComponent(
      "Twin:te"
    )}&url=${encodeURIComponent(url)}`
);

const openIcsFile = createRegisterHandler((url) => url);

type RegisterTarget = {
  id: string;
  title: string;
  description?: string;
  note?: string;
  isAvailable: boolean;
  open: () => void;
};

/** モーダルに並べる登録先。 */
const registerTargets = computed<RegisterTarget[]>(() => [
  {
    id: "google",
    title: "Googleカレンダー",
    note: canRegisterGoogle.value ? undefined : "PCで操作してください",
    isAvailable: canRegisterGoogle.value,
    open: openGoogleCalendar,
  },
  // Apple Calendar (webcal) cannot be used on the Android app, so omit the item there.
  ...(canRegisterApple.value
    ? [
        {
          id: "apple",
          title: "Appleカレンダー",
          description: "iOS / macOS",
          isAvailable: true,
          open: openAppleCalendar,
        },
      ]
    : []),
  {
    id: "outlook",
    title: "Outlook",
    description: "Microsoft 365",
    isAvailable: true,
    open: openOutlookCalendar,
  },
  {
    id: "ics",
    title: ".icsファイル",
    description: "手動インポート",
    isAvailable: true,
    open: openIcsFile,
  },
]);

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
        width: 100%;
      }
      .ical-description {
        line-height: $single-line;
        color: getColor(--color-text-sub);
        font-weight: 400;
      }
      .ical-actions {
        display: flex;
        justify-content: flex-end;
        align-items: center;
        gap: $spacing-2;
      }
      .ical-register-button {
        display: flex;
        align-items: center;
        gap: $spacing-2;
      }
      .ical-register-button__icon,
      .ical-register-button__label {
        // Button はクリック対象が button 要素のときだけ click を emit するため、
        // 中身がイベントを奪わないようにする
        pointer-events: none;
      }
      .ical-register-button__icon {
        font-size: $font-medium;
        line-height: 1;
      }
      .ical-url-text {
        width: 100%;
        padding: 0;
        border: none;
        background: transparent;
        color: getColor(--color-text-main);
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        font-variant-numeric: tabular-nums;
      }
      .ical-cautions {
        margin-top: 0.8rem;
      }
      .ical-cautions__title {
        font-weight: 700;
        color: getColor(--color-text-main);
        margin-bottom: 0.4rem;
      }
      .ical-cautions__list {
        padding-left: 1.8rem;
        color: getColor(--color-text-sub);
        line-height: 1.9;
        li {
          list-style: disc;
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
.calendar-register-modal {
  .register-targets {
    display: flex;
    flex-direction: column;
    list-style: none;
    margin: 0;
    padding: 0;
  }
  .register-targets__item {
    @include button-cursor;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 0.2rem;
    width: 100%;
    padding: $spacing-3 $spacing-4;
    background: transparent;
    border: none;
    border-radius: $radius-2;
    text-align: left;
    transition: background 0.18s ease;
    &:hover {
      background: getColor(--color-base);
    }
    &:active {
      background: getColor(--color-base);
      opacity: 0.85;
    }
    &--disabled {
      cursor: not-allowed;
      opacity: 0.55;
      &:hover,
      &:active {
        background: transparent;
        opacity: 0.55;
      }
    }
  }
  .register-targets__title {
    font-size: $font-medium;
    font-weight: 700;
    color: getColor(--color-text-main);
  }
  .register-targets__desc {
    font-size: $font-small;
    color: getColor(--color-text-sub);
    font-weight: 400;
  }
  .register-targets__note {
    display: inline-flex;
    align-items: center;
    gap: 0.4rem;
    font-size: $font-small;
    color: getColor(--color-text-sub);
    font-weight: 400;
  }
  .register-targets__note-icon {
    font-size: $font-small;
    line-height: 1;
  }
}
</style>
