<template>
  <Banner>【ベータ版】{{ currentAcademicYear }}年度のシラバスを使用しています。</Banner>
</template>

<script setup lang="ts">
import { currentAcademicYear } from "~/domain/year";
import Banner from "~/ui/components/Banner.vue";
</script>
