import axios from "axios";

export const BASE_URL = `${import.meta.env.VITE_APP_BASE_URL}`

const productApi = axios.create({
    baseURL: BASE_URL,
});

export const addNewProduct = async (data:Record<string, string>) => {
    try {
        const res = await productApi.post('/add-product', data)
        return res.data
    } catch (err: any) {
        return err.response
    }
}

export const getAllProducts = async () => {
    try {
        const res = await productApi.get('/products')
        return res.data
    } catch (err: any) {
        return err.response
    }
}