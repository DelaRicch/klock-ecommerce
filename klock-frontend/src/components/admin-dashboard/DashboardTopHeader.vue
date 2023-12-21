<template>
  <div
    class="relative flex h-[5.2rem] items-center justify-end gap-4 bg-[#F9FAFB] px-3 py-3 sm:px-[2rem] md:gap-10"
  >
    <div class="absolute left-8 top-[50%] translate-y-[-50%] md:hidden">
      <ButtonComponent is-icon @update:model-value="handleToggleDisplayMenu">
        <HamburgerMenu :is-open="displayMenu" />
      </ButtonComponent>
    </div>
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
        fill="none"
        height="24"
        viewBox="0 0 24 24"
        width="24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <g clip-path="url(#clip0_408_6447)">
          <path
            d="M15 19C15.0001 19.5046 14.8096 19.9906 14.4665 20.3605C14.1234 20.7305 13.6531 20.9572 13.15 20.995L13 21H11C10.4954 21.0002 10.0094 20.8096 9.63942 20.4665C9.26944 20.1234 9.04281 19.6532 9.00497 19.15L8.99997 19H15ZM12 2C13.8149 1.99997 15.559 2.70489 16.8642 3.96607C18.1694 5.22726 18.9337 6.94609 18.996 8.76L19 9V12.764L20.822 16.408C20.9015 16.567 20.9413 16.7429 20.9379 16.9206C20.9346 17.0984 20.8882 17.2727 20.8028 17.4286C20.7174 17.5845 20.5955 17.7174 20.4475 17.8158C20.2994 17.9143 20.1298 17.9754 19.953 17.994L19.838 18H4.16197C3.98413 18.0001 3.80894 17.957 3.65139 17.8745C3.49385 17.792 3.35865 17.6725 3.25739 17.5264C3.15613 17.3802 3.09182 17.2116 3.06997 17.0351C3.04813 16.8586 3.0694 16.6795 3.13197 16.513L3.17797 16.408L4.99997 12.764V9C4.99997 7.14348 5.73747 5.36301 7.05022 4.05025C8.36298 2.7375 10.1435 2 12 2Z"
            fill="black"
          />
        </g>
        <defs>
          <clipPath id="clip0_408_6447">
            <rect fill="white" height="24" width="24" />
          </clipPath>
        </defs>
      </svg>
    </ButtonComponent>
    <Avatar :alt="userProfile?.Name" :src="userProfile?.Photo" />
  </div>
</template>
<script lang="ts" setup>
import { useUserStore } from "@/stores/user";
import Avatar from "../Avatar.vue";
import ButtonComponent from "../ButtonComponent.vue";
import HamburgerMenu from "@/assets/HamburgerMenu.vue";
import { displayMenu } from "@/stores/resuableState";

const handleToggleDisplayMenu = () => {
  displayMenu.value = !displayMenu.value;
};

const {userProfile} = useUserStore();
</script>
