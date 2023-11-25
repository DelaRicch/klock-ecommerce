<template>
  <div class="flex items-center">
    <div class="" v-for="(bc, index) in breadcrumbs">
    <router-link :to="bc.path"
                 :class="{'text-[#1D2939] bg-[#BDBDBD]': index === breadcrumbs.length -1,
                 'bg-[#F5F5F5] text-[#BDBDBD]': index !== breadcrumbs.length -1,
                 }"
                 class="font-medium text-sm rounded-2xl px-4 py-1">
      {{ bc.label }}</router-link>
    <span v-if="index !== breadcrumbs.length -1" class="mx-1 text-[#BDBDBD]">/</span>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import {PropType} from "vue";

  interface BreadcrumbsType {
    path: string
    label: string
  }

  defineProps({
    breadcrumbs: {
      type: Array as PropType<BreadcrumbsType[]>,
      required: true
    }
  })
</script>