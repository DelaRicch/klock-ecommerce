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
    :class="
      isDesktop
        ? 'text-base font-medium text-black'
        : 'text-2xl font-black text-white'
    "
  >
    {{ title }}
  </h1>
</template>
