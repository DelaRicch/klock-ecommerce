<template>
    <div class="relative h-max w-full flex flex-col gap-1">
        <div class="absolute right-2">
            <ButtonComponent @update:model-value="handleToggleLike" is-icon>
                <HeartIcon :is-liked="isLiked" />
            </ButtonComponent>
        </div>
        <div
          class="flex h-auto w-full items-center justify-center rounded-md border p-4"
        >
          <img src="../assets/watch.png" alt="watch" class="w-2/3 sm:w-full h-full" />
        </div>
        <div class="flex justify-between pr-4">
          <div class="flex flex-col">
            <span class="text-xl">Nixon Leather</span>
            <div class="flex gap-4">
              * * * *
              <span class="text-sm text-[#667085]">(200)</span>
            </div>
          </div>
          <span class="text-sm">$65</span>
        </div>
        <div class="mx-auto w-[10rem]">
          <ButtonComponent label="Add to cart" allow-hover background-color="#000000" color="#FFFFFF"  />
        </div>
      </div>
</template>
<script setup lang="ts">
import HeartIcon from '@/assets/HeartIcon.vue';
import ButtonComponent from './ButtonComponent.vue';
import { ref } from 'vue';

const isLiked = ref(false);

const handleToggleLike = () => {
    isLiked.value = !isLiked.value;
}

</script>