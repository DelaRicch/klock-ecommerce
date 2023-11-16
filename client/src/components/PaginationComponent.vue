<template>
  <div
      class="rounded-xl h-[3rem] w-max mx-auto my-2 border-2 border-[#D0D5DD] overflow-clip flex items-center">

    <button
        :disabled="currentPage === 1"
        class="rounded-l-xl outline-[#0408E7] w-[3rem] sm:w-[7.2rem] h-full flex items-center justify-center gap-1 border-r border-[#D0D5DD] disabled:cursor-not-allowed disabled:opacity-50"
        @click="previousPage"
    >
      <svg fill="none" height="20" viewBox="0 0 20 20" width="20"
           xmlns="http://www.w3.org/2000/svg">
        <path d="M15.8333 10.0013H4.16666M4.16666 10.0013L9.99999 15.8346M4.16666 10.0013L9.99999 4.16797" stroke="#344054" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.67"/>
      </svg>
      <span v-if="displayButtonLabels" class="text-[#1D2939] text-[0.875rem]">Previous</span>
    </button>

    <button
        v-for="page in displayedPages"
        :key="page"
        :class="{ 'bg-slate-100': currentPage === page }"
        :disabled="currentPage === page"
        class="outline-[#0408E7] w-[3rem] h-full text-[#344054] flex items-center justify-center p-[0.625rem] border-x border-[#D0D5DD] disabled:cursor-not-allowed"
        @click="goToPage(page)"
    >
      {{page}}
    </button>

    <button
        :disabled="currentPage === totalPages"
        class="rounded-r-xl outline-[#0408E7] w-[3rem] sm:w-[7.2rem] h-full flex items-center justify-center gap-1 border-l border-[#D0D5DD] disabled:cursor-not-allowed disabled:opacity-50"
        @click="nextPage">
      <span v-if="displayButtonLabels" class="text-[#1D2939] text-[0.875rem]">Next</span>
      <svg fill="none" height="20" viewBox="0 0 20 20" width="20"
           xmlns="http://www.w3.org/2000/svg">
        <path d="M4.16669 10.0013H15.8334M15.8334 10.0013L10 4.16797M15.8334 10.0013L10 15.8346" stroke="#344054" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.67"/>
      </svg>
    </button>
  </div>
</template>
<script lang="ts" setup>
import {computed, onBeforeUnmount, onMounted, ref} from "vue";

const emit = defineEmits(['update:currentPage']);
const props = defineProps({
  totalPages: {
    type: Number,
    required: true,
  },
})

const displayButtonLabels = ref(false);
const currentPage = ref(1);
const buttonsToShow = ref(0);
const range = ref(2);
const totalPages = ref(props.totalPages);

const updateDisplayButtonLabels = () => {
  displayButtonLabels.value = window.innerWidth > 640;
}

const updateButtonsToShow = () => {
  if (window.innerWidth >= 375) {
    buttonsToShow.value = 5;
    range.value = 2;
  } else {
    buttonsToShow.value = 3;
    range.value = 1;
  }
};

const displayedPages = computed(() => {
  let start = Math.max(1, currentPage.value - range.value);
  let end = Math.min(totalPages.value, currentPage.value + range.value);

  if (end - start < buttonsToShow.value - 1) {
    const diff = buttonsToShow.value - (end - start + 1);

    if (start > 1) {
      start = Math.max(1, start - diff);
    } else {
      end = Math.min(totalPages.value, end + diff);
    }
  }
  return Array.from({ length: end - start + 1 }, (_, i) => start + i);
});

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page;
    emit('update:currentPage', currentPage.value);
  }
};

const previousPage = () => {
  goToPage(currentPage.value - 1);
};

const nextPage = () => {
  goToPage(currentPage.value + 1);
};

onMounted(() => {
      updateDisplayButtonLabels();
  updateButtonsToShow();
      window.addEventListener("resize", updateDisplayButtonLabels);
      window.addEventListener("resize", updateButtonsToShow);
    }
)

onBeforeUnmount(() => {
  window.removeEventListener("resize", updateDisplayButtonLabels)
    window.removeEventListener("resize", updateButtonsToShow);
})

</script>