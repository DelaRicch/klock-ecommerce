import {defineStore} from "pinia";
import {reactive, ref} from "vue";
import {accessTokenType} from "@/types";

export const useUserStore = defineStore("userStore", () => {
    const accessToken = reactive<accessTokenType>({
        value: "",
        expiry: 0
    });
    const refreshToken = ref("");
    const userProfile = ref({});

    const setAccessToken = (token: accessTokenType) => {
        accessToken.value = token.value;
        accessToken.expiry = token.expiry;
    }

    const setRefreshToken = (token: string) => {
        refreshToken.value = token;
    }

    return {
        accessToken,
        userProfile,
        refreshToken,
        setAccessToken,
        setRefreshToken
    }
})