<template>
  <!--  Overlay for hiding dropdown menu when clicked outside-->
  <div
    v-if="buttonValue"
    class="absolute left-0 top-0 h-screen w-screen"
    @click="handleToggleDropdown(false)"
  ></div>

  <div :class="className" class="relative flex h-max w-[12rem] flex-col gap-2">
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
      <div
        v-if="buttonValue"
        class="absolute left-0 top-14 z-[100] flex max-h-[10rem] w-full flex-col gap-1.5 overflow-clip overflow-y-auto rounded-lg border bg-white"
      >
        <button
          v-for="(option, index) in options"
          :key="index"
          class="bg-slate-50 py-2 font-medium text-[#1D2939] transition-all duration-200 ease-linear hover:bg-slate-200"
          @click="handleSelectOption(option)"
        >
          {{ option }}
        </button>
      </div>
    </Transition>
  </div>
</template>
<script lang="ts" setup>
import { buttonValue } from "../store/resuableState.ts";
import ButtonComponent from "../components/ButtonComponent.vue";
import { PropType, ref } from "vue";

const emit = defineEmits(["selectedOption"]);
const props = defineProps({
  options: {
    type: Array as PropType<string[]>,
    required: true,
  },
  className: {
    type: String,
    required: false,
  },
});

const dropdownOption = ref<string>("");
const isDefaultValue = ref(true);

const handleToggleDropdown = (value: boolean) => {
  buttonValue.value = value;
};

const handleSelectOption = (option: string) => {
  buttonValue.value = false;
  dropdownOption.value = option;
  isDefaultValue.value = false;
  emit("selectedOption", option);
};
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
