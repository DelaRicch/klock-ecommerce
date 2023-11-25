<template>
  <button
      @click="handleSelectPage(name)"
      :class="(currentRoute === name || parentRoute === name) ? 'bg-[#155EEF] text-white' :
      'bg-transparent text-[#1D2939]'"
          class="capitalize rounded-lg h-[3.2rem] flex items-center justify-center gap-3.5 outline-offset-4 outline-blue-600">
    <DashboardIcon :name="name" :color="(currentRoute === name || parentRoute === name) ? 'white' : '#1D2939' "/>
    {{ label }}
  </button>
</template>
<script lang="ts" setup>
import DashboardIcon from "../../assets/DashboardIcon.vue";
import {useRouter} from "vue-router";

defineProps({
  label: {
    type: String,
    required: true,
  },
  name: {
    type: String,
    required: true,
  },
  color: {
    type: String,
    required: true,
  },
  backgroundColor: {
    type: String,
    required: true,
  },
  route:{
    type: String,
    required: true,
  },
  currentRoute: {
    // type: String,
    required: true,
  },
  parentRoute: {
    required: false
  }
})

const router = useRouter()


const handleSelectPage = (name: string) => {
  router.push({name})
}

</script>