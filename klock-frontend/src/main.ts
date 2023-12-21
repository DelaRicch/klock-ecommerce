import "./style.css";
import { createApp } from 'vue'
import { PiniaVuePlugin, createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import { initializeApp } from "firebase/app";
import Vue3Lottie from "vue3-lottie";
import VueApexCharts from "vue3-apexcharts";

import App from './App.vue'
import router from './router'

const app = createApp(App)
const pinia = createPinia()

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

  initializeApp(firebaseConfig);

pinia.use(piniaPluginPersistedstate)
app.use(pinia)
.use(router)
.use(Vue3Lottie)
.use(PiniaVuePlugin)
.use(VueApexCharts)
app.mount('#app')
