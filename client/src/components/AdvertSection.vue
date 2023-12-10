<template>
  <section class="relative flex h-max items-center justify-center">
    <img
      src="../assets/hero-two.png"
      alt="hero-two"
      class=" h-auto w-full"
    />
    <div
      class="z-20  absolute left-0 top-1/4 w-full flex flex-col gap-1 text-center text-white sm:gap-4 md:gap-6 xl:gap-8"
    >
      <h3 class="text-xl font-bold capitalize sm:text-2xl md:text-4xl">
        men's watches
      </h3>
      <p class="text-base font-medium sm:text-xl md:text-3xl">
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
import ButtonComponent from "./ButtonComponent.vue";
</script>
