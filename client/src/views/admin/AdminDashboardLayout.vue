<template>
  <div class="flex w-screen h-screen">
    <DashboardSidebar />
    <div
        class="flex flex-col w-full md:w-[75%] 2xl:w-[80%]">
      <DashboardTopHeader />
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup lang="ts">
import DashboardSidebar from "../../components/admin-dashboard/DashboardSidebar.vue";
import DashboardTopHeader from "../../components/admin-dashboard/DashboardTopHeader.vue";
</script>