import { createApp } from 'vue'
import App from './App.vue'
import './style.css';
import router from './router/router.js';
import PrimeVue from 'primevue/config';
import Aura from '@primeuix/themes/aura';

createApp(App).use(router).use(PrimeVue, {
    theme: {
        preset: Aura
    }
}).mount('#app')
