<template>
  <div class="w-[3rem] h-[3rem] md:w-[3.4rem] md:h-[3.4rem] rounded-full relative overflow-clip">
    <img class=" bg-no-repeat bg-center bg-cover w-full h-full "
        :src="src" :alt="alt" />
    <div
        class="absolute top-0 left-0 w-full h-full bg-slate-300 border-2 border-slate-400 rounded-full grid place-items-center">
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
  })
</script>