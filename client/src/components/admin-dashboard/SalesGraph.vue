<template>
  <div class="border border-slate-400 rounded-lg p-4 flex flex-col w-full xl:w-[180%]">
    <h3 class="text-[1.25rem] text-[#1D2939] font-semibold">Sales Graph</h3>
    <div
        class="flex flex-col sm:flex-row gap-4 items-center sm:items-start my-4 sm:my-0 justify-between">
      <div class="flex flex-col">
        <span class="text-[#1D2939] text-[2.125rem] font-bold">$37.5K</span>
        <div class="flex items-center gap-2">
          <span class="text-[#667085] text-[0.875rem]">Total</span>
          <IncomePercent :is-increased="false" :percentage="2.44" />
        </div>
        <div class="flex items-center gap-2 mt-2">
          <svg fill="none" height="16" viewBox="0 0 16 16" width="16" xmlns="http://www.w3.org/2000/svg">
            <rect fill="#05CD99" height="16" rx="8" width="16"/>
            <g clip-path="url(#clip0_415_5647)">
              <path d="M6.82068 9.74969L5.21651 8.14552C5.03776 7.96677 4.75359 7.96677 4.57484 8.14552C4.39609 8.32427 4.39609 8.60844 4.57484 8.78719L6.49526 10.7076C6.67401 10.8864 6.96276 10.8864 7.14151 10.7076L11.9998 5.85385C12.1786 5.6751 12.1786 5.39094 11.9998 5.21219C11.8211 5.03344 11.5369 5.03344 11.3582 5.21219L6.82068 9.74969Z" fill="white"/>
            </g>
            <defs>
              <clipPath id="clip0_415_5647">
                <rect fill="white" height="11" transform="translate(2.69568 2.32422)" width="11"/>
              </clipPath>
            </defs>
          </svg>
          <span class="ont-bold text-[#05CD99]">On track</span>
        </div>
      </div>
      <div>
      <date-picker
          v-model="selectedDate"
          separator=" to "
          :formatter="formatter"
          placeholder="Select date"
          class="text-center"
          :shortcuts="false"
      />
      </div>

    </div>
   <div
       class="-mt-3 w-full md:w-11/12 xl:w-9/12 h-max flex self-center overflow-y-clip overflow-x-auto">
     <apexchart :options="chartOptions" :series="series" type="line"
                class="w-full"
                ></apexchart>
   </div>
  </div>
</template>
<script lang="ts" setup>
import IncomePercent from "../admin-dashboard/IncomePercent.vue";
import {ref, watch} from "vue";
import DatePicker from "vue-tailwind-datepicker";

const selectedDate = ref({
  startDate: "",
  endDate: "",
});

watch(selectedDate, (value) => {
  console.log(value);
});

// Custom date picker format
const formatter = ref({
  date: "DD MMM YYYY",
  month: "MMM",
});

const chartOptions = {
  chart: {
    id: 'Sales Graph',
    toolbar: {
      show: false,
    },
  },
  xaxis: {
    categories: ["Xorla", "Dela", "Ricch", "Champ", "Emma" ],
  },
  stroke: {
curve: 'smooth'
  },
  tooltip: {
    enabled: true,
  },
  grid: {
    show: false,
  },
  legend: {
    show: false,
  },
  dataLabels: {
    enabled: false,
  },
  yaxis: {
    show: false,
  },
  fill: {
    colors: ['#4B4EFC'],
  },

}

const series = [{
  name: 'Sales',
  data: [38, 37, 40, 35, 41]
}]

const handleSelectedOption = (option: string) => {
  console.log(option);
}
</script>