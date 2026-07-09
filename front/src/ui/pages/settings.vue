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
            <div ref="registerRef" class="ical-register">
              <Button
                class="ical-register__toggle"
                size="medium"
                layout="fill"
                color="primary"
                :pauseActiveStyle="false"
                :state="isRegisterMenuOpen ? 'active' : 'default'"
                @click="toggleRegisterMenu"
              >
                <span class="material-icons ical-register__toggle-icon"
                  >event</span
                >
                <span class="ical-register__toggle-label">登録する</span>
                <span class="material-icons ical-register__toggle-chevron">{{
                  isRegisterMenuOpen ? "expand_less" : "expand_more"
                }}</span>
              </Button>
              <ul
                v-if="isRegisterMenuOpen"
                :class="[
                  'ical-register__menu',
                  `ical-register__menu--${menuDirection}`,
                ]"
              >
                <li>
                  <button
                    class="ical-register__menu-item"
                    @click="openGoogleCalendar"
                  >
                    <span class="ical-register__menu-title"
                      >Googleカレンダー</span
                    >
                  </button>
                </li>
                <li>
                  <button
                    class="ical-register__menu-item"
                    @click="openAppleCalendar"
                  >
                    <span class="ical-register__menu-title"
                      >Appleカレンダー</span
                    >
                    <span class="ical-register__menu-desc">iOS / macOS</span>
                  </button>
                </li>
                <li>
                  <button
                    class="ical-register__menu-item"
                    @click="openOutlookCalendar"
                  >
                    <span class="ical-register__menu-title">Outlook</span>
                    <span class="ical-register__menu-desc">Microsoft 365</span>
                  </button>
                </li>
                <li>
                  <button class="ical-register__menu-item" @click="openIcsFile">
                    <span class="ical-register__menu-title">.icsファイル</span>
                    <span class="ical-register__menu-desc">手動インポート</span>
                  </button>
                </li>
              </ul>
            </div>
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
import { computed, onBeforeUnmount, onMounted, ref, watch } from "vue";
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
import { isiOS, isMobile } from "~/ui/ua";
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

const isRegisterMenuOpen = ref(false);
const registerRef = ref<HTMLDivElement | null>(null);
const menuDirection = ref<"down" | "up">("down");
const MENU_ESTIMATED_HEIGHT = 280;
const toggleRegisterMenu = () => {
  if (!isRegisterMenuOpen.value && registerRef.value) {
    const rect = registerRef.value.getBoundingClientRect();
    const spaceBelow = window.innerHeight - rect.bottom;
    menuDirection.value =
      spaceBelow < MENU_ESTIMATED_HEIGHT && rect.top > spaceBelow
        ? "up"
        : "down";
  }
  isRegisterMenuOpen.value = !isRegisterMenuOpen.value;
};
const closeRegisterMenu = () => {
  isRegisterMenuOpen.value = false;
};

const handleOutsideClick = (event: MouseEvent) => {
  if (registerRef.value && !registerRef.value.contains(event.target as Node)) {
    closeRegisterMenu();
  }
};

watch(isRegisterMenuOpen, (open) => {
  if (open) {
    document.addEventListener("click", handleOutsideClick);
  } else {
    document.removeEventListener("click", handleOutsideClick);
  }
});

onBeforeUnmount(() => {
  document.removeEventListener("click", handleOutsideClick);
});

const openGoogleCalendar = () => {
  if (!icalUrl.value) return;
  const url = `https://calendar.google.com/calendar/r?cid=${encodeURIComponent(
    icalUrl.value
  )}`;
  window.open(url, "_blank");
  closeRegisterMenu();
};

const openAppleCalendar = () => {
  if (!icalUrl.value) return;
  const url = icalUrl.value.replace(/^https?:\/\//, "webcal://");
  window.open(url, "_blank");
  closeRegisterMenu();
};

const openOutlookCalendar = () => {
  if (!icalUrl.value) return;
  const url = `https://outlook.office.com/calendar/addcalendar?name=${encodeURIComponent(
    "Twin:te"
  )}&url=${encodeURIComponent(icalUrl.value)}`;
  window.open(url, "_blank");
  closeRegisterMenu();
};

const openIcsFile = () => {
  if (!icalUrl.value) return;
  window.open(icalUrl.value, "_blank");
  closeRegisterMenu();
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
      .ical-register {
        width: 100%;
        position: relative;
        margin-bottom: $spacing-3;
      }
      .ical-register__toggle {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: $spacing-2;
        width: 100%;
        position: relative;
      }
      .ical-register__toggle-icon,
      .ical-register__toggle-chevron {
        font-size: $font-large;
        line-height: 1;
        pointer-events: none;
      }
      .ical-register__toggle-label {
        pointer-events: none;
      }
      .ical-register__toggle-chevron {
        position: absolute;
        right: $spacing-4;
      }
      .ical-register__menu {
        list-style: none;
        margin: 0;
        padding: $spacing-2 0;
        display: flex;
        flex-direction: column;
        background: getColor(--color-white);
        border-radius: $radius-3;
        box-shadow: $shadow-convex;
        overflow: hidden;
        position: absolute;
        left: 0;
        right: 0;
        z-index: 10;
      }
      .ical-register__menu--down {
        top: calc(100% + #{$spacing-2});
      }
      .ical-register__menu--up {
        bottom: calc(100% + #{$spacing-2});
      }
      .ical-register__menu-item {
        @include button-cursor;
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        gap: 0.2rem;
        width: 100%;
        padding: $spacing-3 $spacing-5;
        background: transparent;
        border: none;
        text-align: left;
        transition: background 0.18s ease;
        &:hover {
          background: getColor(--color-base);
        }
        &:active {
          background: getColor(--color-base);
          opacity: 0.85;
        }
      }
      .ical-register__menu-title {
        font-size: $font-medium;
        font-weight: 700;
        color: getColor(--color-text-main);
        pointer-events: none;
      }
      .ical-register__menu-desc {
        font-size: $font-small;
        color: getColor(--color-text-sub);
        font-weight: 400;
        pointer-events: none;
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
</style>
