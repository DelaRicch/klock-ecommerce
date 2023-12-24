<template>
  <Transition name="fade">
    <div v-if="showTokenExpiry">
      <ModalComponent persistence>
        <template #header>
          <div class="mx-auto flex items-center gap-3">
            <span>Token expires in</span>
            <span class="text-lg font-semibold text-[#1D2939]">{{
              formattedExpiration
            }}</span>
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
          <div class="mt-2 flex flex-col items-center gap-3 sm:flex-row">
            <ButtonComponent
              @update:model-value="handleCloseModal"
              :model-value="buttonValue"
              label="Cancel"
              background-color="#FBE6E6"
              color="#D20101"
            />
            <ButtonComponent
              @update:model-value="handleStayOnPage"
              :model-value="buttonValue"
              label="Continue shopping"
              background-color="#1D2939"
              color="white"
            />
          </div>
        </template>
      </ModalComponent>
    </div>
  </Transition>
</template>
<script lang="ts" setup>
import TypingMan from "../assets/typing-man.json";
import { Vue3Lottie } from "vue3-lottie";
import { onUnmounted } from "vue";
import { requestNewToken } from "@/api/user";
import { buttonValue, showTokenExpiry } from "@/stores/resuableState";
import { errorApiRequest, successApiRequest, tokenExpiryCalculator } from "@/lib/helperFunctions";
import { useUserStore } from "@/stores/user";
import ModalComponent from "./ModalComponent.vue";
import ButtonComponent from "./ButtonComponent.vue";
import { storeToRefs } from "pinia";


const {accessToken} = storeToRefs(useUserStore());

const { formattedExpiration, countdownInterval, remainingTime } =
  tokenExpiryCalculator(accessToken.value.expiry);

const handleCloseModal = (newValue: boolean) => {
  buttonValue.value = newValue;
  showTokenExpiry.value = false;
  clearInterval(countdownInterval);
  remainingTime.value = 0;
};

const handleStayOnPage = (newValue: boolean) => {
  requestNewToken()
  buttonValue.value = newValue;
  showTokenExpiry.value = false;
  remainingTime.value = 0;
  clearInterval(countdownInterval);

  requestNewToken()
    .then((res) => {
      successApiRequest(res);
    })
    .catch((err) => {
      errorApiRequest(err);
    });
};

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
