<template>
  <Transition name="fade">
    <div
      v-if="showSnackbar"
      :class="
        success ? 'border-green-500 bg-green-50' : 'border-red-500 bg-red-50'
      "
      class="absolute right-2 top-2 z-[100] flex h-[5rem] min-w-[22rem] cursor-pointer items-center gap-4 rounded-md border py-1.5 pl-3 pr-4"
    >
      <Checkmark :success="success" />
      <div class="flex flex-col">
        <span class="text-lg font-semibold uppercase text-[#1D2939]">{{
          title
        }}</span>
        <p class="mt-1">{{ message }}</p>
      </div>
    </div>
  </Transition>
</template>
<script lang="ts" setup>
import Checkmark from "../components/Checkmark.vue";

defineProps({
  success: {
    type: Boolean,
    required: true,
  },
  title: {
    type: String,
    required: true,
  },
  message: {
    type: String,
    required: true,
  },
  showSnackbar: {
    type: Boolean,
    required: true,
  },
});
</script>

<style lang="css" scoped>
.fade-enter-active,
.fade-leave-active {
  @apply transition-all duration-300 ease-linear;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-100%);
}
</style>
