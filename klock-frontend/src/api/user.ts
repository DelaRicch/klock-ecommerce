import { useUserStore } from "@/stores/user";
import type { socialAuthType } from "@/types";
import axios from "axios";
import { storeToRefs } from "pinia";

export const BASE_URL = `${import.meta.env.VITE_APP_BASE_URL}`;

const userApi = axios.create({
  baseURL: BASE_URL,
});


export const registerUser = async (data: Record<string, string>) => {
  try {
    const res = await userApi.post("/register", data);
    return res?.data;
  } catch (err: any) {
    return err?.response?.data;
  }
};

export const loginUser = async (data: Record<string, string | boolean>) => {
  try {
    const res = await userApi.post("/login", data);
    return res?.data;
  } catch (err: any) {
    return err?.response?.data;
  }
};

export const socialLogin = async (data: socialAuthType) => {
  try {
    const res = await userApi.post("/social-login", data);
    return res?.data;
  } catch (err: any) {
    return err?.response?.data;
  }
};

export const requestNewToken = async () => {
  try {
    const { refreshToken} = storeToRefs(useUserStore());
    const res = await userApi.get("/refresh-token", {
      headers: {
        Authorization: `Bearer ${refreshToken.value.value}`,
      },
    });

    return res?.data;
  } catch (err: any) {
    return err?.response?.data;
  }
};

export const getUserProfile = async () => {
  try {
const {accessToken} = storeToRefs(useUserStore());
    const res = await userApi.get("/user-profile", {
      headers: {
        Authorization: `Bearer ${accessToken.value.value}`,
      },
    });

    return res?.data;
  } catch (err: any) {
    return err?.response?.data;
  }
};

export const updateUser = async (data: Record<string, string>) => {
  const {accessToken} = storeToRefs(useUserStore());
  try {
    const res = await userApi.patch("/update-user", data, {
      headers: {
      Authorization: `Bearer ${accessToken.value.value}`,
    },
});

    return res?.data;
  } catch (err: any) {
    return err?.response?.data;
  }
}

export const validateCurrentPassword = async (data: string) => {
  const {userProfile} = storeToRefs(useUserStore());
  const request = {
    userId: userProfile.value.UserID,
    password: data
    ,
  }
  try {
    const res = await userApi.post("/validate-current-password", request);
    return res?.data;
  } catch (err: any) {
    return err?.response?.data;
  }
}

export const updatePassword = async (data: Record<string, string>) => {
  const {accessToken} = storeToRefs(useUserStore());
  try {
    const res = await userApi.put("/update-password", data, {
      headers: {
        Authorization: `Bearer ${accessToken.value.value}`,
      },
    });
    return res?.data;
  } catch (err: any) {
    return err?.response?.data;
  }
}

export const updateAvatar = async (data: File) => {
  const {accessToken} = storeToRefs(useUserStore());
  const payload = {
    avatar: data
  }
  try {
    const res = await userApi.put("/update-avatar", payload, {
      headers: {
        Authorization: `Bearer ${accessToken.value.value}`,
      },
    });
    return res?.data;
  } catch (err: any) {
    return err?.response?.data;
  }
}