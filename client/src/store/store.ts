import {defineStore} from "pinia";
import {onBeforeMount, reactive, ref} from "vue";
import {accessTokenType} from "@/types";

export const useUserStore = defineStore("userStore", () => {
    const accessToken = reactive<accessTokenType>({
        value: "",
        expiry: 0
    });
    const refreshToken = ref("");
    const userProfile = ref({});

    // Retrieve token from localStorage if it exists
    const initializeTokenFromLocalStorage = () => {
        const storedToken = localStorage.getItem("accessToken");
        const storedExpiry = localStorage.getItem("accessTokenExpiry");

        if (storedToken && storedExpiry) {
            accessToken.value = storedToken;
            accessToken.expiry = +storedExpiry;
        }
    };

    const setAccessToken = (token: accessTokenType) => {
        accessToken.value = token.value;
        accessToken.expiry = token.expiry;
    }

    const setRefreshToken = (token: string) => {
        refreshToken.value = token;
    }

    // Automatically retrieve token from localStorage on store setup
    onBeforeMount(() => {
        initializeTokenFromLocalStorage();
    });

    return {
        accessToken,
        userProfile,
        refreshToken,
        setAccessToken,
        setRefreshToken
    }
})