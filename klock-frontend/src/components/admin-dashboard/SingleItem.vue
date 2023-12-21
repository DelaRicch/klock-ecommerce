<template>
  <div class="flex items-center justify-between">
    <div class="flex items-center gap-[0.54rem]">
      <div
        class="flex h-[4.13rem] w-[3.6875rem] items-center justify-center rounded border border-slate-200 bg-white"
      >
        <img :alt="title" :src="image" />
      </div>
      <div class="flex flex-col gap-2">
        <div class="flex flex-col">
          <span class="font-bold text-[0.7.2rem]">{{ title }}</span>
          <span class="text-[0.7rem] text-[#667085]">{{ description }}</span>
        </div>
        <span>${{ price }}</span>
      </div>
    </div>
    <div class="flex flex-col font-semibold">
      <span class="text-[0.7rem] text-[#1D2939]">Sales</span>
      <span class="text-[0.7rem] text-[#667085]">{{ totalSales }}</span>
    </div>
  </div>
</template>
<script lang="ts" setup>
defineProps({
  title: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: true,
  },
  price: {
    type: Number,
    required: true,
  },
  totalSales: {
    type: Number,
    required: true,
  },
  image: {
    type: String,
    required: true,
  },
});
</script>
