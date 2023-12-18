<template>
  <HeaderComponent />
  <MobileHeaderComponent />
  <section class="px-[2.5vw] my-[3rem]">
    <h3 class="text-[#1D2939] w-max text-xl relative after:absolute after:-bottom-1 after:left-0 after:h-[1px] after:w-full after:bg-[#1D2939] after:content-['']">Account Settings</h3>
    <RadioButtons :radios="radios" />
    <div class="flex gap-2 mt-4 border-b-[1px] border-b-[#E4E7EC]">
        <button :class="isActiveTab ? 'text-[#1D2939]' : 'text-[#667085]'" 
        >General</button>
        <button>Password</button>
    </div>
    <div class="my-[4rem] w-[38rem] mx-auto flex flex-col gap-8">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-8"></div>
      </div>
    </div>
  </section>
  <FooterComponent />
</template>
<script setup lang="ts">
import FooterComponent from "@/components/FooterComponent.vue";
import HeaderComponent from "@/components/HeaderComponent.vue";
import MobileHeaderComponent from "@/components/MobileHeaderComponent.vue";
import RadioButtons from "@/components/RadioButtons.vue";
import { ref } from "vue";

const radios = [{ value: "man" }, { value: "woman" }, { value: "other" }];

const isActiveTab = ref(false);
</script>
