<template>
  <Transition name="fade">
    <aside
        v-if="displayMenu"
        class="absolute left-0 top-[4.5rem] w-8/12 sm:w-6/12 min-h-[50vh] bg-white shadow-black/50 shadow-md backdrop-blur-sm md:w-0 md:hidden z-50 overflow-hidden flex flex-col justify-between py-3 rounded-br-md"
    >

            <div>
              <div class="w-full -ml-3 flex justify-center">
                <Logo text-color="#1D2939" inner-color="#1D2939" outer-color="#F9FAFB" />
              </div>
              <div class="mt-[3.5rem] px-0.5 w-full flex flex-col gap-5">
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
                    @update:model-value="displayMenu = false"
                />
              </div>
            </div>
            <div class="hover:bg-slate-200 transition-all duration-300 ease-linear mt-10 sm:mt-4">
              <ButtonComponent label="Logout" @update:model-value="">
            <span><svg xmlns="http://www.w3.org/2000/svg" width="21" height="19" viewBox="0 0 21 19" fill="none">
        <path d="M15.05 4.5L13.64 5.91L16.22 8.5H6.05005V10.5H16.22L13.64 13.08L15.05 14.5L20.05 9.5L15.05 4.5ZM2.05005 2.5H10.05V0.5H2.05005C0.950049 0.5 0.0500488 1.4 0.0500488 2.5V16.5C0.0500488 17.6 0.950049 18.5 2.05005 18.5H10.05V16.5H2.05005V2.5Z" fill="black"/>
      </svg></span>
              </ButtonComponent>
      </div>
    </aside>
  </Transition>
</template>

<script setup>
import { displayMenu } from "@/store/resuableState.ts";
import ButtonComponent from "@/components/ButtonComponent.vue";
import Logo from "@/assets/Logo.vue";
import DashboardNavButton from "@/components/admin-dashboard/DashboardNavButton.vue";
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

<style lang="css" scoped>

.fade-enter-active,
.fade-leave-active {
  @apply transition-all duration-300 ease-linear;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateX(-100%);
}
</style>
