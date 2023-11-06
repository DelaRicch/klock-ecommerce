<script setup lang="ts">
import EyeViewComponent from "./EyeViewComponent.vue";
import {ref, watch} from "vue";

const { label, placeholder, type, id, requiredTag, isPassword } = defineProps({
  label: {
    type: String,
required: true,
  },
  placeholder: {
    type: String,
    required: true,
  },
  type: {
    type: String,
    default: 'text',
  },
  id: {
    type: String,
    required: true,
  },
  requiredTag: {
    type: Boolean,
    default: false,
  },
  isPassword: {
    type: Boolean,
    default: false,
  },

  error: {
    type: String,
    default: '',
  },
})

defineEmits(['updateValue'])

const isHidden = ref(false);
const inputType = ref(type);

const handleToggleIsHidden = () => isHidden.value = !isHidden.value;

watch(isHidden, (value) => {
  inputType.value = value ? "text" : type;
});


</script>

<template>
  <div class="flex flex-col gap-1 w-full">
    <div class="flex gap-1 items-center">

    <label :for="id" class="text-[#1D2939] font-medium">
      {{ label }}
    </label>
      <span v-if="requiredTag" class="text-red-600 font-semibold">*</span>
    </div>
    <div class="relative h-[2.9rem]">
      <input
          @input.prevent="$emit('updateValue', ($event.target as HTMLInputElement).value);"
             :type="inputType"
          :id="id"
          :placeholder="placeholder"
             class="placeholder:text-[#667085] rounded-lg h-full border border-[#667085] w-full px-3 outline-none outline-offset-4 hover:outline-1 hover:outline-[#0408E7] focus:outline-1 focus:outline-[#0408E7] focus:ring-2 focus:ring-[#4B4EFC] hover:border-[#0408E7] pr-8" />

      <button
          type="button"
          v-if="isPassword"
               @click="handleToggleIsHidden"
              class="absolute right-2 top-[50%] translate-y-[-50%] rounded-lg outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]">
        <EyeViewComponent :is-hidden="isHidden" />
      </button>
    </div>
    <span v-if="!!error"
          class="text-sm text-red-600">{{ error }}</span>
  </div>


</template>