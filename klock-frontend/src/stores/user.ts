import type { UserProfileProps, accessTokenType } from "@/types";
import { defineStore } from "pinia";
import { reactive, ref } from "vue";

export const useUserStore = defineStore(
  "userStore",
  () => {
    const accessToken = reactive<accessTokenType>({
      value: "",
      expiry: 0,
    });
    const refreshToken = ref("");
    const userProfile = ref<UserProfileProps>({
      Name: "",
      Email: "",
      UserID: "",
      Role: "",
      Photo: "",
      Phone: "",
    });

    const setAccessToken = (token: accessTokenType) => {
      accessToken.value = token.value;
      accessToken.expiry = token.expiry;
    };

    const setUserProfile = (profile: UserProfileProps) => {
      userProfile.value = profile;
    };

    const setRefreshToken = (token: string) => {
      refreshToken.value = token;
    };

    return {
      accessToken,
      userProfile,
      refreshToken,
      setAccessToken,
      setRefreshToken,
      setUserProfile,
    };
  },
  { persist: true },
);
