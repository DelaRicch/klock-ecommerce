<template>
  <div class="flex h-max w-full md:flex-col gap-2 rounded-t-md pb-4">
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
        class="h-auto w-full"
      />
    </div>
      <div class="flex flex-col gap-0.5">
        <span class="text-xl">Nixon Leather</span>
        <div class="flex gap-4 flex-col">
          <div class="flex items-center gap-1">
            <star-rating-component v-model="rating" :number-of-stars="5" :star-size="18" readonly />
            <span class="text-sm text-[#667085]">(200)</span>
          </div>
          <span class="text-sm font-medium">$65</span>
          <div class="md:mx-auto max-w-[10rem] mt-2">
            <ButtonComponent
              label="Add to cart"
              allow-hover
              background-color="#000000"
              color="#FFFFFF"
            />
        </div>
      </div>
    </div>
  
  </div>
</template>
<script setup lang="ts">
import HeartIcon from "@/assets/HeartIcon.vue";
import ButtonComponent from "./ButtonComponent.vue";
import StarRatingComponent from "./StarRatingComponent.vue";
import { ref, watch } from "vue";


const isLiked = ref(false);
const rating = ref(3.5);

watch(rating, (val) => {
  console.log(val);
});

const handleToggleLike = () => {
  isLiked.value = !isLiked.value;
};
</script>
