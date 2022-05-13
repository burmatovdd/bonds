import { createApp } from 'vue'
import App from './App.vue'
import router from "./components/login/router";


createApp(App)
    .use(router)
    .mount('#app')
