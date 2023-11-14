<template>
  <div class="w-[3rem] h-[3rem] md:w-[3.5rem] md:h-[3.5rem] rounded-full bg-blue-600 overflow-clip">
    <img class=" bg-no-repeat bg-center bg-cover w-full h-full"
        :src="src" :alt="alt" />
  </div>
</template>

<script setup lang="ts">
  defineProps({
    src: {
      type: String,
      required: true,
    },
    alt: {
      type: String,
      required: true,
    },
  })
</script>