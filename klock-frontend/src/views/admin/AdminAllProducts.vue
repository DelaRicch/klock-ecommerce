<template>
  <section
    class="main-dashboard flex flex-col gap-12 px-3 py-[2.25rem] sm:px-[2rem]"
  >
    <div class="flex w-full flex-wrap items-center justify-between gap-3">
      <h3 class="text-2xl font-semibold text-[#1D2939]">All Products</h3>
      <div>
        <ButtonComponent
          @update:model-value="handleAddNewProduct"
          :label="isSmallerScreen ? '' : 'Add New Product'"
          color="white"
          background-color="#1D2939"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="32"
            height="33"
            viewBox="0 0 32 33"
            fill="none"
          >
            <path
              d="M16 2.5C12.3008 2.54463 8.76575 4.03395 6.14985 6.64985C3.53395 9.26575 2.04463 12.8008 2 16.5C2.04463 20.1992 3.53395 23.7343 6.14985 26.3501C8.76575 28.966 12.3008 30.4554 16 30.5C19.6992 30.4554 23.2343 28.966 25.8501 26.3501C28.466 23.7343 29.9554 20.1992 30 16.5C29.9554 12.8008 28.466 9.26575 25.8501 6.64985C23.2343 4.03395 19.6992 2.54463 16 2.5ZM24 17.5H17V24.5H15V17.5H8V15.5H15V8.5H17V15.5H24V17.5Z"
              fill="#F9FAFB"
            />
          </svg>
        </ButtonComponent>
      </div>
    </div>
    <div
      class="grid grid-cols-1 gap-6 sm:grid-cols-2 md:gap-5 lg:grid-cols-3 xl:gap-6"
    >
      <SingleProduct
        v-for="product in displayedItems"
        :id="product.ProductID"
        :product-name="product.ProductName"
        :product-image="product.ProductCoverImage"
        :product-description="product.ProductDescription"
        :product-price="product.ProductPrice"
        :product-quantity="product.ProductQuantity"
        :products-sold="product.ProductsSold"
        :products-remaining="product.ProductsRemaining"
      />
    </div>
    <div
      v-if="isLoading"
      class="my-[20vh] flex w-full items-center justify-center"
    >
      <loading-component loading-text="fetching products" />
    </div>
    <div
    v-if="!displayedItems.length && !isLoading"
    class="flex h-max w-full justify-center mt-[12vh]"
  >
    <EmptyStateComponent />
  </div>
    <div v-if="totalPages > 1" class="my-1 md:my-2">
      <pagination-component
        :total-pages="totalPages"
        @update:current-page="handleDisplayCurrentPage"
      />
    </div>
  </section>
</template>
<script lang="ts" setup>
import ButtonComponent from "../../components/ButtonComponent.vue";
import SingleProduct from "../../components/admin-all-products/SingleProduct.vue";
import { computed, onBeforeMount, onBeforeUnmount, ref, watch } from "vue";
import { useRouter } from "vue-router";
import PaginationComponent from "@/components/PaginationComponent.vue";
import { getAllProducts } from "@/api/products.ts";
import { AllProductsProps } from "@/types";
import { errorApiRequest } from "@/lib/helperFunctions.ts";
import LoadingComponent from "@/components/LoadingComponent.vue";
import EmptyStateComponent from "@/components/EmptyStateComponent.vue";

const router = useRouter();
const isSmallerScreen = ref(false);
const allProducts = ref<AllProductsProps[]>([]);
const currentPage = ref(1);
const itemsPerPage = ref(10);
const isLoading = ref(false);

const totalPages = computed(() => {
  return Math.ceil(allProducts.value.length / itemsPerPage.value);
});

// Compute the items to be displayed on the current page
const displayedItems: any = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage.value;
  const endIndex = startIndex + itemsPerPage.value;
  return allProducts.value.slice(startIndex, endIndex);
});

const handleDisplayCurrentPage = (page: number) => {
  currentPage.value = page;
};

const checkScreenSize = () => {
  isSmallerScreen.value = window.innerWidth <= 430;
};

const handleAddNewProduct = () => {
  router.push("add-new-product");
};

const getAllAddedProducts = () => {
  isLoading.value = true;
  getAllProducts()
  .then((response) => {
      isLoading.value = false;
      allProducts.value = response;
    })
    .catch((error) => {
      isLoading.value = false;
      errorApiRequest(error);
    });
};

watch(isLoading, (data: boolean) => {
  console.log('====================================');
  console.log(data);
  console.log('====================================');
})

onBeforeMount(() => {
  getAllAddedProducts();
  window.addEventListener("resize", checkScreenSize);
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", checkScreenSize);
});
</script>
