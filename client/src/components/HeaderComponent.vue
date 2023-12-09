<template>
  <header
    class="sticky top-0 z-50  hidden md:flex h-[5rem] md:px-[5vw] items-center justify-between bg-white"
  >
    <router-link to="/" class="cursor-pointer">
      <LabelLogo
        text-color="#1D2939"
        inner-color="#1D2939"
        outer-color="#FFFFFF"
      />
    </router-link>
    <nav>
      <ul class="flex items-center gap-6">
        <li
          :class="activeRoute === item.name ? 'after:w-full' : 'after:w-0'"
          class="relative text-[0.875rem] font-semibold capitalize text-[#667085] after:absolute after:bottom-[-10px] after:left-0 after:h-[2px] after:bg-[#1D2939] after:transition-all after:duration-300 after:ease-linear after:content-[''] hover:after:w-full"
          v-for="item in menuItems"
          :key="item.id"
        >
          <button
            @click="handleDisplayCategories"
            v-if="item.id === 'a0a0a0a0'"
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
          <router-link v-else :to="item.link!">{{ item.name }}</router-link>
        </li>
      </ul>
    </nav>
    <div class="flex items-center gap-3">
      <div class="flex items-center gap-2">
        <ButtonComponent is-icon>
          <svg
            fill="none"
            height="24"
            viewBox="0 0 24 24"
            width="24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M21 21.0002L16.657 16.6572M16.657 16.6572C17.3999 15.9143 17.9892 15.0324 18.3912 14.0618C18.7933 13.0911 19.0002 12.0508 19.0002 11.0002C19.0002 9.9496 18.7933 8.90929 18.3912 7.93866C17.9892 6.96803 17.3999 6.08609 16.657 5.34321C15.9141 4.60032 15.0322 4.01103 14.0615 3.60898C13.0909 3.20693 12.0506 3 11 3C9.94939 3 8.90908 3.20693 7.93845 3.60898C6.96782 4.01103 6.08588 4.60032 5.34299 5.34321C3.84266 6.84354 2.99979 8.87842 2.99979 11.0002C2.99979 13.122 3.84266 15.1569 5.34299 16.6572C6.84332 18.1575 8.87821 19.0004 11 19.0004C13.1218 19.0004 15.1567 18.1575 16.657 16.6572Z"
              stroke="#1D2939"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="1.5"
            />
          </svg>
        </ButtonComponent>
        <ButtonComponent is-icon>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="25"
            height="25"
            viewBox="0 0 25 25"
            fill="none"
          >
            <path
              d="M17.9341 18.825C18.4645 18.825 18.9732 19.0357 19.3483 19.4108C19.7234 19.7859 19.9341 20.2946 19.9341 20.825C19.9341 21.3554 19.7234 21.8641 19.3483 22.2392C18.9732 22.6143 18.4645 22.825 17.9341 22.825C17.4036 22.825 16.8949 22.6143 16.5199 22.2392C16.1448 21.8641 15.9341 21.3554 15.9341 20.825C15.9341 19.715 16.8241 18.825 17.9341 18.825ZM1.93408 2.825H5.20408L6.14408 4.825H20.9341C21.1993 4.825 21.4537 4.93035 21.6412 5.11789C21.8287 5.30543 21.9341 5.55978 21.9341 5.825C21.9341 5.995 21.8841 6.165 21.8141 6.325L18.2341 12.795C17.8941 13.405 17.2341 13.825 16.4841 13.825H9.03408L8.13408 15.455L8.10408 15.575C8.10408 15.6413 8.13042 15.7049 8.17731 15.7518C8.22419 15.7987 8.28778 15.825 8.35408 15.825H19.9341V17.825H7.93408C7.40365 17.825 6.89494 17.6143 6.51987 17.2392C6.1448 16.8641 5.93408 16.3554 5.93408 15.825C5.93408 15.475 6.02408 15.145 6.17408 14.865L7.53408 12.415L3.93408 4.825H1.93408V2.825ZM7.93408 18.825C8.46452 18.825 8.97322 19.0357 9.3483 19.4108C9.72337 19.7859 9.93408 20.2946 9.93408 20.825C9.93408 21.3554 9.72337 21.8641 9.3483 22.2392C8.97322 22.6143 8.46452 22.825 7.93408 22.825C7.40365 22.825 6.89494 22.6143 6.51987 22.2392C6.1448 21.8641 5.93408 21.3554 5.93408 20.825C5.93408 19.715 6.82408 18.825 7.93408 18.825ZM16.9341 11.825L19.7141 6.825H7.07408L9.43408 11.825H16.9341Z"
              fill="#1D2939"
            />
          </svg>
        </ButtonComponent>
        <ButtonComponent is-icon>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="21"
            height="20"
            viewBox="0 0 21 20"
            fill="none"
          >
            <path
              d="M11.0341 16.375L10.9341 16.475L10.8241 16.375C6.07408 12.065 2.93408 9.215 2.93408 6.325C2.93408 4.325 4.43408 2.825 6.43408 2.825C7.97408 2.825 9.47408 3.825 10.0041 5.185H11.8641C12.3941 3.825 13.8941 2.825 15.4341 2.825C17.4341 2.825 18.9341 4.325 18.9341 6.325C18.9341 9.215 15.7941 12.065 11.0341 16.375ZM15.4341 0.824997C13.6941 0.824997 12.0241 1.635 10.9341 2.905C9.84408 1.635 8.17408 0.824997 6.43408 0.824997C3.35408 0.824997 0.934082 3.235 0.934082 6.325C0.934082 10.095 4.33408 13.185 9.48408 17.855L10.9341 19.175L12.3841 17.855C17.5341 13.185 20.9341 10.095 20.9341 6.325C20.9341 3.235 18.5141 0.824997 15.4341 0.824997Z"
              fill="#101828"
            />
          </svg>
        </ButtonComponent>
      </div>

      <!--    Display User Avatar if logged in-->
      <Avatar
        v-if="
          userProfile?.Name !== '' &&
          userProfile !== undefined &&
          userProfile !== null
        "
        :src="userProfile?.Photo"
        :alt="userProfile?.Name"
      />

      <!--    Display login signup if not logged in -->
      <div v-else class="flex items-center gap-1 text-[#181CF9]">
        <router-link
          to="sign-in"
          class="transition-all duration-300 ease-linear hover:underline"
          >Login </router-link
        >/
        <router-link
          to="sign-up"
          class="transition-all duration-300 ease-linear hover:underline"
          >Register
        </router-link>
      </div>
    </div>
  </header>
  <CategoriesDropdown />
</template>
<script setup lang="ts">
import LabelLogo from "../assets/LabelLogo.vue";
import ButtonComponent from "@/components/ButtonComponent.vue";
import { useUserStore } from "@/store/store.ts";
import Avatar from "@/components/Avatar.vue";
import { displayCategories } from "@/store/resuableState.ts";
import { computed } from "vue";
import { useRoute } from "vue-router";
import CategoriesDropdown from "@/components/CategoriesDropdown.vue";

const activeRoute = computed(() => {
  return useRoute().name;
});

const userStore = useUserStore();
const userProfile = computed(() => userStore?.userProfile);

const handleDisplayCategories = () => {
  displayCategories.value = !displayCategories.value;
};

const menuItems = [
  {
    name: "home",
    link: "/",
    id: "5ba48fa3",
  },
  {
    name: "categories",
    id: "a0a0a0a0",
  },
  {
    name: "Blog",
    link: "/blog",
    id: "4096c3fe",
  },
  {
    name: "about",
    link: "/about",
    id: "9c3b2e28",
  },
  {
    name: "contact",
    link: "/contact",
    id: "a5a5a5a5",
  },
];
</script>
