<template>
  <div class="relative w-[6rem] h-2 bg-slate-200 rounded-2xl">
    <div :style="{ width: `${(itemsLeft / totalQuantity) * 100}%` }"
         class="absolute top-0 left-0 h-full bg-[#05CD99] rounded-2xl"></div>
  </div>
</template>
<script lang="ts" setup>
defineProps({
  itemsLeft: {
    type: Number,
    required: true,
  },
  totalQuantity: {
    type: Number,
    required: true,
  },
})
</script>