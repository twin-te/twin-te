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
        <template v-if="isAuthenticated && connectedProviders">
          <div class="main__content">
            <p>ログイン方法</p>
            <span class="provider-count"
              >{{ connectedProviders.length }} /
              {{ providers.length }} 連携中</span
            >
          </div>
          <Card class="provider-card">
            <div
              v-for="(provider, index) in displayedProviders"
              :key="provider"
              class="provider"
            >
              <div v-if="index > 0" class="provider-card__divider"></div>
              <div class="provider__row">
                <div :class="['provider__mark', `provider__mark--${provider}`]">
                  <img :src="providerMarkMap[provider]" alt="" />
                </div>
                <div class="provider__text">
                  <span class="provider__name">{{
                    providerMap[provider]
                  }}</span>
                  <span class="provider__status">{{
                    connectedProviders.includes(provider)
                      ? "連携済み"
                      : "未連携"
                  }}</span>
                </div>
                <div
                  v-if="isLocked(provider)"
                  class="provider__lock provider__lock--inline"
                >
                  <span class="material-icons">lock</span
                  >ログイン方法が1つだけのため解除できません
                </div>
                <Button
                  v-if="connectedProviders.includes(provider)"
                  class="provider__button"
                  size="small"
                  color="base"
                  :state="isLocked(provider) ? 'disabled' : 'default'"
                  :pauseActiveStyle="false"
                  @click="openDisconnectionModal(provider)"
                >
                  <span class="material-icons provider__button-icon--danger"
                    >link_off</span
                  ><span class="provider__button-label--danger">解除</span>
                </Button>
                <Button
                  v-else-if="canConnect"
                  class="provider__button"
                  size="small"
                  color="base"
                  :pauseActiveStyle="false"
                  @click="connect(provider)"
                >
                  <span class="material-icons provider__button-icon--liner"
                    >add</span
                  ><span class="provider__button-label--liner">接続</span>
                </Button>
              </div>
              <div
                v-if="isLocked(provider)"
                class="provider__lock provider__lock--below"
              >
                <span class="material-icons">lock</span
                >ログイン方法が1つだけのため解除できません
              </div>
            </div>
            <div v-if="!canConnect" class="provider">
              <div class="provider-card__divider"></div>
              <div class="provider__row">
                <div class="provider__mark provider__mark--add">
                  <span class="material-icons">add</span>
                </div>
                <div class="provider__text">
                  <span class="provider__name">ログイン方法を追加</span>
                  <span class="provider__status"
                    >ブラウザ版の設定画面から追加できます</span
                  >
                </div>
              </div>
            </div>
          </Card>
        </template>
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
    <Modal
      v-if="providerToDisconnect"
      class="provider-disconnect-modal"
      size="small"
      @click="closeDisconnectionModal"
    >
      <template #title
        >{{
          providerMap[providerToDisconnect]
        }}との連携を解除しますか？</template
      >
      <template #contents>
        <p class="modal__text">
          解除すると、選択した{{
            providerMap[providerToDisconnect]
          }}アカウントではTwin:teにログインできなくなります。他のログイン方法を使って引き続きTwin:teを使用できます。
        </p>
      </template>
      <template #button>
        <Button
          size="medium"
          layout="fill"
          color="base"
          @click="closeDisconnectionModal"
          >キャンセル</Button
        >
        <Button
          size="medium"
          layout="fill"
          color="danger"
          @click="confirmDisconnect"
          >解除</Button
        >
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { useHead } from "@vueuse/head";
import { computed, onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import {
  InternalServerError,
  isResultError,
  NetworkError,
  UnauthenticatedError,
} from "~/domain/error";
import { Provider } from "~/domain/user";
import { academicYears } from "~/domain/year";
import { providerMap } from "~/presentation/presenters/provider";
import logoX from "~/ui/assets/login-page/logo-x.png";
import markAppleWhite from "~/ui/assets/login-page/mark-apple-white.svg";
import markGoogle from "~/ui/assets/login-page/mark-google.svg";
import Card from "~/ui/components/Card.vue";
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
import { getConnectUrl, getLogoutUrl, redirectToUrl } from "../url";

const router = useRouter();
const route = useRoute();
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

/** login providers */
const providers: Provider[] = ["google", "apple", "twitter"];

const providerMarkMap: Record<Provider, string> = {
  google: markGoogle,
  apple: markAppleWhite,
  twitter: logoX,
};

const connectedProviders = ref<Provider[] | undefined>(undefined);

onMounted(async () => {
  if (!isAuthenticated.value) {
    return;
  }
  const result = await authUseCase.getMe();
  if (!isResultError(result)) {
    connectedProviders.value = result.providers;
  }
});

// The Android app ignores the redirect url of Google and always returns to the top page (twin-te/twin-te#482),
// so connecting is not provided in the Android app until the app is updated.
const canConnect = !isAndroid();

// Only the connected providers are displayed if connecting is not provided.
const displayedProviders = computed<Provider[]>(() =>
  canConnect
    ? providers
    : providers.filter((provider) =>
        connectedProviders.value?.includes(provider)
      )
);

// The last provider cannot be disconnected, since the user has at least one authentication.
const isLocked = (provider: Provider): boolean =>
  connectedProviders.value?.length === 1 &&
  connectedProviders.value.includes(provider);

const connect = (provider: Provider) => {
  redirectToUrl(getConnectUrl(provider));
};

/** result of connecting a provider, which is passed as query by the back end */
const getConnectResultMessage = (result: string): string | undefined => {
  switch (result) {
    case "connected":
      return "連携しました。";
    case "already_connected":
      return "このアカウントは既に連携済みです。";
  }
  return undefined;
};

const getConnectErrorMessage = (error: string): string => {
  switch (error) {
    case "cancelled":
      return "連携がキャンセルされました。";
    case "unauthenticated":
      return "ログインの確認に失敗しました。お手数ですが、再度ログインした上でお試しいただけますと幸いです。";
    case "already_used_by_another_user":
      return "このアカウントは既に別のTwin:teアカウントに連携されているため、連携できません。";
    case "provider_already_connected":
      return "同じサービスの別のアカウントが既に連携されています。連携を解除してからお試しください。";
  }
  return "連携に失敗しました。お手数ですが、再度お試しください。";
};

onMounted(() => {
  const connectResult = route.query.connect_result?.toString();
  const connectError = route.query.connect_error?.toString();
  if (connectResult === undefined && connectError === undefined) {
    return;
  }

  if (connectError !== undefined) {
    displayToast(getConnectErrorMessage(connectError), { type: "danger" });
  } else if (connectResult !== undefined) {
    const message = getConnectResultMessage(connectResult);
    if (message) displayToast(message, { type: "primary" });
  }

  // prevent the message from being displayed again on reload
  router.replace({
    query: {
      ...route.query,
      connect_result: undefined,
      connect_error: undefined,
    },
  });
});

/** disconnect provider */
const providerToDisconnect = ref<Provider | undefined>(undefined);

const openDisconnectionModal = (provider: Provider) => {
  providerToDisconnect.value = provider;
};

const closeDisconnectionModal = () => {
  providerToDisconnect.value = undefined;
};

const confirmDisconnect = async () => {
  const provider = providerToDisconnect.value;
  if (provider === undefined) return;

  const result = await authUseCase.deleteUserAuthentication(provider);
  if (!isResultError(result)) {
    closeDisconnectionModal();
    connectedProviders.value = connectedProviders.value?.filter(
      (connectedProvider) => connectedProvider !== provider
    );
    displayToast(`${providerMap[provider]}との連携を解除しました。`, {
      type: "primary",
    });
  } else if (result instanceof UnauthenticatedError) {
    displayToast(
      "ログインの確認に失敗しました。お手数ですが、再度ログインした上でお試しいただけますと幸いです。",
      { type: "danger" }
    );
    router.push("/login");
  } else if (result instanceof NetworkError) {
    displayToast(
      "ネットワークエラーが発生しました。お使いの端末がインターネットに接続されているか、今一度確認ください。",
      { type: "danger" }
    );
  } else if (result instanceof InternalServerError) {
    displayToast("サーバーエラーが発生しました。", { type: "danger" });
  }
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

.provider-count {
  margin-left: auto;
  font-size: $font-small;
  color: getColor(--color-text-sub);
}

.main .provider-card {
  padding: $spacing-1 $spacing-5;
  margin-bottom: $spacing-3;
  &__divider {
    height: 0.2rem;
    border-radius: 0.2rem;
    box-shadow: $shadow-concave;
  }
}

.provider {
  display: flex;
  flex-direction: column;
  &__row {
    display: flex;
    align-items: center;
    gap: $spacing-3;
    padding: 1.4rem 0;
  }
  &__mark {
    @include center-flex;
    flex-shrink: 0;
    width: 3.4rem;
    height: 3.4rem;
    border-radius: 50%;
    box-shadow: $shadow-drop;
    img {
      width: 1.5rem;
      height: 1.5rem;
    }
    &--google {
      background: #ffffff;
    }
    &--apple,
    &--twitter {
      background: #000000;
    }
    &--add {
      background: var(--base-liner);
      box-shadow: $shadow-convex;
      .material-icons {
        font-size: 2rem;
        @include text-liner;
      }
    }
  }
  &__text {
    display: flex;
    flex-direction: column;
    gap: 0.2rem;
    flex: 1;
    min-width: 0;
  }
  &__name {
    color: getColor(--color-text-main);
    font-weight: 500;
  }
  &__status {
    @include ellipsis;
    font-size: $font-small;
    font-weight: 400;
    color: getColor(--color-text-sub);
  }
  &__lock {
    display: flex;
    align-items: center;
    gap: $spacing-1;
    white-space: nowrap;
    font-size: $font-small;
    font-weight: 400;
    color: getColor(--color-text-sub);
    .material-icons {
      font-size: 1.4rem;
      color: getColor(--color-button-gray);
    }
    // The reason is displayed below the row in portrait, and in the row in landscape.
    &--below {
      margin: -0.6rem 0 1.4rem;
      line-height: $multi-line;
      @include landscape {
        display: none;
      }
    }
    &--inline {
      display: none;
      @include landscape {
        display: flex;
      }
    }
  }
  &__row &__button {
    display: inline-flex;
    align-items: center;
    flex-shrink: 0;
    gap: 0.6rem;
    height: 2.8rem;
    padding: 0 1.4rem 0 1rem;
    background: var(--base-liner);
    .material-icons {
      font-size: 1.8rem;
    }
    span {
      // Button ignores the click whose target is not the button itself.
      pointer-events: none;
    }
    &:active:not(.--disabled) span {
      color: getColor(--color-white);
      @include void-text-liner;
    }
  }
  &__button-icon--danger,
  &__button-label--danger {
    color: getColor(--color-danger);
  }
  &__button-icon--liner,
  &__button-label--liner {
    @include text-liner;
  }
}

.provider-disconnect-modal .modal {
  .button {
    width: calc(50% - 0.6rem);
    &:first-child {
      margin-right: 0.6rem;
    }
    &:last-child {
      margin-left: 0.6rem;
    }
  }
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
