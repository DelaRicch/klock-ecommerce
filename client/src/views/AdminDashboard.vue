<template>
  <div class="flex w-screen">
    <aside class="md:w-[25%] 2xl:w-[20%]"></aside>
    <div class="flex flex-col px-3 sm:px-[2rem] w-full md:w-[75%] 2xl:w-[80%] bg-red-300">
      <DashboardTopHeader />
      <section
          class="py-[2.25rem] flex flex-col gap-[2.5rem]">

        <div class="flex flex-wrap gap-[0.87rem] xl:gap-6 2xl:gap-8">
        <DashboardCard
            :title="'Total Orders'"
            :value="120.00"
            :percentage="34.7"
            :summary="'Compared to Jan 2022'"
        />
        <DashboardCard
            :title="'Active Orders'"
            :value="120.00"
            :percentage="34.7"
            :summary="'Compared to Jan 2022'"
        />
        <DashboardCard
            :title="'Shipped Orders'"
            :value="120.00"
            :percentage="34.7"
            :summary="'Compared to Jan 2022'"
        />
        </div>
        <div class="flex items-center justify-between">
          <div class="p-4 rounded-lg">

          </div>
        </div>
      </section>
    </div>
  </div>
</template>
<script setup lang="ts">

import DashboardTopHeader from "../components/admin-dashboard/DashboardTopHeader.vue";
import DashboardCard from "../components/admin-dashboard/DashboardCard.vue";
</script>