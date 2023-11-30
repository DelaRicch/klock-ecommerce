<template>
  <div class="p-4 2xl:py-6 flex flex-col gap-4 w-full rounded-2xl bg-[#FAFAFA]">
      <div class="flex items-center gap-4">
        <div class="w-[5rem] h-[5rem] border border-slate-200">
          <img :src="productImage" :alt="productName" class="w-full h-full">
        </div>
        <div class="flex flex-col gap-4">
          <div class="flex flex-col">
            <span class="font-semibold">{{ productName }}</span>
            <span class="text-[#667085] text-[0.875rem]"></span>
          </div>
          <span class="font-semibold">${{ productPrice }}</span>
        </div>
      </div>
      <div class="flex flex-col">
        <span class="font-semibold">Summary</span>
        <p class="text-[#667085] leading-8 line-clamp-6 sm:line-clamp-4">{{
            productDescription
          }}</p>
      </div>
      <div class="border border-slate-400 rounded-lg flex flex-col p-4 w-full">
        <div class="flex items-center justify-between text-[#232321] text-[0.875rem]">
          <span>Sales</span>
          <div class="flex items-center gap-2">
            <TopDownArrow />
          <span>{{ productsSold }}</span>
          </div>
        </div>
        <div class="h-[1px] w-full my-2 2xl:my-4 bg-slate-400"></div>
        <div class="flex items-center justify-between flex-wrap text-[#232321] text-[0.875rem]">
          <span>Remaining Products</span>
          <div class="flex items-center gap-2 flex-wrap mt-2 lg:mt-0">
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
  required: true
 },
  productImage: {
    type: String,
    required: true
  },
  productDescription: {
    type: String,
    required: true
  },
  productPrice: {
    type: Number,
    required: true
  },
  productQuantity: {
    type: Number,
    required: true
  },
  productsSold: {
    type: Number,
    required: true
  },
  productsRemaining: {
    type: Number,
    required: true
  }
})
</script>