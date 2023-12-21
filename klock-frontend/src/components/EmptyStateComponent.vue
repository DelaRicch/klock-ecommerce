<template>
  <div class="flex flex-col text-center">
    <Vue3Lottie
      :animationData="EmptyData"
      :height="200"
      :width="200"
      :loop="true"
      :speed="1"
      :autoPlay="true"
      :pauseAnimation="false"
      :pauseOnHover="false"
    />
    <span class="-mt-6 font-semibold text-slate-500">{{ emptyMessage }}</span>
  </div>
</template>
<script lang="ts" setup>
defineProps({
  emptyMessage: {
    type: String,
    default: "No data available",
  },
});
import { Vue3Lottie } from "vue3-lottie";
import EmptyData from "../assets/empty-state.json";
</script>
