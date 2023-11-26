<template>
  <aside
      class="w-0 hidden md:block md:w-[25%] 2xl:w-[20%] h-full bg-[#F9FAFB] border-r border-slate-300">
    <div class="w-full flex justify-center h-max mt-3">
      <LabelLogo text-color="#1D2939" inner-color="#1D2939" outer-color="#F9FAFB" />
    </div>
    <div class="mt-[3.5rem] px-0.5 w-full flex flex-col gap-[3rem]">
        <DashboardNavButton
            v-for="button in navButtons"
            :key="button.id"
            :background-color="''"
            :color="''"
            :name="button.name"
            :label="button.title"
            :route="button.route"
            :current-route="currentRoute"
            :parent-route="parentRoute"
            />
    </div>
  </aside>
</template>
<script setup lang="ts">
import LabelLogo from "../../assets/LabelLogo.vue";
import DashboardNavButton from "../../components/admin-dashboard/DashboardNavButton.vue";
import {useRoute} from "vue-router";
import {computed} from "vue";

const navButtons = [
  {id: 1, title: 'Dashboard', name: 'dashboard', route: ''},
  {id: 2, title: 'All Products', name: 'all-products', route: 'all-products'},
  {id: 3, title: 'Order List', name: 'order-list', route: 'order-list'},
]

const route = useRoute()
const currentRoute = computed(() => route.name ?? '');
const parentRoute = computed(() => route.meta?.parent ?? '')

</script>