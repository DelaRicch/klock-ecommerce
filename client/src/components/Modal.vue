<template>
  <Teleport to="body">
    <div
        @click.self="handleCloseModal"
        class="absolute top-0 left-0 w-screen h-screen flex items-center justify-center bg-slate-400/40 backdrop-blur-[4px] z-[100]">
    <div
        :class=" wiggle && 'animate-wiggle' "
        class=" shadow-black/40 shadow-lg rounded-lg p-4 flex flex-col min-h-[10rem] min-w-[90vw] md:min-w-[30rem] bg-white">
        <slot name="header"></slot>
      <slot name="content"></slot>
      <slot name="footer"></slot>
    </div>
    </div>
  </Teleport>
</template>
<script lang="ts" setup>
const props = defineProps({
  persistence: {
    type: Boolean,
  required: true,
  }
})
import {showTokenExpiry} from "../store/resuableState.ts";
import {ref} from "vue";

const wiggle = ref(false);


const handleCloseModal = () => {
if (props.persistence) {
    showTokenExpiry.value = true;
  wiggle.value = true;
  setTimeout(() => {
    wiggle.value = false;
  }, 600);
} else {
  showTokenExpiry.value = false;
}
}

</script>

<style lang="css" scoped>
@keyframes zoomInOut {
  0% {
    transform: scale(1);;
  }
  50% {
    transform: scale(1.03);
  }
  100% {
    transform: scale(1);
  }
}

.animate-wiggle {
  animation: zoomInOut .2s linear;
}
</style>
