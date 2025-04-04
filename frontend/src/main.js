import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";
import Alpine from 'alpinejs';

const options = {
    transition: "Vue-Toastification__bounce",
    maxToasts: 20,
    newestOnTop: true
}

const app = createApp(App)
app.use(router)
app.use(store)
app.use(Toast, options)
window.Alpine = Alpine;
Alpine.start();
app.mount('#app')
