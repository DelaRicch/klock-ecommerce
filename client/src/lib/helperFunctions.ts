import {Ref} from "vue";

export const handleSetSnackbarState = (showSnackbar:  Ref<boolean>, duration: number) => {
    showSnackbar.value = true;
    setTimeout(() => {
        showSnackbar.value = false;
    }, duration);
};