<template>
  <button
    :class="{
      uppercase: uppercase,
      'h-6 w-6 sm:h-[2.5rem] sm:w-[2.5rem] hover:border-[#D9D9D9]': isIcon,
      'h-8 md:h-[45px] w-full gap-[0.8rem] px-6 font-semibold': !isIcon,
      'border-[#667085] ': border,
      'border-transparent': !border,
    }"
    :style="{
      color: color,
      backgroundColor: backgroundColor,
    }"
    :disabled="isDisabled || isSubmitting"
    :type="type"
    class="flex items-center justify-center rounded sm:rounded-lg border disabled:cursor-not-allowed disabled:opacity-70 transition-all duration-300 ease-linear"
    @click="handleClick"
    @mouseover="handleMouseOver" 
    @mouseout="handleMouseOut"
    
  >
    <slot></slot>
    <span v-if="isSubmitting">
      <i class="fa fa-spinner fa-spin"></i> Submitting
    </span>
    <span v-if="!isSubmitting" class="text-sm sm:text-base">
      {{ label }}
    </span>
  </button>
</template>

<script lang="ts" setup>
import { PropType, ref } from "vue";

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
  border: {
    type: Boolean,
    default: false,
  },
  isSubmitting: {
    type: Boolean,
    default: false,
  },
  allowHover: {
    type: Boolean,
  }
});

const backgroundColor = ref(props.backgroundColor);
const handleMouseOver = () => {
  if (props.allowHover) { 
    backgroundColor.value = '#181CF9';
  }
    };

    const handleMouseOut = () => {
      backgroundColor.value = props.backgroundColor;
    };


const handleClick = () => {
  const updatedState = !props.modelValue;
  if (props.type === "button") {
    emit("update:modelValue", updatedState);
  }
};
</script>
