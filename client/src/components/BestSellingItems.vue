<template>
  <section class="flex flex-col gap-4 px-[5vw] py-8">
    <div class="flex justify-between">
      <div class="border-b border-black pb-1 text-xl font-medium capitalize">
        best selling watches
      </div>
      <span class="font-medium text-[#667085]">20 items</span>
    </div>
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 2xl:grid-cols-5 gap-12">
     <ItemComponent />
     <ItemComponent />
     <ItemComponent />
     <ItemComponent />
     <ItemComponent />
     <ItemComponent />
     <ItemComponent />
     <ItemComponent />
     <ItemComponent />
    </div>
    <div class="flex self-end">
        <NavigationComponent
            :total-pages="10"
        />
    </div>
  </section>
</template>
<script setup lang="ts">
import ItemComponent from './ItemComponent.vue';
import NavigationComponent from './NavigationComponent.vue';


</script>
