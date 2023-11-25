<template>
  <section
      class="main-dashboard py-[2.25rem] flex flex-col gap-12 px-3 sm:px-[2rem]">
    <div class="flex items-center justify-between w-full flex-wrap gap-3">
    <h3 class="text-[#1D2939] text-2xl font-semibold">All Products</h3>
      <div>
      <ButtonComponent :label="isSmallerScreen ? '' : 'Add New Product'" color="white"
                       background-color="#1D2939">
        <svg xmlns="http://www.w3.org/2000/svg" width="32" height="33" viewBox="0 0 32 33" fill="none">
          <path d="M16 2.5C12.3008 2.54463 8.76575 4.03395 6.14985 6.64985C3.53395 9.26575 2.04463 12.8008 2 16.5C2.04463 20.1992 3.53395 23.7343 6.14985 26.3501C8.76575 28.966 12.3008 30.4554 16 30.5C19.6992 30.4554 23.2343 28.966 25.8501 26.3501C28.466 23.7343 29.9554 20.1992 30 16.5C29.9554 12.8008 28.466 9.26575 25.8501 6.64985C23.2343 4.03395 19.6992 2.54463 16 2.5ZM24 17.5H17V24.5H15V17.5H8V15.5H15V8.5H17V15.5H24V17.5Z" fill="#F9FAFB"/>
        </svg>
      </ButtonComponent>
      </div>
    </div>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 md:gap-5 xl:gap-6">
        <SingleProduct />
        <SingleProduct />
        <SingleProduct />
        <SingleProduct />
        <SingleProduct />
        <SingleProduct />
        <SingleProduct />
        <SingleProduct />
        <SingleProduct />
    </div>
  </section>
</template>
<script lang="ts" setup>
import ButtonComponent from "../../components/ButtonComponent.vue";
import SingleProduct from "../../components/admin-all-products/SingleProduct.vue";
import {onBeforeUnmount, onMounted, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();
const isSmallerScreen = ref(false);
const checkScreenSize = () => {
  isSmallerScreen.value = window.innerWidth <= 375;
}

const handle addNewPro

onMounted(() => {
  window.addEventListener("resize", checkScreenSize)
})

onBeforeUnmount(() => {
  window.removeEventListener("resize", checkScreenSize)
})
</script>