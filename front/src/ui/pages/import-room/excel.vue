<script setup lang="ts">
import dayjs from "dayjs";
import { ref } from "vue";
import { getKdbClassroom } from "~/infrastructure/local/courseLocationExcel";
import { LocalStorage } from "~/infrastructure/localstorage";
import Button from "~/ui/components/Button.vue";
import IconButton from "~/ui/components/IconButton.vue";
import InputButtonFile from "~/ui/components/InputButtonFile.vue";
import PageHeader from "~/ui/components/PageHeader.vue";
import { useToast } from "~/ui/store";

const { displayToast } = useToast();

const dayjsFormat = "YYYY/MM/DD HH:mm:ss";

const localStorage = LocalStorage.getInstance();
const latestData = ref(localStorage.get("courseLocationInfo"));

const loadLoading = ref(false);
const uploadLoading = ref(false);

async function load(file: File) {
  loadLoading.value = true;
  try {
    const data = await getKdbClassroom(file);
    latestData.value = data;
    localStorage.set("courseLocationInfo", data);
  } finally {
    loadLoading.value = false;
  }
}

async function upload() {
  uploadLoading.value = true;

  // TODO: アップロード処理
  await new Promise<void>((resolve) => {
    setTimeout(() => resolve(), 2000);
  });

  uploadLoading.value = false;
  displayToast("授業場所の登録が完了しました。", {
    displayPeriod: 5000,
    type: "primary",
  });
  return;
}
</script>

<template>
  <div class="import-excel">
    <PageHeader>
      <template #left-button-icon>
        <IconButton
          size="large"
          color="normal"
          icon-name="arrow_back"
          @click="$router.back()"
        ></IconButton>
      </template>
      <template #title>Excel ファイルをアップロード</template>
    </PageHeader>
    <main class="main">
      <section class="description">
        <p>
          最新の年度において登録されている授業に対して、Excel
          ファイルに掲載されている教室情報を登録します。
        </p>
      </section>
      <section class="upload">
        <p class="upload__header">Excel ファイル</p>
        <InputButtonFile name="excel-file" accept=".xlsx" @change-file="load">
          アップロードする
        </InputButtonFile>
      </section>
      <section class="apply">
        <div v-if="loadLoading" class="loading">読み込み中...</div>
        <div v-else-if="latestData" class="data-info">
          <div class="data-info__content">
            <div class="title">最終アップロード日</div>
            <div class="content">
              {{ dayjs(latestData.uploadAt).format(dayjsFormat) }}
            </div>
          </div>
          <div class="data-info__content">
            <div class="title">授業データ数</div>
            <div class="content">
              {{
                Object.keys(latestData.courseLocations).length.toLocaleString()
              }}
              件
            </div>
          </div>
        </div>
        <Button
          class="btn-upload"
          color="primary"
          size="small"
          layout="fill"
          :state="uploadLoading ? 'disabled' : 'default'"
          @click="upload"
        >
          このデータを登録する</Button
        >
      </section>
      <section class="info">
        <h5 class="header">注意事項</h5>
        <ul class="list">
          <li>登録した授業場所の情報は、このアカウントにのみ保存されます。</li>
          <li>
            新たに授業を追加した際には、再度データの登録操作が必要です。
            <br />
            （その際にファイルをアップロードしなかった場合は、
            最後にアップロードしたファイルの情報が使用されます。）
          </li>
          <li>
            手動または自動で既に授業場所が登録されている授業は変更されないため、
            授業場所が変更された場合には手動で修正してください。
          </li>
          <li>
            授業場所を登録した Twin:te
            のスクリーンショット等は、取り扱いに十分注意していただき、
            大学による注意事項を守って利用してください。
          </li>
        </ul>
      </section>
    </main>
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles/variable";
@use "~/ui/styles/mixin";

.import-excel {
  @include mixin.max-width;
}

.description {
  margin-top: variable.$spacing-8;
}

.upload {
  margin-top: variable.$spacing-10;

  &__header {
    font-weight: 500;
  }
}

.apply {
  margin-top: variable.$spacing-8;

  .data-info {
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    gap: variable.$spacing-4;

    &__content {
      .title {
        font-weight: 500;
        font-size: 1.2rem;
      }
      .content {
        margin-top: variable.$spacing-1;
        font-size: 1.5rem;
      }
    }
  }

  .btn-upload {
    margin-top: variable.$spacing-8;
  }
}

.info {
  margin-top: variable.$spacing-8;

  .header {
    font-size: 1.4rem;
    font-weight: 500;
  }
  .list {
    margin-top: variable.$spacing-2;
    li {
      list-style: disc inside;
      line-height: 130%;
      margin-bottom: variable.$spacing-2;
    }
  }
}
</style>
