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
        <div class="flex flex-col lg:flex-row gap-6 md:gap-4 md:justify-between xl:gap-8">
          <SalesGraph />
          <div
              class="p-4 rounded-lg border border-slate-400 flex flex-col gap-4 w-full">
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
/>
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
/>
          </div>
        </div>
<!--        <div class="h-[40rem] border w-full">-->
        <table-component
            :title="'Recent Orders'"
            :headers="tableHeaders"
            :items="tableItems"
            :columnWidths="tableColumnWidths"
        />
<!--        </div>-->
      </section>
    </div>
  </div>
</template>
<script lang="ts" setup>
import DashboardTopHeader from "../components/admin-dashboard/DashboardTopHeader.vue";
import DashboardCard from "../components/admin-dashboard/DashboardCard.vue";
import SingleItem from "../components/admin-dashboard/SingleItem.vue";
import Watch from "../assets/watch.png";
import SalesGraph from "../components/admin-dashboard/SalesGraph.vue";
import TableComponent from "../components/TableComponent.vue";
import {ProductProps} from "@/types";

// Table section
const tableColumnWidths = {
  product: "25%",
  order_id: "17%",
  date: "12%",
  customer_name: "25%",
  status: "10%",
  amount: "7%",
};

const tableHeaders = [
  {key: 'product', label: 'Product'},
  {key: 'order_id', label: 'Order ID'},
  {key: 'date', label: 'Date'},
  {key: 'customer_name', label: 'Customer Name'},
  {key: 'status', label: 'Status'},
  {key: 'amount', label: 'Amount'},
]

const tableItems: ProductProps[] = [
  {
    id: 1,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 2,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 3,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 4,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 5,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 6,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 7,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 8,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 9,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 10,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 11,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
  {
    id: 12,
    product: "Breitling leather",
    order_id: '123RGR231567Y',
    date: '12/08/2023',
    customer_name: 'Ekow Smith',
    status: 'Completed',
    amount: '50'
  },
]

</script>
<style lang="css" scoped>
.main-dashboard {
min-height: calc(100vh - 5rem);
  overflow-y: auto;
}
</style>