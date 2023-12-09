<template>
  <div
    class="relative h-[3rem] w-[3rem] overflow-clip rounded-full md:h-[3.4rem] md:w-[3.4rem]"
  >
    <img
      class="h-full w-full bg-cover bg-center bg-no-repeat"
      :src="src"
      :alt="alt"
    />
    <div
      class="absolute left-0 top-0 grid h-full w-full place-items-center rounded-full border-2 border-slate-400 bg-slate-300"
    >
      <UserIcon />
    </div>
  </div>
</template>

<script setup lang="ts">
import UserIcon from "../assets/UserIcon.vue";

defineProps({
  src: {
    type: String,
  },
  alt: {
    type: String,
    required: true,
  },
});
</script>
