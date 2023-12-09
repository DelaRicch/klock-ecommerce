<template>
  <aside
    class="hidden h-full w-0 flex-col justify-between border-r border-slate-300 bg-[#F9FAFB] pb-20 md:flex md:w-[25%] 2xl:w-[20%]"
  >
    <div class="h-max">
      <div class="mt-3 flex h-max w-full justify-center">
        <LabelLogo
          text-color="#1D2939"
          inner-color="#1D2939"
          outer-color="#F9FAFB"
        />
      </div>
      <div class="mt-[3.5rem] flex w-full flex-col gap-[3rem] px-0.5">
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
    </div>
    <div class="transition-all duration-300 ease-linear hover:bg-slate-200">
      <ButtonComponent label="Logout" @update:model-value="handleLogout">
        <span
          ><svg
            xmlns="http://www.w3.org/2000/svg"
            width="21"
            height="19"
            viewBox="0 0 21 19"
            fill="none"
          >
            <path
              d="M15.05 4.5L13.64 5.91L16.22 8.5H6.05005V10.5H16.22L13.64 13.08L15.05 14.5L20.05 9.5L15.05 4.5ZM2.05005 2.5H10.05V0.5H2.05005C0.950049 0.5 0.0500488 1.4 0.0500488 2.5V16.5C0.0500488 17.6 0.950049 18.5 2.05005 18.5H10.05V16.5H2.05005V2.5Z"
              fill="black"
            /></svg
        ></span>
      </ButtonComponent>
    </div>
  </aside>
</template>
<script setup lang="ts">
import LabelLogo from "../../assets/LabelLogo.vue";
import DashboardNavButton from "../../components/admin-dashboard/DashboardNavButton.vue";
import { useRoute } from "vue-router";
import { computed } from "vue";
import ButtonComponent from "@/components/ButtonComponent.vue";
import { handleLogout } from "@/lib/helperFunctions.ts";

const navButtons = [
  { id: 1, title: "Dashboard", name: "dashboard", route: "" },
  { id: 2, title: "All Products", name: "all-products", route: "all-products" },
  { id: 3, title: "Order List", name: "order-list", route: "order-list" },
];

const route = useRoute();
const currentRoute = computed(() => route.name ?? "");
const parentRoute = computed(() => route.meta?.parent ?? "");
</script>
