<template>
    <div class="hidden">
<!-- Token Expiry Checker  -->
    </div>
</template>
<script setup lang="ts">
import { useHandleLogoutAfterTokenExpires } from '@/lib/helperFunctions';
import { onMounted, onUnmounted } from 'vue';

const {startChecking, stopChecking} = useHandleLogoutAfterTokenExpires();
onMounted(() => {
    startChecking()
})

onUnmounted(() => {
      stopChecking();
    });
</script>