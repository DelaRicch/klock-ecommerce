<template>
  <teleport to="body">
    <Transition name="fade">
      <div
        v-if="displayCategories"
        class="top-[5rem] mx-auto flex w-max flex-col bg-white shadow"
      >
        <div class="flex gap-8 border-b border-[#667085] px-[3rem] py-2">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            :id="(tab.id).toString()"
            class="flex cursor-pointer tabs-center gap-2 capitalize"
          >
            <span class="font-medium">{{ tab.label }}</span>
            <svg
              class="rotate-0 transition-all duration-300 ease-linear"
              xmlns="http://www.w3.org/2000/svg"
              width="20"
              height="20"
              viewBox="0 0 20 20"
              fill="none"
            >
              <path
                d="M6 9L12 15L18 9"
                stroke="#101828"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </button>
        </div>
        <div class="mt-5 flex justify-between px-[3rem] py-[1rem]">
          <div class="flex w-[20rem] flex-col gap-3">
            <span>New Arrivals</span>
            <span>Best Sellers</span>
            <span>Watch Bands</span>
            <span>Shop All</span>
          </div>
        </div>
      </div>
    </Transition>
  </teleport>
</template>
<script setup lang="ts">
import { displayCategories } from '@/stores/resuableState';
import { ref } from 'vue';


const tabs = ref([
  { label: "Men's Watches", id: 1 },
  { label: "Women's Watches", id: 2 },
]);
</script>

<style lang="css" scoped>
.fade-enter-active,
.fade-leave-active {
  @apply transition-all duration-300 ease-linear;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-20%);
}
</style>
