<template>
  <div class="flex items-center justify-between">
    <div class="flex items-center gap-[0.54rem]">
      <div
          class="w-[3.6875rem] h-[4.13rem] rounded border border-slate-200 flex items-center justify-center bg-white">
        <img :alt="title" :src="image" />
      </div>
      <div class="flex flex-col gap-2">
        <div class="flex flex-col">
          <span class="text-[0.70506rem] font-bold">{{title}}</span>
          <span class="text-[0.47006rem] text-[#667085]">{{description}}</span>
        </div>
        <span>${{price}}</span>
      </div>
    </div>
    <div class="flex flex-col font-semibold">
      <span class="text-[#1D2939] text-[0.47006rem]">Sales</span>
      <span class="text-[#667085] text-[0.47006rem]">{{ totalSales }}</span>
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
  price:{
    type: Number,
    required: true,
  },
  totalSales: {
    type: Number,
    required: true,
  },
  image:{
    type: String,
    required: true,
  }
})
</script>