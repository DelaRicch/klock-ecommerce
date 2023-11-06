<script setup lang="ts">
defineProps({
  outerColor: {
    type: String,
required: true,
  },
  innerColor: {
    type: String,
required: true
  },
  textColor: {
    type: String,
required: true
  }
})
</script>

<template>
  <svg width="71" height="66" viewBox="0 0 71 66" fill="none" xmlns="http://www.w3.org/2000/svg">
    <rect x="70.2696" y="43.3837" width="13.6298" height="23"
          transform="rotate(114.997 70.2696 43.3837)" :fill="innerColor" :stroke="textColor"/>
    <rect x="62.6016" y="9.68301" width="13.6298" height="23"
          transform="rotate(60 62.6016 9.68301)" :fill="innerColor" :stroke="textColor"/>
    <rect x="29" y="1" width="20" height="64" :fill="innerColor" :stroke="outerColor"
          stroke-width="2"/>
    <path d="M13.879 15.7501L47.7458 33.6261L13.4733 50.7113L13.879 15.7501Z" :fill="innerColor"
          :stroke="outerColor" stroke-width="2"/>
  </svg>

</template>