<template>
<!--  Overlay for hiding dropdown menu when clicked outside-->
  <div v-if="buttonValue" class="w-screen h-screen absolute left-0 top-0"
       @click="handleToggleDropdown(false)"></div>

  <div :class="className"
       class="w-[12rem] h-max flex flex-col gap-2 relative">
    <ButtonComponent
        :label="dropdownOption"
        :model-value="buttonValue"
        background-color="#FFFFFF"
        :color="isDefaultValue ? '#667085' : '#1D2939'"
        @update:model-value="handleToggleDropdown"
        border
    >
     <slot></slot>
    </ButtonComponent>
<Transition>
    <div v-if="buttonValue"
         class="absolute top-14 left-0 border w-full flex flex-col gap-1.5 rounded-lg overflow-clip max-h-[10rem] overflow-y-auto z-[100] bg-white">
      <button
          v-for="(option, index) in options"
          :key="index"
          class="py-2 bg-slate-50 text-[#1D2939] font-medium hover:bg-slate-200 transition-all duration-200 ease-linear"
          @click="handleSelectOption(option)"
      >{{ option }}
      </button>
    </div>
</Transition>
  </div>
</template>
<script lang="ts" setup>

import {buttonValue} from "../store/resuableState.ts";
import ButtonComponent from "../components/ButtonComponent.vue";
import {PropType, ref} from "vue";

const emit = defineEmits(["selectedOption"]);
const props = defineProps({
  options: {
    type: Array as PropType<string[]>,
    required: true,
  },
  className: {
    type: String,
    required: false,
  }
})

const dropdownOption = ref<string>("");
const isDefaultValue = ref(true);

const handleToggleDropdown = (value: boolean) => {
  buttonValue.value = value;
}

const handleSelectOption = (option: string) => {
  buttonValue.value = false;
  dropdownOption.value = option;
  isDefaultValue.value = false;
  emit("selectedOption", option);
}

</script>

<style scoped>
.v-enter-active,
.v-leave-active {
  @apply transition-all duration-200 ease-linear;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
  transform: translateY(-40%);
}

* {
  scrollbar-width: none;
}

::-webkit-scrollbar {
  width: 0; /* width of the scrollbar */
}
</style>