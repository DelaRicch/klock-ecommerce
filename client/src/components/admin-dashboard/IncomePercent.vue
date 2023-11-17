<template>
  <div class="flex items-center">
    <svg fill="none" height="18" viewBox="0 0 21 18" width="21" xmlns="http://www.w3.org/2000/svg">
      <rect :fill="isIncreased ? '#05CD99' : '#CD0505'" height="17.948" transform="translate(0 -0.00390625)"
            width="20.1242"/>
      <path :transform="isIncreased ? '' : 'rotate(180 10 10)'"
            d="M5.86957 10.4657L10.0621 6.72656L14.2547 10.4657H5.86957Z" fill="#F9FAFB"/>
    </svg>
  <span :class="isIncreased ? 'text-[#05CD99]' : 'text-[#CD0505]'" class="font-bold text-[0.75rem]"> {{isIncreased ? '+' : '-'}}{{ percentage }}%</span>
  </div>
</template>
<script lang="ts" setup>
const props = defineProps({
  percentage: {
    type: Number,
    required: true,
  },
  isIncreased: {
    type: Boolean,
    required: true,
  },
})
</script>