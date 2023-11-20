export interface signUpFormFields {
    [key: string]: {
        value: string;
        error: string;
    };
}
export interface accessTokenType {
    value: string;
    expiry: number;
}

export interface socialAuthType {
    name: string;
    email: string;
    photo: string;
    socialId: string;
    provider: string;
}

export interface ProductProps {
    id: number;
    product: string;
    order_id: string;
    amount: string;
    status: string;
    customer_name: string;
    date: string;
}