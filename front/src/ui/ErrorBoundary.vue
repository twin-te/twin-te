<script setup lang="ts">
import { onErrorCaptured, ref } from "vue";
import { useRouter } from "vue-router";
import {
  InternalServerError,
  NetworkError,
  UnauthenticatedError,
} from "~/domain/error";
import Button from "./components/Button.vue";

const router = useRouter();

const errorMessage = ref<string>("");

const errorDetail = ref<
  | {
      messages: string[];
      buttonText: string;
      onClickButton: () => void;
    }
  | undefined
>(undefined);

onErrorCaptured((error) => {
  console.log("onErrorCaptured");
  console.log(error);

  if (error instanceof UnauthenticatedError) {
    errorDetail.value = {
      messages: ["未認証です。ログインして下さい。"],
      buttonText: "ログイン",
      onClickButton: () => {
        errorDetail.value = undefined;
        router.push("/login");
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
        errorDetail.value = undefined;
        location.reload();
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
        errorDetail.value = undefined;
        location.reload();
      },
    };
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
      errorDetail.value = undefined;
      location.reload();
    },
  };
  errorMessage.value = error.message;
  return true;
});
</script>

<template>
  <div v-if="errorDetail" class="error-boundary">
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
