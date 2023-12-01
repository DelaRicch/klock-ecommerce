<template>
  <div class=" w-[10rem] h-[10rem] p-2 rounded-full flex items-center justify-center relative">
    <div
        class="border-[3px] border-[#1D2939] border-x-transparent w-full h-full rounded-full motion-safe:animate-spin-slow absolute top-0 left-0">
    </div>
    <span
        class="whitespace-normal uppercase text-center text-sm font-semibold animate-pulse text-[#1D2939]">{{ loadingText
      }}</span>
  </div>
</template>

<script lang="ts" setup>
defineProps({
  loadingText: {
    type: String,
    default: 'loading'
  }
})
</script>
