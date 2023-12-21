<template>
  <HeaderComponent />
  <MobileHeaderComponent />

  <section class="px-[2.5vw] my-[3rem]">
    <h3 class="text-[#1D2939] w-max text-xl relative after:absolute after:-bottom-1 after:left-0 after:h-[1px] after:w-full after:bg-[#1D2939] after:content-['']">Account Settings</h3>
    <div class="flex gap-2 mt-4 border-b-[1px] border-b-[#E4E7EC]">
        <button @click="activeTab = 'general'" 
        :class="activeTab === 'general' ? 'text-[#1D2939] relative after:absolute after:-bottom-0 after:left-0 after:h-[1px] after:w-full after:bg-[#1D2939]' : 'text-[#667085]'" 
        >General</button>
        <button @click="activeTab = 'password'" 
        :class="activeTab === 'password' ? 'text-[#1D2939] relative after:absolute after:-bottom-0 after:left-0 after:h-[1px] after:w-full after:bg-[#1D2939]' : 'text-[#667085]'" 
        >Password</button>
    </div>
    <div class="my-[4rem] w-[38rem] mx-auto flex flex-col gap-9">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-8">
          <Avatar is-upload :alt="userProfile?.Name" />
          <div class="flex flex-col">
            <span class="text-xl font-medium" >{{ userProfile?.Name }}</span>
            <span class="text-[#667085] w-36" >{{ userProfile?.Email }}</span>
          </div>
        </div>
        <div>
          <ButtonComponent color="#EF0816" border-color="#EF0816" border label="Delete Account" background-color="#ef08160d" />
        </div>
      </div>

      <GeneralSettings v-if="activeTab === 'general'" />
    </div>
  </section>
  <FooterComponent />
</template>
<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import ButtonComponent from "@/components/ButtonComponent.vue";
import FooterComponent from "@/components/FooterComponent.vue";
import HeaderComponent from "@/components/HeaderComponent.vue";
import MobileHeaderComponent from "@/components/MobileHeaderComponent.vue";
import GeneralSettings from "@/components/settings/GeneralSettings.vue";
import { useUserStore } from "@/store/store";
import { ref } from "vue";

const activeTab = ref("general");

const {userProfile} = useUserStore();
</script>
