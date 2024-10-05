<script setup lang="ts">
import * as Sentry from "@sentry/vue";
import { onErrorCaptured, ref } from "vue";
import { useRouter } from "vue-router";
import {
  InternalServerError,
  NetworkError,
  UnauthenticatedError,
} from "~/domain/error";
import Button from "../components/Button.vue";
import IconButton from "../components/IconButton.vue";
import PageHeader from "../components/PageHeader.vue";

const router = useRouter();

const errorDetail = ref<
  | {
      messages: string[];
      buttonText: string;
      onClickButton: () => void;
    }
  | undefined
>(undefined);

onErrorCaptured((error) => {
  console.log("captured the following error in ErrorBoundary");
  console.log(error);

  if (error instanceof UnauthenticatedError) {
    errorDetail.value = {
      messages: ["未認証です。ログインして下さい。"],
      buttonText: "ログイン",
      onClickButton: async () => {
        await router.push("/login");
        errorDetail.value = undefined;
      },
    };
    return true;
  }

  if (error instanceof NetworkError) {
    errorDetail.value = {
      messages: [
        "ネットワークエラーが発生しました。",
        "通信状況をご確認下さい。",
      ],
      buttonText: "リロード",
      onClickButton: () => {
        location.reload();
        errorDetail.value = undefined;
      },
    };
    return true;
  }

  if (error instanceof InternalServerError) {
    errorDetail.value = {
      messages: [
        "申し訳ございません。サーバー内でエラーが発生しました。",
        "リロードして再度お試しください。",
        "改善されない場合は運営にお問い合わせ下さい。",
        error.message,
      ],
      buttonText: "リロード",
      onClickButton: () => {
        location.reload();
        errorDetail.value = undefined;
      },
    };

    Sentry.captureException(error);
    return true;
  }

  errorDetail.value = {
    messages: [
      "申し訳ございません。予期せぬエラーが発生しました。",
      "リロードして再度お試しください。",
      "改善されない場合は運営にお問い合わせ下さい。",
      error.message,
    ],
    buttonText: "リロード",
    onClickButton: () => {
      location.reload();
      errorDetail.value = undefined;
    },
  };

  Sentry.captureException(error);
  return true;
});

const onClickBackButton = () => {
  router.back();
  errorDetail.value = undefined;
};
</script>

<template>
  <div v-if="errorDetail" class="error-boundary">
    <PageHeader>
      <template #left-button-icon>
        <IconButton
          size="large"
          color="normal"
          icon-name="arrow_back"
          @click="onClickBackButton"
        ></IconButton>
      </template>
      <template #title></template>
    </PageHeader>
    <div class="error">
      <div class="error__icon material-icons">error_outline</div>
      <p
        v-for="(message, i) in errorDetail.messages"
        :key="i"
        class="error__message"
      >
        {{ message }}
      </p>
      <Button size="large" color="base" @click="errorDetail.onClickButton">
        {{ errorDetail.buttonText }}
      </Button>
    </div>
  </div>
  <slot v-else />
</template>

<style scoped lang="scss">
@import "~/ui/styles";

.error-boundary {
  width: 100%;
  height: 100vh;

  padding: $spacing-0 $spacing-4;
  @include landscape {
    padding: $spacing-0 $spacing-9;
  }

  background: var(--base-liner);
  color: var(--text-sub-light);
}

.error {
  @include center-asolute();
  @include center-flex(column);
  gap: 3.2rem;

  &__icon {
    font-size: 10rem;
    opacity: 0.2;
  }
  &__message {
    opacity: 0.7;
  }
}
</style>
