<script setup lang="ts">
import PageHeader from "~/ui/components/PageHeader.vue";
</script>

<template>
  <div class="wrapper">
    <PageHeader>
      <template #title>タグ</template>
    </PageHeader>
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles/mixin" as *;
@use "~/ui/styles/variable" as *;

.wrapper {
  @include max-width;
  height: 100vh;

  display: flex;
  flex-direction: column;
  gap: $spacing-6;

  padding-bottom: $spacing-4;

  @include pc {
    display: grid;
    grid-template: "header header" auto "tags main" 1fr / auto 1fr;
  }
}

.header {
  grid-area: header;
}
</style>
