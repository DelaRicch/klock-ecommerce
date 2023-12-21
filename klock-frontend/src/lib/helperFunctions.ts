import {
  apiResponse,
  showSnackbar,
  showTokenExpiry,
  success,
} from "@/stores/resuableState";
import { useUserStore } from "@/stores/user";
import type { accessTokenType, ErrorObject, UserProfileProps } from "@/types";
import { ref, computed, type ComputedRef, type Ref } from "vue";

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

  if(res?.refresh_token || res?.access_token) {
    userStore.setAccessToken(accessToken);
  userStore.setRefreshToken(refreshToken);
  }
};

export const errorApiRequest = (err: any) => {
  handleSetSnackbarState(showSnackbar, 6000);
  apiResponse.value = err?.message;
  success.value = false;
  console.error(err);
};

export const tokenExpiryCalculator = (tokenExpiration: ComputedRef<number>) => {
  
  const getRemainingTime = () => {
    return Math.max(0, tokenExpiration.value - Math.floor(Date.now() / 1000));
  }

  const remainingTime = ref(getRemainingTime());


  function formatTime(seconds: number) {
    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;
    return `${minutes}:${remainingSeconds < 10 ? "0" : ""}${remainingSeconds}`;
  }

  const formattedExpiration = computed(() => {
    return formatTime(remainingTime.value);
  });

  const countdownInterval = setInterval(() => {
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

export const getCurrentLocation = () => {

  const success = (position: { coords: { latitude: number; longitude: number; }; }) => {
      const {latitude, longitude} = position.coords;
      
      const geoApiUrl = `https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=${latitude}&longitude=${longitude}}&localityLanguage=en`;

     fetch(geoApiUrl).then(res => res.json()).then(data => console.log(data)).catch(err => console.log(err));
  }

  const error = () => {
      return "unable to retrieve your location"
  }

  navigator.geolocation.getCurrentPosition(success, error)
}

export const handleLogout = () => {
  const userStore = useUserStore();
  userStore.setUserProfile({} as UserProfileProps);
  userStore.setAccessToken({} as accessTokenType);
  userStore.setRefreshToken("");
  localStorage.removeItem("userProfile");
  localStorage.removeItem("accessToken");
  localStorage.removeItem("accessTokenExpiry");
  localStorage.removeItem("refreshToken");
  window.location.reload();
};
