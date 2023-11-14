<template>
  <header class="mobile h-[5rem] bg-red-500 flex items-center justify-between">
    <router-link to="/" class="cursor-pointer">
  <LabelLogo text-color="#1D2939" inner-color="#1D2939" outer-color="#FFFFFF" />
    </router-link>
  </header>
</template>
<script setup lang="ts">
import LabelLogo from "../assets/LabelLogo.vue";
</script>