import axios from "axios";
import {socialAuthType} from "@/types";


export const BASE_URL = `${import.meta.env.VITE_APP_BASE_URL}`

const userApi = axios.create({
    baseURL: BASE_URL,
});


export const registerUser = async (data:Record<string, string>) => {
   try {

       const res = await userApi.post('/register', data)
       return res.data
   } catch (err: any) {
       return err.response.data
   }
}

export const loginUser = async (data:Record<string, string>) => {
   try {
       const res = await userApi.post('/login', data)
       return res.data
   } catch (err: any) {
       return err.response.data
   }
}

export const socialLogin = async (data: socialAuthType) => {
    try {
        const res = await userApi.post('/social-login', data)
        return res.data
    } catch (err: any) {
        return err.response.data
    }
}

export const requestNewToken = async () => {
    try {
        const refreshToken = localStorage.getItem("refreshToken");

        const res = await userApi.get('/refresh-token', {
            headers: {
                Authorization: `Bearer ${refreshToken}`,
            },
        });

        return res.data;
    } catch (err) {
        return err.response.data;
    }
};

export const getUserProfile = async () => {
    try {
        const accessToken = localStorage.getItem("accessToken");

        const res = await userApi.get('/user-profile', {
            headers: {
                Authorization: `Bearer ${accessToken}`,
            },
        });

        return res.data;
    } catch (err) {
        return err.response.data;
    }
};