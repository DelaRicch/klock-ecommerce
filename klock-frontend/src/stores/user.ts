import type { UserProfileProps, TokenType } from '@/types';
import { defineStore } from 'pinia';
import { reactive, ref } from 'vue';

export const useUserStore = defineStore(
  'userStore',
  () => {
    const accessToken = reactive<TokenType>({} as TokenType);
    const refreshToken = reactive<TokenType>({} as TokenType);
    const userProfile = ref<UserProfileProps>({
      Name: '',
      Email: '',
      UserID: '',
      Role: '',
      Photo: '',
      Phone: '',
      Gender: '',
      Location: ''
    });

    const setAccessToken = (token: TokenType) => {
      accessToken.expiry = token.expiry;
      accessToken.value = token.value;
    };
    
    const setRefreshToken = (token: TokenType) => {
      refreshToken.expiry = token.expiry;
      refreshToken.value = token.value;
    };

    const setUserProfile = (profile: UserProfileProps) => {
      userProfile.value = profile;
    };

    return {
      accessToken,
      userProfile,
      refreshToken,
      setAccessToken,
      setRefreshToken,
      setUserProfile
    };
  },
  { persist: true }
);
