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
  </main>
</template>

<script setup lang="ts">
import SnackbarComponent from '@/components/SnackbarComponent.vue';
import TokenExpirationNotif from '@/components/TokenExpirationNotif.vue';
import {
  apiResponse,
  success,
} from "@/stores/resuableState";
</script>
