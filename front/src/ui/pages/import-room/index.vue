<script setup lang="ts">
import CardAdd from "~/ui/components/CardAdd.vue";
import IconButton from "~/ui/components/IconButton.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import { useSidebar } from "~/ui/store";

const { openSidebar } = useSidebar();
</script>

<template>
  <div class="import-room">
    <PageHeader>
      <template #left-button-icon>
        <div class="header__left-button-icon">
          <IconButton
            size="large"
            color="normal"
            iconName="menu"
            @click="openSidebar"
          ></IconButton>
        </div>
      </template>
      <template #title>授業場所の登録</template>
    </PageHeader>
    <main class="main">
      <CardAdd
        iconName="view_list"
        heading="Excel ファイルをアップロード"
        text="配布されている Excel ファイルから授業場所情報を登録します。"
        @click-next-button="$router.push('/import-room/excel')"
      />
    </main>
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles/variable";
@use "~/ui/styles/mixin";

@include mixin.header-left-button-delete;

.import-room {
  @include mixin.max-width;
}

.main {
  display: flex;
  flex-direction: column;
  gap: variable.$spacing-4;
  margin-top: variable.$spacing-8;
}
</style>
