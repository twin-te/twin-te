<template>
  <div class="ical-settings">
    <PageHeader>
      <template #left-button-icon>
        <IconButton
          size="large"
          color="normal"
          icon-name="arrow_back"
          @click="$router.back()"
        ></IconButton>
      </template>
      <template #title>詳細設定</template>
    </PageHeader>
    <div class="main">
      <div class="main__contents">
        <IcalDetailSettings
          v-model="value"
          :tags="tags"
          @create-tag="onCreateTag"
        />
      </div>
      <div class="main__footer">
        <Button
          size="medium"
          layout="fill"
          color="primary"
          :pauseActiveStyle="false"
          :state="tags.length === 0 ? 'disabled' : 'default'"
          @click="onSave"
          >保存する</Button
        >
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useHead } from "@vueuse/head";
import { onMounted } from "vue";
import { useRouter } from "vue-router";
import {
  InternalServerError,
  isResultError,
  NetworkError,
  UnauthenticatedError,
} from "~/domain/error";
import Button from "~/ui/components/Button.vue";
import IcalDetailSettings from "~/ui/components/IcalDetailSettings.vue";
import IconButton from "~/ui/components/IconButton.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import { useIcalDetailSettings } from "~/ui/hooks/useIcalDetailSettings";
import { useToast } from "~/ui/store";

useHead({
  title: "Twin:te | カレンダー連携の詳細設定",
});

const router = useRouter();
const { displayToast } = useToast();

const { tags, value, load, save } = useIcalDetailSettings();

onMounted(load);

const onCreateTag = () => {
  router.push("/credit");
};

const onSave = async () => {
  const result = await save();
  if (!isResultError(result)) {
    displayToast("表示設定を保存しました", { type: "primary" });
    router.back();
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
</script>

<style scoped lang="scss">
@import "~/ui/styles";
.ical-settings {
  @include max-width;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.main {
  margin-top: $spacing-5;
  display: flex;
  flex-direction: column;
  height: 100%;
  width: calc(100% + 2rem);
  padding: 0 1rem 100px;
  margin-left: -1rem;
  overflow-y: auto;
  overflow-x: hidden;
  &__contents {
    overflow-y: auto;
    padding: 0 1.2rem 1rem;
    margin: 0 -1.2rem;
  }
  &__footer {
    position: absolute;
    bottom: 0;
    width: calc(100% - 3rem);
    flex-shrink: 0;
    display: flex;
    padding: $spacing-4 0;
  }
}
</style>
