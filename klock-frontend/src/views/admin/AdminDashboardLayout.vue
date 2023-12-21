<template>
  <div class="relative flex h-screen w-screen">
    <DashboardSidebar />
    <MobileSidebarMenu />
    <div class="flex w-full flex-col md:w-[75%] 2xl:w-[80%]">
      <DashboardTopHeader />
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup lang="ts">
import DashboardSidebar from "@/components/admin-dashboard/DashboardSidebar.vue";
import DashboardTopHeader from "@/components/admin-dashboard/DashboardTopHeader.vue";
import MobileSidebarMenu from "@/components/admin-dashboard/MobileSidebarMenu.vue";
</script>
