import {computed, ComputedRef, ref, Ref} from "vue";
import {apiResponse, showSnackbar, showTokenExpiry, success} from "../store/resuableState.ts";
import {accessTokenType} from "@/types";
import {useUserStore} from "../store/store.ts";

export const handleSetSnackbarState = (showSnackbar:  Ref<boolean>, duration: number) => {
    showSnackbar.value = true;
    setTimeout(() => {
        showSnackbar.value = false;
    }, duration);
};

export const successApiRequest = (res: any) => {
    const userStore = useUserStore();

    handleSetSnackbarState(showSnackbar, 6000);
    const accessToken: accessTokenType = {
        value: res?.access_token,
        expiry: res?.exp,
    };
    const refreshToken = res?.refresh_token;
    apiResponse.value = res?.message;
    success.value = !!res?.success;

    userStore.setAccessToken(accessToken);
    userStore.setRefreshToken(refreshToken);

    localStorage.setItem("accessToken", accessToken.value);
    localStorage.setItem("accessTokenExpiry", JSON.stringify(accessToken.expiry));
    localStorage.setItem("refreshToken", refreshToken);
}

export const errorApiRequest = (err: any) => {
    handleSetSnackbarState(showSnackbar, 6000);
    apiResponse.value = err?.message;
    success.value = false;
    console.error(err);
}

export const tokenExpiryCalculator = (tokenExpiration: ComputedRef<number>) => {
    let remainingTime = ref(getRemainingTime());

    function getRemainingTime() {
        return Math.max(0, tokenExpiration.value - Math.floor(Date.now() / 1000));
    }

    function formatTime(seconds: number) {
        const minutes = Math.floor(seconds / 60);
        const remainingSeconds = seconds % 60;
        return `${minutes}:${remainingSeconds < 10 ? '0' : ''}${remainingSeconds}`;
    }

    const formattedExpiration = computed(() => {
        return formatTime(remainingTime.value);
    });

    let countdownInterval = setInterval(() => {
        remainingTime.value = getRemainingTime();
        if (remainingTime.value <= 30 && remainingTime.value > 0) {
            showTokenExpiry.value = true;
        }
        else if (remainingTime.value < 1) {
          showTokenExpiry.value = false;
        }
    }, 1000);

    return {
        formattedExpiration,
        countdownInterval,
        remainingTime
    }
}
