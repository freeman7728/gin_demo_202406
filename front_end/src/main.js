import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/css/reset.css'
import '@/assets/css/font.css'
import * as echarts from 'echarts'
import axios from 'axios'
import Qs from 'qs'
import Print from 'vue-print-nb'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue' 

const app = createApp(App);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(router);
app.use(Print);
app.use(ElementPlus); 

app.config.globalProperties.$echarts = echarts;
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$qs = Qs;
app.config.globalProperties.$serverUrl_test = "http://8.146.198.97:4000";

app.mount('#app');
