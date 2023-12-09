import { computed, ComputedRef, ref, Ref } from "vue";
import {
  apiResponse,
  showSnackbar,
  showTokenExpiry,
  success,
} from "../store/resuableState.ts";
import { accessTokenType, ErrorObject, UserProfileProps } from "@/types";
import { useUserStore } from "../store/store.ts";

export const handleSetSnackbarState = (
  showSnackbar: Ref<boolean>,
  duration: number,
) => {
  showSnackbar.value = true;
  setTimeout(() => {
    showSnackbar.value = false;
  }, duration);
};

export const handleUserProfile = (res: UserProfileProps) => {
  const userStore = useUserStore();
  userStore.setUserProfile(res);
  localStorage.setItem("userProfile", JSON.stringify(res));
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
};

export const errorApiRequest = (err: any) => {
  handleSetSnackbarState(showSnackbar, 6000);
  apiResponse.value = err?.message;
  success.value = false;
  console.error(err);
};

export const tokenExpiryCalculator = (tokenExpiration: ComputedRef<number>) => {
  let remainingTime = ref(getRemainingTime());

  function getRemainingTime() {
    return Math.max(0, tokenExpiration.value - Math.floor(Date.now() / 1000));
  }

  function formatTime(seconds: number) {
    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;
    return `${minutes}:${remainingSeconds < 10 ? "0" : ""}${remainingSeconds}`;
  }

  const formattedExpiration = computed(() => {
    return formatTime(remainingTime.value);
  });

  let countdownInterval = setInterval(() => {
    remainingTime.value = getRemainingTime();
    if (remainingTime.value <= 30 && remainingTime.value > 0) {
      showTokenExpiry.value = true;
    } else if (remainingTime.value < 1) {
      showTokenExpiry.value = false;
    }
  }, 1000);

  return {
    formattedExpiration,
    countdownInterval,
    remainingTime,
  };
};

// Image validation
export const isValidImage = (file: File): Promise<File> => {
  return new Promise((resolve, reject) => {
    const allowedTypes: string[] = [
      "image/svg+xml",
      "image/png",
      "image/jpeg",
      "image/jpg",
      "image/gif",
    ];

    if (!allowedTypes.includes(file.type)) {
      const error: ErrorObject = {
        message: "Invalid file type or extension: " + file.name,
      };
      errorApiRequest(error);
      reject(error);
      return;
    }

    if (file.size > 4 * 1024 * 1024) {
      const error: ErrorObject = {
        message: "Image size should be less than 4MB",
      };
      errorApiRequest(error);
      reject(error);
      return;
    }

    const reader = new FileReader();

    reader.onload = (e) => {
      const img = new Image();

      img.onload = () => {
        const dimensions = {
          width: img.width,
          height: img.height,
        };

        // if (dimensions.width > 900 || dimensions.height > 450) {
        if (dimensions.width > 5000 || dimensions.height > 4450) {
          const error: ErrorObject = {
            message: "Image dimensions should be 2000 x 1450px or less",
          };
          errorApiRequest(error);
          reject(error);
        } else {
          resolve(file);
        }
      };

      if (typeof e.target?.result === "string") {
        img.src = e.target?.result;
      }
    };

    reader.readAsDataURL(file);
  });
};

export const handleLogout = () => {
  const userStore = useUserStore();
  userStore.setUserProfile("");
  userStore.setAccessToken("");
  userStore.setRefreshToken("");
  localStorage.removeItem("userProfile");
  localStorage.removeItem("accessToken");
  localStorage.removeItem("accessTokenExpiry");
  localStorage.removeItem("refreshToken");
  window.location.reload();
};
