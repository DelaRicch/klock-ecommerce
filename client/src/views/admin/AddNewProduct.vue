<template>
  <section
      class="main-dashboard py-[2.25rem] flex flex-col gap-5 px-3 sm:px-[2rem]">

  <BreadCrumbs :breadcrumbs="breadcrumbs" />
  <h3>Add New Product</h3>
  </section>
</template>
<script setup lang="ts">
import BreadCrumbs from "../../components/BreadCrumbs.vue";

const breadcrumbs = [
  {label: "Home", path: '/dashboard'},
  {label: "Product", path: '/dashboard/all-products'},
  {label: "New Product", path: ''},
]
</script>