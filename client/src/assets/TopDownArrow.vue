<template>
  <svg
    v-if="salesIncrease"
    xmlns="http://www.w3.org/2000/svg"
    width="16"
    height="17"
    viewBox="0 0 16 17"
    fill="none"
  >
    <path
      d="M3.5 8.125L8 3.625L12.5 8.125M8 4.25V13.375"
      stroke="#05CD99"
      stroke-width="1.5"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
  </svg>

  <svg
    v-else
    class="rotate-180 transform"
    xmlns="http://www.w3.org/2000/svg"
    width="16"
    height="17"
    viewBox="0 0 16 17"
    fill="none"
  >
    <path
      d="M3.5 8.125L8 3.625L12.5 8.125M8 4.25V13.375"
      stroke="#FF3535"
      stroke-width="1.5"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
  </svg>
</template>

<script lang="ts" setup>
defineProps({
  salesIncrease: {
    type: Boolean,
    default: true,
  },
});
</script>
