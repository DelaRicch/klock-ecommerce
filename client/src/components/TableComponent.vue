<template>
  <div class="flex flex-col gap-6 w-full min-h-[37rem] overflow-x-auto">
  <div
      class="min-w-[50rem] rounded-lg border border-[#667085] pt-5 pb-[0.62rem] flex flex-col">
      <h4 class="text-[#1D2939] text-2xl font-bold ml-3">{{title}}</h4>
    <table class="mt-2 w-full h-full px-[1rem]" :aria-describedby="title">
      <thead>
        <tr>
          <th>
            <input class="cursor-pointer" type="checkbox" v-model="selectAll" @change="selectAllRows">
          </th>
          <th
              v-for="header in headers"
              :key="header.key"
              :style="{ width: getColumnWidths(header) }"
          >
            {{ header.label }}
          </th>
        </tr>
      </thead>
      <tbody v-if="displayedItems.length">
      <tr v-for="item in displayedItems" :key="item .id"
          class="text-[#474D66] hover:bg-slate-200 transition-all duration-250 ease-linear cursor-default">
        <td>
          <input class="cursor-pointer" type="checkbox" v-model="selectedRows" :value="item.id">
        </td>
        <td v-for="header in headers" :key="header.key">
          {{ item[header.key] }}
        </td>
      </tr>
      </tbody>
    </table>
    <div v-if="!displayedItems.length" class="h-max flex justify-center w-full">
      <EmptyStateComponent />
    </div>
  </div>
    <PaginationComponent v-if="totalPages > 1" :total-pages="totalPages"
                         @update:current-page="handleDisplayCurrentPage" />
  </div>
</template>
<script lang="ts" setup>
import {computed, PropType, ref, watch} from "vue";
import PaginationComponent from "../components/PaginationComponent.vue";
import EmptyStateComponent from "../components/EmptyStateComponent.vue";
import {ProductProps} from "@/types";

interface HeaderProps {
  key: string;
  label : string;
}

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  headers: {
    type: Array as PropType<HeaderProps[]>,
    required: true,
  },
  items: {
    type: Array as PropType<ProductProps[]>,
    required: true,
  },
  columnWidths: {
    type: Object as () => Record<string, string>,
    default: () => ({}),
  },
  itemsPerPage: {
    type: Number,
    default: 10
  },
})

const currentPage = ref(1);
const selectedRows = ref<number[]>([]);
const selectAll = ref(false);

watch(selectedRows, (data) => {
  console.log(data)
})
const getColumnWidths = (header: HeaderProps) => {
  return props.columnWidths[header.key] || "auto";
};

// Compute the total number of pages based on the number of items and items per page
const totalPages = computed(() =>
    Math.ceil(props.items?.length / props.itemsPerPage)
);

// Compute the items to be displayed on the current page
const displayedItems: any = computed(() => {
  const startIndex = (currentPage.value - 1) * props.itemsPerPage;
  const endIndex = startIndex + props.itemsPerPage;
  return props.items?.slice(startIndex, endIndex);
});

// Display current page based on pagination button clicked
const handleDisplayCurrentPage = (page: number) => {
  currentPage.value = page;
};

// Select all rows
const selectAllRows = () => {
  if (selectAll.value) {
    selectedRows.value = props.items.map((item: ProductProps) => item?.id);
  } else {
    selectedRows.value = [];
  }
};

</script>

<style scoped>
th {
  @apply border-y border-[#98A2B3] px-3 py-2 font-semibold text-[#1D2939] text-left
}

td {
  @apply px-3 py-2 text-left text-[0.88rem]
}

td:nth-child(2) {
  @apply font-bold
}
</style>