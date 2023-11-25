<template>
  <div class="p-4 2xl:py-6 flex flex-col gap-4 w-full rounded-2xl bg-[#FAFAFA]">
      <div class="flex items-center gap-4">
        <div class="w-[5rem] h-[5rem] border border-slate-200">
          <img src="" alt="" class="w-full h-full">
        </div>
        <div class="flex flex-col gap-4">
          <div class="flex flex-col">
            <span class="font-semibold">Breitling leather</span>
            <span class="text-[#667085] text-[0.875rem]">DV One automatic 43mm</span>
          </div>
          <span class="font-semibold">$50</span>
        </div>
      </div>
      <div class="flex flex-col">
        <span class="font-semibold">Summary</span>
        <p class="text-[#667085] leading-8">Breitling Leather for displaying House Motifs</p>
      </div>
      <div class="border border-slate-400 rounded-lg flex flex-col p-4 w-full">
        <div class="flex items-center justify-between text-[#232321] text-[0.875rem]">
          <span>Sales</span>
          <div class="flex items-center gap-2">
            <TopDownArrow />
          <span>1269</span>
          </div>
        </div>
        <div class="h-[1px] w-full my-2 2xl:my-4 bg-slate-400"></div>
        <div class="flex items-center justify-between flex-wrap text-[#232321] text-[0.875rem]">
          <span>Remaining Products</span>
          <div class="flex items-center gap-2 flex-wrap mt-2 lg:mt-0">
           <SlideIndicator />
          <span>1269</span>
          </div>
        </div>
      </div>
  </div>
</template>
<script lang="ts" setup>
import SlideIndicator from "../../components/SlideIndicator.vue";
import TopDownArrow from "../../assets/TopDownArrow.vue";
</script>