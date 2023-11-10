<template>
  <Transition name="fade">
    <div
        v-if="showSnackbar"
        :class="success ? 'bg-green-50 border-green-500' : 'bg-red-50 border-red-500'"
        class="cursor-pointer absolute w-[22rem] h-[5rem] rounded-md border top-2 right-2 flex items-center gap-4 px-3 py-1.5 z-[100]">
      <Checkmark :success="success" />
      <div class="flex flex-col">
        <span class="text-[#1D2939] text-lg font-semibold uppercase">{{ title }}</span>
        <p class=" mt-1">{{message}}</p>
      </div>
    </div>

  </Transition>
</template>
<script lang="ts" setup>
import Checkmark from "../components/Checkmark.vue";

defineProps({
  success: {
    type: Boolean,
    required: true
  },
  title: {
    type: String,
    required: true
  },
  message: {
    type: String,
    required: true
  },
  showSnackbar: {
    type: Boolean,
    required: true
  }
})
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