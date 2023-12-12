<template>
  <Teleport to="body">
    <Transition name="fade">
      <div
        v-if="displayUserDropdownMenu"
        class="fixed right-0 top-0 z-50 h-screen w-screen bg-black/50 md:hidden"
      >
        <div
          class="absolute right-0 top-0 flex h-[40%] w-4/6 flex-col justify-between bg-white px-8 pb-8 pt-5"
        >
          <div class="flex flex-col">
            <div class="flex self-end">
              <ButtonComponent
                is-icon
                @update:model-value="displayUserDropdownMenu = false"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="14"
                  height="14"
                  viewBox="0 0 14 14"
                  fill="none"
                >
                  <path
                    d="M14 1.41L12.59 0L7 5.59L1.41 0L0 1.41L5.59 7L0 12.59L1.41 14L7 8.41L12.59 14L14 12.59L8.41 7L14 1.41Z"
                    fill="black"
                  />
                </svg>
              </ButtonComponent>
            </div>
            <ul class="flex flex-col gap-6">
              <li
                class="relative text-sm w-max font-bold capitalize text-[#1D2939] after:absolute after:bottom-[-5px] after:left-0 after:h-[2px] after:bg-[#1D2939] after:transition-all after:duration-300 after:ease-linear after:content-[''] hover:after:w-2/5"
                v-for="item in menuItems"
                :key="item.id"
              >
                <button
                  v-if="item.id === '5ba48fa3'"
                  class="flex cursor-pointer items-center gap-1 capitalize"
                >
                  <span>{{ item.name }}</span>
                  <svg
                    :class="{
                      'rotate-180 transform transition-all duration-300 ease-linear':
                        displayCategories,
                      'rotate-0 transition-all duration-300 ease-linear':
                        !displayCategories,
                    }"
                    xmlns="http://www.w3.org/2000/svg"
                    width="20"
                    height="20"
                    viewBox="0 0 20 20"
                    fill="none"
                  >
                    <path
                      d="M6 9L12 15L18 9"
                      stroke="#667085"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    />
                  </svg>
                </button>
                <router-link v-else :to="item.link!">{{
                  item.name
                }}</router-link>
              </li>
            </ul>
          </div>
          <button
            @click="handleLogoutFunc"
            class="w-max cursor-pointer font-bold text-[#EF0816] relative after:absolute after:bottom-[-5px] after:left-0 after:h-[2px] after:bg-[#EF0816] after:transition-all after:duration-300 after:ease-linear after:content-[''] hover:after:w-full"
          >
            Log out
          </button>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>
<script setup lang="ts">
import {
  displayCategories,
  displayUserDropdownMenu,
} from "@/store/resuableState";
import ButtonComponent from "./ButtonComponent.vue";
import { handleLogout } from "@/lib/helperFunctions";
import { ref } from "vue";

const menuItems = ref<{ name: string; link?: string; id: string }[]>([
  {
    name: "my orders",
    id: "5ba48fa3",
  },
  {
    name: "shipping address",
    id: "a0a0a0a0",
    link: "/shipping-address",
  },
  {
    name: "payment methods",
    link: "/payment-methods",
    id: "4096c3fe",
  },
  {
    name: "my reviews",
    link: "/my-reviews",
    id: "9c3b2e28",
  },
  {
    name: "account settings",
    link: "/account-settings",
    id: "a5a5a5a5",
  },
]);

const handleLogoutFunc = () => {
  handleLogout();
  displayUserDropdownMenu.value = false;
};

const handleDisplayCategories = () => {
  displayCategories.value = !displayCategories.value;
};
</script>

<style lang="css" scoped>
.fade-enter-active,
.fade-leave-active {
  @apply transition-all duration-300 ease-linear;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateX(100%);
}
</style>
