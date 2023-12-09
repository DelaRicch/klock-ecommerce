<template>
  <main>
    <router-view></router-view>

    <!--  Display notification when token is about to expire-->
    <TokenExpirationNotif :show-token-expiry="showTokenExpiry" />

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
import TokenExpirationNotif from "./components/TokenExpirationNotif.vue";
import {
  apiResponse,
  showSnackbar,
  showTokenExpiry,
  success,
} from "./store/resuableState.ts";
import Snackbar from "./components/Snackbar.vue";
</script>
