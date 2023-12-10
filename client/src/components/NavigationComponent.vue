<template>
  <div class="flex items-center gap-2">
    <ButtonComponent
      @update:model-value="previousPage"
      :is-disabled="currentPage === 1"
      is-icon
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="20"
        height="20"
        viewBox="0 0 20 20"
        fill="none"
      >
        <path
          d="M12.5 5L7.5 10L12.5 15"
          :stroke="isDisableButton ? '#667085' : '#1E1E1E'"
          stroke-width="1.66667"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </ButtonComponent>
    <span class="text-xs text-[#667085]"
      >Results {{ currentPage }} of
      <span class="text-black">{{ totalPages }}</span></span
    >
    <ButtonComponent
      @update:model-value="nextPage"
      :is-disabled="currentPage === totalPages"
      is-icon
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="20"
        height="20"
        viewBox="0 0 20 20"
        fill="none"
      >
        <path
          d="M7.5 15L12.5 10L7.5 5"
          :stroke="isDisableButton ? '#667085' : '#1E1E1E'"
          stroke-width="1.66667"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </ButtonComponent>
  </div>
</template>
<script setup lang="ts">
import { ref, watch } from "vue";
import ButtonComponent from "./ButtonComponent.vue";

const emit = defineEmits(["update:currentPage"]);

defineProps({
  totalPages: {
    type: Number,
    required: true,
  },
});

const currentPage = ref(1);
const isDisableButton = ref(false);

const previousPage = () => {
  currentPage.value -= 1;
};

const nextPage = () => {
      currentPage.value += 1;
};

watch(currentPage, data => {
    emit("update:currentPage", data);
})
</script>
