import {createApp} from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import {createPinia, PiniaVuePlugin} from "pinia";

const pinia = createPinia()
const app = createApp(App)

router.beforeEach((to, _, next) => {
    document.title = to?.meta?.title as string ?? 'Klock';
    next();
});


 app.use(router)
     .use(pinia)
.use(PiniaVuePlugin)
     .mount('#app')
