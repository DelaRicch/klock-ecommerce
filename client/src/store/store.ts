import { defineStore } from "pinia";
import { onBeforeMount, reactive, ref } from "vue";
import { accessTokenType, UserProfileProps } from "@/types";

export const useUserStore = defineStore("userStore", () => {
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

  // Retrieve token from localStorage if it exists
  const initializeTokenFromLocalStorage = () => {
    const storedToken = localStorage.getItem("accessToken");
    const storedExpiry = localStorage.getItem("accessTokenExpiry");

    if (storedToken !== "undefined" && storedExpiry !== "undefined") {
      accessToken.value = storedToken!;
      accessToken.expiry = Number(storedExpiry);
    }
  };

  // Retrieve user profile from localStorage if it exists
  const initializeUserProfileFromLocalStorage = () => {
    const storedUserProfile = localStorage.getItem("userProfile");
    if (storedUserProfile !== "undefined") {
      userProfile.value = JSON.parse(storedUserProfile as string);
    }
  };

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

  // Automatically retrieve token and user profile from localStorage on store setup
  onBeforeMount(() => {
    initializeTokenFromLocalStorage();
    initializeUserProfileFromLocalStorage();
  });

  return {
    accessToken,
    userProfile,
    refreshToken,
    setAccessToken,
    setRefreshToken,
    setUserProfile,
  };
});
