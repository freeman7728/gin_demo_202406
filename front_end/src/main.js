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
import * as ElementPlusIconsVue from '@element-plus/icons-vue' // ElementPlus-icon

const app = createApp(App);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}


app.use(router);
app.use(Print);
app.use(ElementPlus); //  注册到vue 业务中

app.config.globalProperties.$echarts = echarts;
app.config.globalProperties.axios = axios;
app.config.globalProperties.qs = Qs;
app.config.globalProperties.serverUrl = "http://101.42.232.154:9001";
app.config.globalProperties.serverUrl2 = "http://101.42.232.154:9002";

app.mount('#app');
