<template>
  <div class="flex w-full flex-col gap-4 rounded-2xl bg-[#FAFAFA] p-4 2xl:py-6">
    <div class="flex items-center gap-4">
      <div class="h-[5rem] w-[5rem] border border-slate-200">
        <img :src="productImage" :alt="productName" class="h-full w-full" />
      </div>
      <div class="flex flex-col gap-4">
        <div class="flex flex-col">
          <span class="font-semibold">{{ productName }}</span>
          <span class="text-[0.875rem] text-[#667085]"></span>
        </div>
        <span class="font-semibold">${{ productPrice }}</span>
      </div>
    </div>
    <div class="flex flex-col">
      <span class="font-semibold">Summary</span>
      <p class="line-clamp-6 leading-8 text-[#667085] sm:line-clamp-4">
        {{ productDescription }}
      </p>
    </div>
    <div class="flex w-full flex-col rounded-lg border border-slate-400 p-4">
      <div
        class="flex items-center justify-between text-[0.875rem] text-[#232321]"
      >
        <span>Sales</span>
        <div class="flex items-center gap-2">
          <TopDownArrow />
          <span>{{ productsSold }}</span>
        </div>
      </div>
      <div class="my-2 h-[1px] w-full bg-slate-400 2xl:my-4"></div>
      <div
        class="flex flex-wrap items-center justify-between text-[0.875rem] text-[#232321]"
      >
        <span>Remaining Products</span>
        <div class="mt-2 flex flex-wrap items-center gap-2 lg:mt-0">
          <SlideIndicator
            :items-left="productsRemaining"
            :total-quantity="productQuantity"
          />
          <span>{{ productsRemaining }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import SlideIndicator from "../../components/SlideIndicator.vue";
import TopDownArrow from "../../assets/TopDownArrow.vue";

defineProps({
  productName: {
    type: String,
    required: true,
  },
  productImage: {
    type: String,
    required: true,
  },
  productDescription: {
    type: String,
    required: true,
  },
  productPrice: {
    type: Number,
    required: true,
  },
  productQuantity: {
    type: Number,
    required: true,
  },
  productsSold: {
    type: Number,
    required: true,
  },
  productsRemaining: {
    type: Number,
    required: true,
  },
});
</script>
