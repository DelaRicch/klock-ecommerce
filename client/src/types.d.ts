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