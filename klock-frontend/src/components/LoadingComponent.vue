<template>
  <div
    class="relative flex h-[10rem] w-[10rem] items-center justify-center rounded-full p-2"
  >
    <div
      class="absolute left-0 top-0 h-full w-full rounded-full border-[3px] border-[#1D2939] border-x-transparent motion-safe:animate-spin-slow"
    ></div>
    <span
      class="animate-pulse whitespace-normal text-center text-sm font-semibold uppercase text-[#1D2939]"
      >{{ loadingText }}</span
    >
  </div>
</template>

<script lang="ts" setup>
defineProps({
  loadingText: {
    type: String,
    default: "loading",
  },
});
</script>
