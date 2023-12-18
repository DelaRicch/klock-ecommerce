<template>
  <section class="relative flex h-max items-center justify-center">
    <img
      v-if="isMobile"
      src="../assets/hero-two-mobile.png"
      alt="hero-two"
      class="h-auto w-full"
    />

    <img
      v-else
      src="../assets/hero-two.png"
      alt="hero-two"
      class="h-auto w-full"
    />

    <div
      class="absolute left-0 top-[50%] z-20 flex w-full -translate-y-[50%] flex-col gap-2 text-center text-white sm:gap-4 lg:gap-6 xl:gap-8"
    >
      <h3 class="text-xl font-bold capitalize sm:text-2xl lg:text-4xl">
        men's watches
      </h3>
      <p class="text-base font-medium sm:text-xl lg:text-3xl">
        Quality men’s watch at your disposal
      </p>
      <div class="mx-auto w-40">
        <ButtonComponent
          label="Shop Now"
          color="#FFFFFF"
          background-color="#181CF9"
        />
      </div>
    </div>
  </section>
</template>
<script lang="ts" setup>
import { onMounted, ref, onBeforeUnmount } from "vue";
import ButtonComponent from "./ButtonComponent.vue";
const isMobile = ref(true);

const updateBackgroundImage = () => {
  if (window.innerWidth > 425) {
    isMobile.value = false;
  } else {
    isMobile.value = true;
  }
};

onMounted(() => {
  updateBackgroundImage();
  window.addEventListener("resize", updateBackgroundImage);
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", updateBackgroundImage);
});
</script>
