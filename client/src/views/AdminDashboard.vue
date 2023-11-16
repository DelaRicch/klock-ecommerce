<template>
  <div class="flex w-screen h-screen">
    <aside class="md:w-[25%] 2xl:w-[20%]"></aside>
    <div
        class="flex flex-col w-full md:w-[75%] 2xl:w-[80%]">
      <DashboardTopHeader />
      <section
          class="main-dashboard py-[2.25rem] flex flex-col gap-[2.5rem] px-3 sm:px-[2rem]">

        <div class="grid gr-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-[0.87rem] xl:gap-6 2xl:gap-8">
        <DashboardCard
            :percentage="34.7"
            :summary="'Compared to Jan 2022'"
            :title="'Total Orders'"
            :value="120.00"
        />
        <DashboardCard
            :percentage="34.7"
            :summary="'Compared to Jan 2022'"
            :title="'Active Orders'"
            :value="120.00"
        />
        <DashboardCard
            :percentage="34.7"
            :summary="'Compared to Jan 2022'"
            :title="'Shipped Orders'"
            :value="120.00"
        />
        </div>
        <div class="flex flex-col md:flex-row items-center md:justify-between">
<DropdownComponent :options="options" @selected-option="handleSelectedOption" />

          <div
              class="p-4 rounded-lg border border-slate-400 flex flex-col gap-4 w-full md:w-[22.5rem]">
            <h3 class="text-[1.25rem] font-semibold">Best Sellers</h3>
<SingleItem
    :image="Watch"
    :price="50"
    :total-sales="5612"
    description="DV One automatic 43mm"
    title="Breitling leather"
/>
            <SingleItem
    :image="Watch"
    :price="50"
    :total-sales="5612"
    description="DV One automatic 43mm"
    title="Breitling leather"
/><SingleItem
    :image="Watch"
    :price="50"
    :total-sales="5612"
    description="DV One automatic 43mm"
    title="Breitling leather"
/>
          </div>
        </div>
<!--      <div class="my-4">-->
<!--      <PaginationComponent :total-pages="20" @update:current-page="handleDisplayCurrentPage" />-->
<!--      </div>-->
      </section>
    </div>
  </div>
</template>
<script lang="ts" setup>

import DashboardTopHeader from "../components/admin-dashboard/DashboardTopHeader.vue";
import DashboardCard from "../components/admin-dashboard/DashboardCard.vue";
import SingleItem from "../components/admin-dashboard/SingleItem.vue";
import Watch from "../assets/watch.png";
import DropdownComponent from "../components/DropdownComponent.vue";

const options = ['Today', 'This week', 'This month', 'Last month', 'This year', 'Last year'];

const handleSelectedOption = (option: string) => {
  console.log(option);
}

const handleDisplayCurrentPage = (page: number) => {
  console.log(page);
}

</script>
<style lang="css" scoped>
.main-dashboard {
min-height: calc(100vh - 5rem);
  overflow-y: auto;
}
</style>