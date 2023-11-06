<script setup lang="ts">
  defineProps({
    type: {
      type: String,
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
required: true,
    },
      backgroundColor: {
      type: String,
        required: true
      },
    color: {
      type: String,
      required: true
    }
  });
</script>

<template>
<button :style="{background: backgroundColor, color: color}" :disabled="isDisabled"
        :class="uppercase &&
'uppercase'" :type="type"
        class="w-full font-semibold h-[45px] rounded-lg flex items-center justify-center outline-offset-4 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC] disabled:cursor-not-allowed disabled:opacity-70">{{label
  }}</button>
</template>