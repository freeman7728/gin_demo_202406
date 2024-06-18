<template>
    <div v-if="show" class="modal-overlay" @click.self="close">
      <div class="modal">
        <div class="modal-header">
          <h3>异常提示</h3>
          <button class="close-button" @click="close">&times;</button>
        </div>
        <div class="modal-body">
          <p>{{ message }}</p>
        </div>
        <div class="modal-footer">
          <button @click="close">关闭</button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'ErrorModal',
    props: {
      message: {
        type: String,
        required: true
      },
      show: {
        type: Boolean,
        required: true
      }
    },
    methods: {
      close() {
        this.$emit('close');
      }
    }
  };
  </script>
  
  <style scoped>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .modal {
    background: white;
    padding: 20px;
    border-radius: 5px;
    width: 300px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }
  
  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .close-button {
    background: none;
    border: none;
    font-size: 1.5em;
    cursor: pointer;
  }
  
  .modal-body {
    margin: 20px 0;
  }
  
  .modal-footer {
    display: flex;
    justify-content: flex-end;
  }
  </style>
  