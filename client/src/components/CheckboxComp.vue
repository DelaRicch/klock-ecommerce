<script setup lang="ts">
defineProps({
  label: {
    type: String,
    required: true,
  },
  id: {
    type: String,
    required: true,
  },
  error: {
    type: String,
    default: "",
  },
});

defineEmits(["updateValue"]);
</script>

<template>
  <div class="relative flex items-center gap-1">
    <input
      type="checkbox"
      :id="id"
      @change.prevent="
        $emit('updateValue', ($event.target as HTMLInputElement).checked)
      "
      class="cursor-pointer rounded-md border border-slate-200 outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]"
    />
    <label :for="id" class="cursor-pointer text-[0.8rem] leading-[1.75rem]">{{
      label
    }}</label>
    <span
      v-if="!!error"
      class="absolute -bottom-4 w-[10rem] text-sm text-red-600"
      >{{ error }}</span
    >
  </div>
</template>
