<template>
  <footer class="flex flex-col gap-8 px-[5vw] pb-4 pt-4 sm:pt-8 md:pt-[4.2rem]">
    <div
      class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between"
    >
      <router-link to="/" class="-ml-4 cursor-pointer">
        <LabelLogo
          text-color="#1D2939"
          inner-color="#1D2939"
          outer-color="#FFFFFF"
        />
      </router-link>
      <div class="hidden items-center gap-4 sm:flex">
        <FooterSocials social="facebook" />
        <FooterSocials social="instagram" />
        <FooterSocials social="twitter" />
        <FooterSocials social="pinterest" />
      </div>
    </div>
    <div
      class="grid grid-cols-1 gap-8 sm:grid-cols-2 md:grid-cols-3 md:gap-14 lg:gap-32"
    >
      <div class="flex flex-col gap-4 capitalize">
        <h5
          class="relative text-2xl text-[#1D2939] after:absolute after:-bottom-1 after:left-0 after:h-[1px] after:w-12 after:bg-[#1D2939] after:content-['']"
        >
          {{ company[0].label }}
        </h5>
        <router-link
          class="text-[#667085]"
          v-for="(item, index) in company.slice(1)"
          :key="index"
          :to="item.link"
        >
          {{ item.label }}
        </router-link>
      </div>

      <div class="flex flex-col gap-4 capitalize">
        <h5
          class="relative text-2xl text-[#1D2939] after:absolute after:-bottom-1 after:left-0 after:h-[1px] after:w-12 after:bg-[#1D2939] after:content-['']"
        >
          {{ customer[0].label }}
        </h5>
        <router-link
          class="text-[#667085]"
          v-for="(item, index) in customer.slice(1)"
          :key="index"
          :to="item.link"
        >
          {{ item.label }}
        </router-link>
      </div>

      <div class="flex flex-col gap-4 capitalize">
        <h5
          class="relative text-2xl text-[#1D2939] after:absolute after:-bottom-1 after:left-0 after:h-[1px] after:w-12 after:bg-[#1D2939] after:content-['']"
        >
          {{ orders[0].label }}
        </h5>
        <router-link
          class="text-[#667085]"
          v-for="(item, index) in orders.slice(1)"
          :key="index"
          :to="item.link"
        >
          {{ item.label }}
        </router-link>
      </div>
    </div>
    <div class="flex flex-col gap-2 sm:hidden">
      <span class="text-sm font-medium capitalize text-[#1D2939]"
        >follow us</span
      >
      <div class="flex items-center gap-8 sm:gap-4">
        <FooterSocials social="facebook" />
        <FooterSocials social="instagram" />
        <FooterSocials social="twitter" />
        <FooterSocials social="pinterest" />
      </div>
    </div>
    <span class="text-xs font-medium text-[#4F5B75] sm:text-base"
      >@ {{ year }} Klock Inc.</span
    >
  </footer>
</template>
<script setup lang="ts">
import LabelLogo from "@/assets/LabelLogo.vue";
import FooterSocials from "./FooterSocials.vue";
import { ref } from "vue";

const year = new Date().getFullYear();

const company = ref([
  { label: "company", link: "/company" },
  { label: "about us", link: "/about-us" },
  { label: "team", link: "/team" },
  { label: "collaboration", link: "/team" },
]);

const customer = ref([
  { label: "customer support", link: "/customer-support" },
  { label: "customer service", link: "/customer-service" },
  { label: "warranty", link: "/warranty" },
  { label: "FAQs", link: "/faqs" },
  { label: "collaboration", link: "/collaboration" },
  { label: "contact us", link: "/contact-us" },
]);

const orders = ref([
  { label: "order & return", link: "/order-and-return" },
  { label: "shipping", link: "/shipping" },
  { label: "returns", link: "/returns" },
  { label: "check gift card ballance", link: "/gift-card-ballance" },
  { label: "promotions", link: "/promotions" },
]);
</script>
