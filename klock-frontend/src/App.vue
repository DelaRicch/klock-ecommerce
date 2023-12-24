<template>
  <main>
    <router-view></router-view>

    <!--  Display notification when token is about to expire-->
    <TokenExpirationNotif />

    <!--  Snackbar toast notification-->
    <SnackbarComponent
      :success="success"
      :title="success ? 'success' : 'error'"
      :message="apiResponse"
    />

    <!-- Token expiry checker  -->
    <TokenExpiryChecker v-if="refreshToken.value" />
  </main>
</template>

<script setup lang="ts">
import SnackbarComponent from '@/components/SnackbarComponent.vue';
import TokenExpirationNotif from '@/components/TokenExpirationNotif.vue';
import {
  apiResponse,
  success,
} from "@/stores/resuableState";
import TokenExpiryChecker from './components/TokenExpiryChecker.vue';
import { useUserStore } from './stores/user';
import { storeToRefs } from 'pinia';

const {refreshToken} = storeToRefs(useUserStore());

</script>
