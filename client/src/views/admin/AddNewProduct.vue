<template>
  <section
      class="main-dashboard py-[0.8rem] flex flex-col gap-5 px-3 sm:px-[2rem]">

  <BreadCrumbs :breadcrumbs="breadcrumbs" />
    <div class="flex flex-col gap-6">
      <h3 class="text-[#1D2939] text-2xl font-semibold">Add New Product</h3>

      <form @submit.prevent="handleSubmit">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-y-7 md:gap-8 lg:gap-12 xl:gap-[4rem]">
          <div class="flex flex-col gap-5">
            <InputField
                id="product-name"
                placeholder="Enter name of product "
                label="Product Name"
                @update-value="handleSelectOption($event, 'productName')"
            />
            <InputField
                id="product-description"
                        placeholder="Enter product Description"
                        label="Product Description"
                text-type="textarea"
                @update-value="handleSelectOption($event, 'productDescription')"
            />

            <div class="flex flex-col gap-1">
              <span class="text-[#1D2939] font-medium">Category</span>
                      <DropdownComponent class-name="w-full"
                                         :options="dropdownOptions"
                                         @selected-option="handleSelectOption($event,'productCategory')">
                        <span v-if="!productData.productCategory"
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
                  type="number"
                  placeholder="Enter regular price"
                  label="Regular Price"
                  @update-value="handleSelectOption($event, 'productPrice')"
              />
              <InputField
                  id="product-discount-percentage"
                  type="number"
                  placeholder="Enter discount percentage"
                  label="Discount Percentage"
                  @update-value="handleSelectOption($event, 'productDiscountPercentage')"
              />
              <InputField
                  id="product-quantity"
                  placeholder="Enter items in stock"
                  label="Stock quantity"
                  type="number"
                  @update-value="handleSelectOption($event, 'productQuantity')"
              />
              <InputField
                  id="product-brand-name"
                  placeholder="Enter brand name"
                  label="Brand Name"
                  @update-value="handleSelectOption($event, 'productBrandName')"
              />
            </div>

          </div>
          <div class="flex flex-col gap-7">
            <InputField
                id="product-cover-image"
                label="Cover Image"
                type="file"
                @update-value="handleSelectOption($event, 'productCoverImage')"
           >
              <div
                  class="w-full h-4/6 flex flex-col gap-2 items-center justify-center cursor-pointer">
                <img class="w-[46px]" :src="FileUpload" alt="file-upload">
                <div class="flex flex-col text-sm text-[#667085]">
                   <span class="line-clamp-1"><span class="text-blue-600">Click to upload </span>or drag and
                     drop</span>
                  <span class="text-[0.7rem] line-clamp-1">SVG, PNG, JPG or GIF (max. 800x400px)</span>
                </div>
              </div>
            </InputField>

            <div class="flex flex-col gap-1">
              <span class="text-[#1D2939] font-medium">Product Gallery</span>
              <span class="text-[#667085] text-[0.875rem]">Enhance your product showcase. Add an
                Image to
                your product
                gallery.</span>
              <div class="flex items-center flex-wrap gap-6">
              <InputField
                  v-for="(image, index) in
              productData.productGalleryImages"
                  :id="image.id"
                  type="file"
                  file-class-name
                  @update-value="handleSelectOption({value: $event, id: image.id},
                  'productGalleryImages')"
              >
                <div class="flex flex-col justify-center items-center my-4 mx-3 cursor-pointer">
                  <span>
                  <img :src="AddImage" alt="add-image">
                  </span>
                  <span class="text-[#1D2939] text-[0.75rem]">Add image {{index +1}}</span>
                </div>
              </InputField>
              </div>
            </div>
              <ButtonComponent :is-disabled="!isValidForm" label="Submit" type="submit"
                               color="#FFFFFF"
                               background-color="#1D2939" />
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
import {computed, reactive } from "vue";
import FileUpload from "@/assets/file-upload.svg";
import AddImage from '@/assets/add-image.svg';
import ButtonComponent from "@/components/ButtonComponent.vue";
import {addNewProduct} from "@/api/products.ts";
import {errorApiRequest, successApiRequest} from "@/lib/helperFunctions.ts";

const breadcrumbs = [
  {label: "Home", path: '/dashboard'},
  {label: "Product", path: '/dashboard/all-products'},
  {label: "New Product", path: ''},
]
const dropdownOptions = ['Men', 'Women']

const productData = reactive({
  productName: "",
  productDescription: "",
  productCategory: "",
  productPrice: 0,
  productDiscountPercentage: 0,
  productQuantity: 0,
  productBrandName: "",
  productCoverImage: "",
  productGalleryImages: [
    { id: "gallery-image-1", image: null },
    { id: "gallery-image-2", image: null },
    { id: "gallery-image-3", image: null },
    { id: "gallery-image-4", image: null },
  ]
})

const handleSelectOption = (option: string, field: string) => {
if (field === 'productGalleryImages') {
  productData[field].forEach((image: File) => {
    if (image.id === option.id) {
      image.image = option.value
    }
  })
} else {
  productData[field] = option
}
};

const isValidForm = computed(() => {
  return productData.productName !== "" &&
      productData.productDescription !== "" &&
      productData.productCategory !== "" &&
      productData.productPrice !== 0 &&
      productData.productDiscountPercentage !== 0 &&
      productData.productQuantity !== 0 &&
      productData.productBrandName !== "" &&
      productData.productCoverImage !== "" &&
      productData.productGalleryImages.every((image: File) => image.image !== null)
})

const handleSubmit = () => {
 const formData = new FormData()
  for (const field in productData) {
    if (field === 'productGalleryImages') {
      productData[field].forEach((image: File) => {
        formData.append(field, image.image)
      })
    } else {
      formData.append(field, productData[field])
    }
  }

  if (isValidForm.value) {
    console.log(formData.get('productCoverImage'))
  addNewProduct(formData)
      .then((res) => { successApiRequest(res)})
      .catch((err) => {
        // errorApiRequest(err)
        console.log(err)
      });
  }
}

</script>