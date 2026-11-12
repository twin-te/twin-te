<script setup lang="ts">
withDefaults(
  defineProps<{
    placeholder?: string;
    type?: "normal" | "slim";
    added?: boolean;
    disabled?: boolean;
  }>(),
  { placeholder: "", type: "normal", disabled: false, added: false }
);
const model = defineModel<string>({ required: true });

defineEmits<{
  "enter-text-field": [];
  close: [];
  change: [Event];
}>();

const handleInput = (e: Event) => {
  if (!(e.target instanceof HTMLInputElement)) return;
  model.value = e.target.value;
};
</script>

<template>
  <div :class="{ 'text-field': true, [`text-field--${type}`]: true }">
    <div class="text-field__box">
      <input
        :class="{
          'text-field': true,
          'text-field__input': true,
          '--disabled': disabled,
        }"
        type="text"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        @input="handleInput"
        @change="(e) => $emit('change', e)"
        @keydown.enter="$emit('enter-text-field')"
      />
    </div>
    <div
      v-if="added"
      class="text-field__icon material-icons"
      @click="$emit('close')"
    >
      close
    </div>
  </div>
</template>

<style scoped lang="scss">
@use "~/ui/styles" as *;

.text-field {
  display: flex;
  align-items: center;
  gap: $spacing-2;

  &__box {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    padding: 0rem $spacing-4;
    box-shadow: $shadow-input-concave;
    border-radius: $radius-input;
    background: getColor(--color-base);
  }

  &--normal {
    height: 4rem;
  }

  &--slim {
    height: 3.4rem;
  }

  &__input {
    width: 100%;
    height: 2rem;
    font-size: 1.6rem; //スマホでのinput入力時拡大防止
    transform: scale(0.875); //$text-mediumにする
    margin: 0 -6%; //scaleで縮んだ表示領域の調整
    line-height: $fit;
    color: getColor(--color-text-main);
    background: getColor(--color-base);

    &:focus {
      outline: none;
    }

    &::placeholder {
      color: getColor(--color-unselected);
    }
  }

  &__icon {
    color: getColor(--color-button-gray);
    font-size: 2rem;
    @include button-cursor;
  }

  &.--disabled {
    opacity: 0.3;
    box-shadow: $shadow-convex;
  }
}
</style>
