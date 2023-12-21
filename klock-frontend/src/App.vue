<template>
  <main>
    <router-view></router-view>

    <!--  Display notification when token is about to expire-->
    <TokenExpirationNotif />

    <!--  Snackbar toast notification-->
    <Snackbar
      :show-snackbar="showSnackbar"
      :success="success"
      :title="success ? 'success' : 'error'"
      :message="apiResponse"
    />
  </main>
</template>

<script setup lang="ts">
import Snackbar from '@/components/Snackbar.vue';
import TokenExpirationNotif from '@/components/TokenExpirationNotif.vue';
import {
  apiResponse,
  showSnackbar,
  success,
} from "@/stores/resuableState";
</script>
