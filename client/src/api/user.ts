import axios from "axios";

export const BASE_URL = `${import.meta.env.VITE_APP_BASE_URL}`

const userApi = axios.create({
    baseURL: BASE_URL,
});

export const registerUser = async (data:Record<string, string>) => {
   try {
       return await userApi.post('/register', data)
   } catch (err: any) {
       return err.response.data.message
   }
}
export const loginUser = async (data:Record<string, string>) => {
   try {
       return await userApi.post('/login', data)
   } catch (err: any) {
       return err.response.data.message
   }
}