import {Ref} from "vue";
import {apiResponse, showSnackbar, success} from "../store/resuableState.ts";
import {accessTokenType} from "@/types";
import {useUserStore} from "../store/store.ts";

export const handleSetSnackbarState = (showSnackbar:  Ref<boolean>, duration: number) => {
    showSnackbar.value = true;
    setTimeout(() => {
        showSnackbar.value = false;
    }, duration);
};

export const successApiRequest = (res) => {
    const userStore = useUserStore();

    handleSetSnackbarState(showSnackbar, 6000);
    const accessToken: accessTokenType = {
        value: res?.access_token,
        expiry: res?.exp,
    };
    const refreshToken = res?.refresh_token;
    apiResponse.value = res?.message;
    success.value = !!res?.success;
    console.log(res);

    userStore.setAccessToken(accessToken);
    userStore.setRefreshToken(refreshToken);
}

export const errorApiRequest = (err) => {
    handleSetSnackbarState(showSnackbar, 6000);
    apiResponse.value = err?.message;
    success.value = false;
    console.error(err);
}