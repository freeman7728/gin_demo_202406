// types/vue-to-pdf.d.ts
declare module 'vue-to-pdf' {
    import { PluginFunction } from 'vue';
    const VueToPdf: { install: PluginFunction<any> };
    export default VueToPdf;
  }
  