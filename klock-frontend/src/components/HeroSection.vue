<template>
  <section class="relative h-max w-full">
    <div class="absolute h-full w-full bg-[#050D5833]/20"></div>
    <img src="../assets/hero.png" alt="hero-bg" class="h-full w-full" />
    <div
      class="absolute bottom-0 left-0 flex h-max w-full gap-3 px-[6vw] pb-2 pt-6 sm:gap-10 sm:pb-4 md:gap-14 xl:gap-28 xl:pb-24"
    >
      <div
        class="flex h-10 w-10 items-center justify-center rounded-full bg-white sm:h-14 sm:w-14 md:h-20 md:w-20 xl:h-[5.5rem] xl:w-[5.5rem]"
      >
        <img
          src="../assets/arrow-left.svg"
          alt="left-arrow"
          class="w-2/3 sm:w-[70%]"
        />
      </div>
      <div
        class="flex flex-1 flex-col gap-1 text-center font-bold uppercase text-white sm:gap-4 md:gap-6 lg:gap-10"
      >
        <h3 class="sm:text-3xl xl:text-5xl">black friday</h3>
        <p
        v-if="isDesktop"
          class="line-clamp-1 text-xs sm:line-clamp-3 sm:text-base lg:line-clamp-none"
        >
          Prepare to be amazed as we slash prices across the store! Enjoy
          jaw-dropping discounts of up to 80% on a wide range of products. It's
          your chance to own the items you've been eyeing all year long.
        </p>
        <div class="mx-auto text-sm sm:w-[10rem] sm:text-base">
          <ButtonComponent
            label="Get Started"
            color="#FFFFFF"
            background-color="#181CF9"
          />
        </div>
        <div class="mx-auto flex h-1 w-[10rem] gap-2 overflow-clip rounded-lg">
          <div class="h-full w-full bg-white"></div>
          <div class="h-full w-full bg-slate-100/70"></div>
        </div>
      </div>
      <div
        class="flex h-10 w-10 items-center justify-center rounded-full bg-white sm:h-14 sm:w-14 md:h-20 md:w-20 xl:h-[5.5rem] xl:w-[5.5rem]"
      >
        <img
          src="../assets/arrow-left.svg"
          alt="left-arrow"
          class="w-2/3 rotate-180 transform sm:w-[70%]"
        />
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { onMounted, ref, onBeforeUnmount } from "vue";
import ButtonComponent from "./ButtonComponent.vue";

const isDesktop = ref(false);

const updateDeviceWidth = () => {
  if (window.innerWidth > 425) {
    isDesktop.value = true;
  } else {
    isDesktop.value = false;
  }
};

onMounted(() => {
  updateDeviceWidth();
  window.addEventListener("resize", updateDeviceWidth);
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", updateDeviceWidth);
});
</script>
