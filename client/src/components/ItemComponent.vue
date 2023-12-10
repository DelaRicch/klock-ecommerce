<template>
  <div class="flex h-max w-full flex-col gap-1 rounded-t-md pb-4">
    <div
      class="relative flex h-auto w-full items-center justify-center rounded-md border p-4 transition-all duration-300 ease-linear hover:shadow-md hover:shadow-black/20"
    >
      <div class="absolute right-2 top-1">
        <ButtonComponent @update:model-value="handleToggleLike" is-icon>
          <HeartIcon :is-liked="isLiked" />
        </ButtonComponent>
      </div>
      <img
        src="../assets/watch.png"
        alt="watch"
        class="h-full w-2/3 sm:w-full"
      />
    </div>
    <div class="flex justify-between pr-4">
      <div class="flex flex-col">
        <span class="text-xl">Nixon Leather</span>
        <div class="flex gap-4 items-center">
          <star-rating-component v-model="rating" :number-of-stars="5" :star-size="18" />
          <span class="text-sm text-[#667085]">(200)</span>
        </div>
      </div>
      <span class="text-sm">$65</span>
    </div>
    <div class="mx-auto w-[10rem]">
      <ButtonComponent
        label="Add to cart"
        allow-hover
        background-color="#000000"
        color="#FFFFFF"
      />
    </div>
  </div>
</template>
<script setup lang="ts">
import HeartIcon from "@/assets/HeartIcon.vue";
import ButtonComponent from "./ButtonComponent.vue";
import StarRatingComponent from "./StarRatingComponent.vue";
import { ref, watch } from "vue";


const isLiked = ref(false);
const rating = ref(0);

watch(rating, (val) => {
  console.log(val);
});

const handleToggleLike = () => {
  isLiked.value = !isLiked.value;
};
</script>
