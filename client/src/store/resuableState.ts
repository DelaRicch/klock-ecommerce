import {ref} from "vue";

// Snackbar
export const success = ref(false);
export const showSnackbar = ref(false);
export const apiResponse = ref('');

// Set colors based on screen size
export const isDesktop = ref(false);
export const textColor = ref("#1D2939");
export const innerColor = ref("black");
export const outerColor = ref("white");

// Show token expiry notification
export const showTokenExpiry = ref(false);

// Button value state
export const buttonValue = ref(false);

// Pagination current page
