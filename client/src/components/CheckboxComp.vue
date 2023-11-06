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
    default: '',
  },
})

defineEmits(['updateValue'])

</script>

<template>

  <div class="flex items-center gap-1 relative">
    <input type="checkbox" :id="id"
           @change.prevent="$emit('updateValue', ($event.target as HTMLInputElement).checked);"
           class="rounded-md border border-slate-200 outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC] cursor-pointer">
  <label :for="id" class=" text-[0.8rem] leading-[1.75rem] cursor-pointer">{{label}}</label>
    <span v-if="!!error"
          class="text-sm text-red-600 absolute w-[10rem] -bottom-4">{{ error }}</span>
  </div>

</template>