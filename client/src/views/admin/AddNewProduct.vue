<template>
  <section
      class="main-dashboard py-[0.8rem] flex flex-col gap-5 px-3 sm:px-[2rem]">

  <BreadCrumbs :breadcrumbs="breadcrumbs" />
    <div class="flex flex-col gap-6">
      <h3 class="text-[#1D2939] text-2xl font-semibold">Add New Product</h3>

      <form @submit.prevent>
        <div class="grid grid-cols-1 md:grid-cols-2 md:gap-8 lg:gap-12 xl:gap-[4rem]">
          <div class="flex flex-col gap-5">
            <InputField
                id="product-name"
                placeholder="Enter name of product "
                label="Product Name"
            />
            <InputField
                id="product-description"
                        placeholder="Enter product Description"
                        label="Product Description"
                text-type="textarea"
            />

            <div class="flex flex-col gap-1">
              <span class="text-[#1D2939] font-medium">Category</span>
                      <DropdownComponent class-name="w-full"
                                         :options="dropdownOptions"
                                         @selected-option="handleSelectOption"
                      >
                        <span v-if="!selectedOption"
                              class="text-[#667085] absolute left-2 font-normal">Select an
                          option</span>
                        <svg :class="{'transform rotate-180 transition-all duration-300 ease-linear':
                        buttonValue,
                        'rotate-0 transition-all duration-300 ease-linear': !buttonValue
                        }" class="absolute right-1"
                             xmlns="http://www.w3.org/2000/svg"
                             width="24"
                             height="24"
                             viewBox="0 0 24 24"
                              fill="none">
                          <path d="M6 9L12 15L18 9" stroke="#101828" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                        </svg>
                  </DropdownComponent>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-x-4 lg:gap-x-8 gap-y-7">
              <InputField
                  id="product-price"
                  placeholder="Enter regular price"
                  label="Regular Price"
              />
              <InputField
                  id="product-discount-percentage"
                  placeholder="Enter discount percentage"
                  label="Discount Percentage"
              />
              <InputField
                  id="product-quantity"
                  placeholder="Enter items in stock"
                  label="Stock quantity"
                  type="number"
              />
              <InputField
                  id="product-brand-name"
                  placeholder="Enter brand name"
                  label="Brand Name"
              />
            </div>

          </div>
          <div class="lex flex-col gap-4">
            <InputField
                id="product-cover-image"
                label="Cover Image"
                type="file"
            />
          </div>
        </div>
      </form>
      </div>
  </section>
</template>
<script setup lang="ts">
import BreadCrumbs from "@/components/BreadCrumbs.vue";
import DropdownComponent from "@/components/DropdownComponent.vue";
import {buttonValue} from "@/store/resuableState.ts";
import InputField from "@/components/InputField.vue";
import {ref} from "vue";

const breadcrumbs = [
  {label: "Home", path: '/dashboard'},
  {label: "Product", path: '/dashboard/all-products'},
  {label: "New Product", path: ''},
]

const dropdownOptions = ['Men', 'Women']
const selectedOption = ref("");

const handleSelectOption = (option: string) => {
  console.log(option)
  selectedOption.value = option;
}
</script>