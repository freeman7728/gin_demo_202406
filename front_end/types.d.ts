// types.d.ts
import { ComponentCustomProperties } from 'vue';
import { AxiosInstance } from 'axios';

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $echarts: typeof echarts;
    $axios: AxiosInstance;
    $qs: typeof Qs;
    $serverUrl_test: string;
  }
}
