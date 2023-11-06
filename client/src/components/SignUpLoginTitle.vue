<script setup lang="ts">
  defineProps({
    title: {
      type: String,
      required: true,
    },
    isDesktop: {
      type: Boolean,
      default: false,
    },
  });
</script>

<template>
  <h1
      class="uppercase"
      :class="isDesktop ? 'text-black font-medium text-base' :
      'font-black text-white text-2xl'">{{title
    }}</h1>
</template>