<template>
  <router-view></router-view>

  <!--  Display notification when token is about to expire-->
  <TokenExpirationNotif :show-token-expiry="showTokenExpiry" />

</template>

<script setup lang="ts">
import TokenExpirationNotif from "./components/TokenExpirationNotif.vue";
import {showTokenExpiry} from "./store/resuableState.ts";

</script>
