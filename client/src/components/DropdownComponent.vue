<template>
<!--  Overlay for hiding dropdown menu when clicked outside-->
  <div v-if="buttonValue" class="w-screen h-screen absolute left-0 top-0"
       @click="handleToggleDropdown(false)"></div>

  <div class="w-[12rem] h-max flex flex-col gap-2 relative">
    <ButtonComponent
        :label="dropdownOption"
        :model-value="buttonValue"
        background-color="#F4F7FE"
        color="#1D2939"
        @update:model-value="handleToggleDropdown"
    >
      <svg fill="none" height="18" viewBox="0 0 18 18" width="18" xmlns="http://www.w3.org/2000/svg">
        <g clip-path="url(#clip0_415_5652)">
          <path d="M15 2.25H14.25V1.5C14.25 1.0875 13.9125 0.75 13.5 0.75C13.0875 0.75 12.75 1.0875 12.75 1.5V2.25H5.25V1.5C5.25 1.0875 4.9125 0.75 4.5 0.75C4.0875 0.75 3.75 1.0875 3.75 1.5V2.25H3C2.175 2.25 1.5 2.925 1.5 3.75V15.75C1.5 16.575 2.175 17.25 3 17.25H15C15.825 17.25 16.5 16.575 16.5 15.75V3.75C16.5 2.925 15.825 2.25 15 2.25ZM14.25 15.75H3.75C3.3375 15.75 3 15.4125 3 15V6H15V15C15 15.4125 14.6625 15.75 14.25 15.75Z" fill="#1D2939"/>
        </g>
        <defs>
          <clipPath id="clip0_415_5652">
            <rect fill="white" height="18" width="18"/>
          </clipPath>
        </defs>
      </svg>
    </ButtonComponent>
<Transition>
    <div v-if="buttonValue"
         class="absolute top-14 left-0 border w-full flex flex-col gap-1.5 rounded-lg overflow-clip max-h-[10rem] overflow-y-auto z-[100] bg-white">
      <button
          v-for="(option, index) in options"
          :key="index"
          class="py-2 bg-slate-200 text-[#1D2939] font-medium hover:bg-slate-50 transition-all duration-200 ease-linear"
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
  }
})

const dropdownOption = ref<string>(props.options[0])

const handleToggleDropdown = (value: boolean) => {
  buttonValue.value = value;
}

const handleSelectOption = (option: string) => {
  buttonValue.value = false;
  dropdownOption.value = option;
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