<template>
  <div class="flex min-h-[48rem] w-full flex-col gap-6 overflow-x-auto">
    <div
      class="flex min-w-[50rem] flex-col rounded-lg border border-slate-400 pb-[0.62rem] pt-5"
    >
      <h4 class="ml-3 text-2xl font-bold text-[#1D2939]">{{ title }}</h4>
      <table class="mt-2 h-full w-full px-[1rem]" :aria-describedby="title">
        <thead>
          <tr>
            <th>
              <input
                class="cursor-pointer"
                type="checkbox"
                v-model="selectAll"
                @change="selectAllRows"
              />
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
          <tr
            v-for="item in displayedItems"
            :key="item.id"
            class="duration-250 cursor-default text-[#474D66] transition-all ease-linear hover:bg-slate-200"
          >
            <td>
              <input
                class="cursor-pointer"
                type="checkbox"
                v-model="selectedRows"
                :value="item.id"
              />
            </td>
            <td v-for="header in headers" :key="header.key">
              <slot :name="header.key" :item="item">{{
                item[header.key]
              }}</slot>
            </td>
          </tr>
        </tbody>
      </table>
      <div
        v-if="!displayedItems.length"
        class="flex h-max w-full justify-center"
      >
        <EmptyStateComponent />
      </div>
    </div>
    <div class="my-1 md:my-2">
      <PaginationComponent
        v-if="totalPages > 1"
        :total-pages="totalPages"
        @update:current-page="handleDisplayCurrentPage"
      />
    </div>
  </div>
</template>
<script lang="ts" setup>
import { computed, PropType, ref, watch } from "vue";
import PaginationComponent from "../components/PaginationComponent.vue";
import EmptyStateComponent from "../components/EmptyStateComponent.vue";
import { ProductProps } from "@/types";

interface HeaderProps {
  key: string;
  label: string;
}

const props = defineProps({
  title: {
    type: String,
    required: true,
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
    default: 10,
  },
});

const currentPage = ref(1);
const selectedRows = ref<number[]>([]);
const selectAll = ref(false);

watch(selectedRows, (data) => {
  console.log(data);
});
const getColumnWidths = (header: HeaderProps) => {
  return props.columnWidths[header.key] || "auto";
};

// Compute the total number of pages based on the number of items and items per page
const totalPages = computed(() =>
  Math.ceil(props.items?.length / props.itemsPerPage),
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
  @apply border-y border-slate-400 px-3 py-2 text-left font-semibold text-[#1D2939];
}

td {
  @apply px-3 py-2 text-left text-[0.88rem];
}

td:nth-child(2) {
  @apply font-bold;
}
</style>
