<template>
  <div class="relative h-2 w-[6rem] rounded-2xl bg-slate-200">
    <div
      :style="{ width: `${(itemsLeft / totalQuantity) * 100}%` }"
      class="absolute left-0 top-0 h-full rounded-2xl bg-[#05CD99]"
    ></div>
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
});
</script>
