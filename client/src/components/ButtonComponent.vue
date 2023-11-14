<template>
  <button
      @click="handleClick"
      :style="{background: backgroundColor, color: color}" :disabled="isDisabled"
          :class="{'uppercase': uppercase,
          'w-[2.5rem] h-[2.5rem] hover:border-[#D9D9D9]': isIcon,
          'w-full px-6 font-semibold h-[45px] gap-[0.8rem]': !isIcon,
          }" :type="type"
          class=" rounded-lg flex items-center justify-center outline-offset-4 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC] disabled:cursor-not-allowed disabled:opacity-70 border border-transparent">
    <slot></slot>
    {{label}}
  </button>
</template>

<script setup lang="ts">
import {PropType} from "vue";

const emit = defineEmits(["update:modelValue"]);

  const props = defineProps({
    type: {
      type: String as PropType<"button" | "submit">,
      default: "button",

    },
  isDisabled: {
    type: Boolean,
    default: false,
  },
  uppercase: {
    type: Boolean,
    default: false,
  },
    label: {
      type: String,
    },
      backgroundColor: {
      type: String,
      },
    color: {
      type: String,
    },
    modelValue: {
      type: Boolean,
    },
    isIcon: {
      type: Boolean,
      default: false,
    },
  });

const handleClick = () => {
    const updatedState = !props.modelValue;
    if (props.type === "button") {
    emit("update:modelValue", updatedState);
    }
  }
</script>