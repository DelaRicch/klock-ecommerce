<template>
  <Transition name="fade">
<div v-if="showTokenExpiry">
  <Modal :persistence="true">
    <template #header>
      <div class="flex items-center gap-3 mx-auto">
        <span>Token expires in</span>
          <span class="text-[#1D2939] font-semibold text-lg">{{ formattedExpiration }}</span>
      </div>
    </template>
    <template #content>
      <div class="my-2">
        <Vue3Lottie
            :animationData="TypingMan"
            :height="200"
            :width="200"
            :loop="true"
            :speed="1"
            :autoPlay="true"
            :pauseAnimation="false"
            :pauseOnHover="false"
        />
      </div>
    </template>
    <template #footer>
      <div class="flex items-center gap-3 mt-2 flex-col sm:flex-row">
        <ButtonComponent
            @update:model-value="handleCloseModal"
            :model-value="buttonValue"
            :label="'Cancel'"
            background-color="#FBE6E6"
            color="#D20101" />
        <ButtonComponent
            @update:model-value="handleStayOnPage"
            :model-value="buttonValue"
            :label="'Stay on page'"
            background-color="#1D2939"
            color="white" />
      </div>
    </template>
  </Modal>
</div>

  </Transition>
</template>
<script lang="ts" setup>
import ButtonComponent from "../components/ButtonComponent.vue";
import Modal from "../components/Modal.vue";
import TypingMan from '../assets/typing-man.json';
import { Vue3Lottie } from "vue3-lottie";
import { useUserStore } from "../store/store";
import {computed, onUnmounted} from "vue";
import {buttonValue, showTokenExpiry} from "../store/resuableState";
import {tokenExpiryCalculator} from "../lib/helperFunctions.ts";
import {requestNewToken} from "../api/user.ts";

 defineProps({
  showTokenExpiry: {
    type: Boolean,
    required: true
  }
})

const userStore = useUserStore();
const tokenExpiration = computed(() => userStore.accessToken.expiry);

const {formattedExpiration, countdownInterval, remainingTime} =
    tokenExpiryCalculator(tokenExpiration);


const handleCloseModal = (newValue: boolean) => {
  console.log(remainingTime.value)
  buttonValue.value = newValue;
    showTokenExpiry.value = false;
  clearInterval(countdownInterval);
  remainingTime.value = 0;
  console.log(remainingTime.value)
}

const handleStayOnPage = (newValue: boolean) => {
  buttonValue.value = newValue;
  showTokenExpiry.value = false;
  remainingTime.value = 0;
  clearInterval(countdownInterval)

  requestNewToken().then(res => {
    console.log(res)
  }).catch(err => {
    console.error(err)
  })
}

onUnmounted(() => {
  clearInterval(countdownInterval);
});
</script>

<style lang="css" scoped>
.fade-enter-active,
.fade-leave-active {
  @apply transition-all duration-300 ease-linear;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>