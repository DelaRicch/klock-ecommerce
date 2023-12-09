import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";
import { createPinia, PiniaVuePlugin } from "pinia";
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import Vue3Lottie from "vue3-lottie";
import VueApexCharts from "vue3-apexcharts";

const pinia = createPinia();
const app = createApp(App);

router.beforeEach((to, _, next) => {
  document.title = (to?.meta?.title as string) ?? "Klock";
  next();
});

const firebaseConfig = {
  apiKey: import.meta.env.VITE_APP_FIREBASE_API_KEY as string,
  authDomain: import.meta.env.VITE_APP_FIREBASE_AUTH_DOMAIN as string,
  projectId: import.meta.env.VITE_APP_FIREBASE_PROJECT_ID as string,
  storageBucket: import.meta.env.VITE_APP_FIREBASE_STORAGE_BUCKET as string,
  messagingSenderId: import.meta.env
    .VITE_APP_FIREBASE_MESSAGING_SENDER_ID as string,
  appId: import.meta.env.VITE_APP_FIREBASE_APP_ID as string,
  measurementId: import.meta.env.VITE_APP_FIREBASE_MEASUREMENT_ID as string,
};

// Initialize Firebase
const firebaseApp = initializeApp(firebaseConfig);
// getAnalytics(firebaseApp);

app
  .use(router)
  .use(pinia)
  .use(Vue3Lottie)
  .use(PiniaVuePlugin)
  .use(VueApexCharts)
  .mount("#app");
