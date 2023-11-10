
<template>
  <button
      @click="signInWithSocialAuth"
      class="outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC] p-[0.63rem] border border-[#D0D5DD] hover:bg-white transition duration-200 ease-linear">
    <SocialIcons :type="authType" />
  </button>
  <Teleport to="body">
    <Snackbar
        :show-snackbar="showSnackbar"
        :success="success"
        :title="success ? 'success' : 'error' "
        :message="apiResponse" />
  </Teleport>
</template>

<script setup lang="ts">
  import SocialIcons from "../assets/SocialIcons.vue";
 import {GoogleAuthProvider, getAuth, signInWithPopup} from "firebase/auth";
  import { socialAuthType} from "@/types";
  import { socialLogin} from "../api/user.ts";
  import {
    errorApiRequest,
    successApiRequest
  } from "../lib/helperFunctions.ts";
  import Snackbar from "../components/Snackbar.vue";
  import {apiResponse, showSnackbar, success} from "../store/resuableState.ts";


  const props = defineProps({
    authType: {
      type: String,
      required: true,
      validator: (value: string) => {
        return ['Facebook', 'Google', 'X'].includes(value);
      },
    },
  })

  const signInWithSocialAuth = () => {
    // Implement authentication logic based on the provided authType
    if (props.authType === 'Facebook') {
      // Facebook authentication logic
    }  else if (props.authType === 'X') {
      // Apple authentication logic
    } else if (props.authType === 'Google') {
      const provider = new GoogleAuthProvider();
      signInWithPopup(getAuth(), provider).then((response) => {
        const user: socialAuthType = {
          email: response.user.email,
          name: response.user.displayName,
          photo: response.user.photoURL,
          socialId: response.user.uid,
        }
        console.log(user)
        socialLogin(user)
            .then((res) => {
              successApiRequest(res);
            })
            .catch((err) => {
            errorApiRequest(err)
            });

      }).catch(err => {
        console.error(err);
      });

    }
  }
</script>
