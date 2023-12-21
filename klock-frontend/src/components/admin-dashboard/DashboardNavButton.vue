<template>
  <button
    @click="handleSelectPage(name)"
    :class="
      currentRoute === name || parentRoute === name
        ? 'bg-[#155EEF] text-white'
        : 'bg-transparent text-[#1D2939]'
    "
    class="flex h-[3.2rem] items-center justify-center gap-3.5 rounded-lg capitalize outline-offset-4 outline-blue-600"
  >
    <DashboardIcon
      :name="name"
      :color="
        currentRoute === name || parentRoute === name ? 'white' : '#1D2939'
      "
    />
    {{ label }}
  </button>
</template>
<script lang="ts" setup>
import DashboardIcon from "../../assets/DashboardIcon.vue";
import { useRouter } from "vue-router";

const emit = defineEmits(["update:modelValue"]);

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
  route: {
    type: String,
    required: true,
  },
  currentRoute: {
    // type: String,
    required: true,
  },
  parentRoute: {
    required: false,
  },
});

const router = useRouter();

const handleSelectPage = (name: string) => {
  emit("update:modelValue", name);
  router.push({ name });
};
</script>
