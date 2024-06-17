  <template>
    <div id="app">
      <router-view />
    </div>
  </template>
  
  <script>
  export default {
    name: "App",
    methods: {
      appclick() {}
    },
    components: {}
  };
  </script>
  
  <style>
  /*
  用作样式切换
  */
  .router-link-active {
    background-color: antiquewhite;
  }
  </style>