<template>
  <button
    @click="signInWithSocialAuth"
    class="border border-[#D0D5DD] p-[0.63rem] outline-offset-2 transition duration-200 ease-linear hover:bg-white focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]"
  >
    <SocialIcons :type="authType" />
  </button>
</template>

<script setup lang="ts">
import {
  GoogleAuthProvider,
  FacebookAuthProvider,
  getAuth,
  signInWithPopup,
} from "firebase/auth";
import { socialAuthType } from "@/types";
import { useRouter } from "vue-router";
import SocialIcons from "@/assets/SocialIcons.vue";
import { getUserProfile, socialLogin } from "@/api/user";
import { errorApiRequest, handleUserProfile, successApiRequest } from "@/lib/helperFunctions";

const router = useRouter();

const props = defineProps({
  authType: {
    type: String,
    required: true,
    validator: (value: string) => {
      return ["Facebook", "Google"].includes(value);
    },
  },
});

const signInWithSocialAuth = () => {
  // Implement authentication logic based on the provided authType
  if (props.authType === "Facebook") {
    const provider = new FacebookAuthProvider();
    provider.addScope("email");
    signInWithPopup(getAuth(), provider)
      .then((response) => {
        console.log(response);
        const user: socialAuthType = {
          email: response?.user?.providerData[0]?.email!,
          name: response?.user?.providerData[0]?.displayName!,
          photo: response?.user?.providerData[0]?.photoURL!,
          socialId: response?.user?.uid,
          provider: response?.providerId!,
        };
        socialLogin(user)
          .then((res) => {
            successApiRequest(res);
          })
          .then(() => {
            getUserProfile().then((res) => {
              handleUserProfile(res.user);
              router.push("/dashboard");
            });
          })
          .catch((err) => {
            errorApiRequest(err);
          });
      })
      .catch((err) => {
        console.error(err);
      });
  } else if (props.authType === "Google") {
    const provider = new GoogleAuthProvider();
    provider.addScope("email");
    signInWithPopup(getAuth(), provider)
      .then((response) => {
        const user: socialAuthType = {
          email: response?.user?.providerData[0]?.email!,
          name: response?.user?.providerData[0]?.displayName!,
          photo: response?.user?.providerData[0]?.photoURL!,
          socialId: response?.user?.uid,
          provider: response?.providerId!,
        };
        socialLogin(user)
          .then((res) => {
            successApiRequest(res);
          })
          .then(() => {
            getUserProfile().then((res) => {
              handleUserProfile(res?.user);
              router.push("/dashboard");
            });
          })
          .catch((err) => {
            errorApiRequest(err);
          });
      })
      .catch((err) => {
        console.error(err);
      });
  }
};
</script>
