<template>
  <div class="login">
    <div class="login__rectangle-logo">
      <img v-if="setting.darkMode" src="../assets//twintelogo-darkmode.svg" />
      <img v-else src="../assets/twintelogo-color.svg" />
    </div>
    <div class="login__main">
      <div class="main__square-logo">
        <img
          v-if="setting.darkMode"
          src="../assets/twintelogo-darkmode-b.svg"
        />
        <img v-else src="../assets/twintelogo-color-b.svg" />
      </div>
      <div class="main__head">ログイン方法を選択</div>
      <div class="main__provider">
        <button @click="login('apple')">
          <img
            class="main__login-button"
            src="../assets/login-page/login-apple.png"
            alt="appleでログイン"
          />
        </button>
        <button @click="login('twitter')" class="main__login-button main__btn-x">
          <img src="../assets/login-page/logo-x.png" alt="Xのロゴマーク">
          <span>Xでサインイン</span>
        </button>
        <button @click="login('google')">
          <img
            class="main__login-button"
            src="../assets/login-page/login-google.svg"
            alt="googleでログイン"
          />
        </button>
      </div>
      <div class="main__note">
        ※ ログインした場合、<a href="https://www.twinte.net/terms">利用規約</a>
        及び
        <a href="https://www.twinte.net/policy">プライバシーポリシー</a
        >に同意したものとします。<br />
        ※Twin:teにログインしたことがない場合は、選択したログイン方法と紐づいたアカウントが自動で作成されます。
      </div>
      <img
        v-if="setting.darkMode"
        class="main__prev-button"
        src="../assets/login-page/button-prev_dark.png"
        alt="元のページへ戻る"
        @click="$router.back()"
      />
      <img
        v-else
        class="main__prev-button"
        src="../assets/login-page/button-prev.svg"
        alt="元のページへ戻る"
        @click="$router.back()"
      />
    </div>
  </div>
  <GrayFilter
    v-show="clicked"
    class="layout__grayfilter"
    @click="() => {}"
  ></GrayFilter>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { Provider } from "~/domain/user";
import GrayFilter from "~/ui/components/GrayFilter.vue";
import { getLoginUrl, redirectToUrl } from "~/ui/url";
import { useSetting } from "../store";

export default defineComponent({
  components: { GrayFilter },
  setup: () => {
    const router = useRouter();
    const route = useRoute();
    const redirectUrl = route.query.redirectUrl?.toString() as string;
    const { setting } = useSetting();
    const clicked = ref(false);
    const login = (provider: Provider) => {
      clicked.value = true;
      redirectToUrl(getLoginUrl(provider, redirectUrl));
    };
    return { router, setting, clicked, login };
  },
});
</script>

<style lang="scss" scoped>
@import "~/ui/styles";

.login {
  @include center-flex(column);
  padding-top: $safe-area-top;
  width: 100%;
  height: 100vh;
  &__rectangle-logo {
    display: none;
    position: absolute;
    top: $spacing-9;
    left: $spacing-9;
    @include landscape {
      display: block;
    }
    img {
      height: 36px;
    }
  }
  &__main {
    width: 287px;
    @include center-flex(column);
    @include landscape {
      align-items: flex-start;
    }
  }
}

.main {
  &__square-logo {
    display: block;
    margin-bottom: $spacing-12;
    img {
      height: 85px;
    }
    @include landscape {
      display: none;
    }
  }
  &__head {
    @include text-main;
    margin-bottom: $spacing-10;
    text-align: center;
    font-size: $font-maximum;
  }
  &__provider {
    display: flex;
    flex-direction: column;
    gap: $spacing-5;
    margin-bottom: $spacing-6;
  }
  &__login-button {
    @include button-cursor;
    width: 100%;
    box-shadow: $shadow-convex;
    border-radius: $radius-button;
    display: flex;
    $icon-size: 18px;
    img{
      width: $icon-size;
      height: $icon-size;
    }
    span{
      width: calc(100% - $icon-size);
      font-size: $font-large;
    }
  }
  &__btn-x{
    color: white;
    background: black;
    padding: 14px 24px;
  }
  &__note {
    @include text-sub-discription;
    margin-bottom: $spacing-5;
    a {
      color: getColor(--color-text-link);
    }
  }
  &__prev-button {
    @include button-cursor;
    display: block;
    height: 70px;
    @include landscape {
      margin-left: -12px;
    }
  }
}
</style>
