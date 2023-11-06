<script setup lang="ts">
  import SocialIcons from "../assets/SocialIcons.vue";

  const props = defineProps({
    authType: {
      type: String,
      required: true,
      validator: (value: string) => {
        return ['Facebook', 'Google', 'Apple'].includes(value);
      },
    },
  })

  const signInWithSocialAuth = () => {
    // Implement authentication logic based on the provided authType
    if (props.authType === 'Facebook') {
      // Facebook authentication logic
    } else if (props.authType === 'Google') {
      // Google authentication logic
    } else if (props.authType === 'Apple') {
      // Apple authentication logic
    }
  }
</script>

<template>
  <button
      @click="signInWithSocialAuth"
      class="outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC] p-[0.63rem] border border-[#D0D5DD] hover:bg-white transition duration-200 ease-linear">
    <SocialIcons :type="authType" />
  </button>
</template>