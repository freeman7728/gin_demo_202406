<template>
    <div id="app" class="container" ref="appContainer" >
      <div class="container__form signin-admin" ref="container1">
        <el-form :model="adminForm" ref="adminForm_ref" class="form" @submit.native.prevent="submitForm('admin')">
          <h2 class="form-title">管理员登录</h2>
          <el-form-item prop="adminAccount">
            <el-input v-model="adminForm.adminAccount" placeholder="管理员账号" class="input" />
          </el-form-item>
          <el-form-item prop="adminPassword">
            <el-input v-model="adminForm.adminPassword" type="password" placeholder="密码" class="input" />
          </el-form-item>
          <a href="#" class="link">忘记密码？</a>
          <el-button type="primary" native-type="submit" class="btn">登录</el-button>
        </el-form>
      </div>
  
      <div class="container__form signin-ordinary" ref="container2">
        <el-form :model="ordinaryForm" ref="ordinaryForm_ref" class="form" @submit.native.prevent="submitForm('ordinary')">
          <h2 class="form-title">普通用户登录</h2>
          <el-form-item prop="ordinaryAccount">
            <el-input v-model="ordinaryForm.ordinaryAccount" placeholder="账号" class="input" />
          </el-form-item>
          <el-form-item prop="ordinaryPassword">
            <el-input v-model="ordinaryForm.ordinaryPassword" type="password" placeholder="密码" class="input" />
          </el-form-item>
          <a href="#" class="link">忘记密码？</a>
          <el-button type="primary" native-type="submit" class="btn">登录</el-button>
        </el-form>

      </div>
  
      <div class="container__overlay">
        <div class="overlay">
          <div class="overlay__panel overlay--left">
            <el-button type="primary" @click="switchToSignIn">普通用户登录</el-button>
          </div>
          <div class="overlay__panel overlay--right">
            <el-button type="primary" @click="switchToSignUp">管理员登录</el-button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <style scoped>
  :root {
  --white: #e9e9e9;
  --gray: #333;
  --blue: #0367a6;
  --lightblue: #008997;
  --input-bg: #f5f7fa; 
  --input-border: #dcdfe6; 
  --input-focus-border: #409eff; 


  --button-radius: 0.7rem;
  --input-radius: 3px; 


  --max-width: 758px;
  --max-height: 420px;

  font-size: 16px;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
    Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
}

.container {
  align-items: center;
  background: url('@/assets/img/background.jpg') no-repeat center center fixed; 
  background-size: cover;
  display: grid;
  height: 100vh;
  place-items: center;
}

.form-title {
  font-family: kai1;
  font-weight: 300;
  margin: 0;
  margin-bottom: 1.25rem;
  align-self: center;
}

.link {
  color: var(--gray);
  font-size: 0.9rem;
  margin: 1.5rem 0;
  text-decoration: none;
}

.container__form {
  background-color: rgb(202, 194, 194);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); 
  border-radius: 8px;
  height: 50%;
  width: 30%; 
  position: absolute;
  top: 170px;
  transition: all 0.6s ease-in-out;
}

.signin-ordinary {
  z-index: 1;
  left: 20%; 
}

.right-panel-active .signin-ordinary {
  transform: translateX(100%);
}

.signin-admin {
  left: 20%; 
  opacity: 0; 
  z-index: 1;
}

.right-panel-active .signin-admin {
  animation: show 0.6s;
  opacity: 1;
  transform: translateX(100%);
  z-index: 5;
}

.container__overlay {
  position: absolute;
  height: 50%;
  left: 50%;
  overflow: hidden;
  top: 170px;
  transition: transform 0.6s ease-in-out;
  border-radius: 8px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); 
  width: 30%;
  z-index: 100;
}

.right-panel-active .container__overlay {
  transform: translateX(-100%);
}

.overlay {
  background-color: var(--lightblue);
  background: url('@/assets/img/background.jpg');
  background-attachment: fixed;
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
  height: 100%;
  left: -100%;
  position: relative;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
  width: 200%;
}

.right-panel-active .overlay {
  transform: translateX(50%);
}

.overlay__panel {
  align-items: center;
  display: flex;
  flex-direction: column;
  height: 100%;
  justify-content: center;
  position: absolute;
  text-align: center;
  top: 0;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
  width: 50%;
}

.overlay--left {
  transform: translateX(0);
}

.container.right-panel-active .overlay--left {
  transform: translateX(0);
}

.overlay--right {
  right: 0;
  transform: translateX(0);
}

.container.right-panel-active .overlay--right {
  transform: translateX(20%);
}

.btn {
  background-color: rgb(132, 213, 236);
  background-image: linear-gradient(90deg, var(--blue) 0%, var(--lightblue) 74%);
  border-radius: 15px;
  border: 1px solid var(--blue);  
  color: var(--white);
  cursor: pointer;
  font-size: 0.8rem;
  font-weight: bold;
  font-family: alimama;
  letter-spacing: 0.1rem;
  padding: 0.9rem 4rem;
  text-transform: uppercase;
  transition: transform 80ms ease-in;
}

.form > .btn {
  margin-top: 1.5rem;
}

.btn:active {
  transform: scale(0.95);
}

.btn:focus {
  outline: none;
}

.form {
  background-color: var(--white);
  display: flex;
  align-items: center;
  border-radius: 8px;
  justify-content: center;
  flex-direction: column;
  padding: 0 3rem;
  height: 100%;
  text-align: center;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); 
}

.input .el-input__inner {
  background-color: var(--input-bg);
  border: 1px solid var(--input-border);
  border-radius: var(--input-radius);
  padding: 0.9rem 0.9rem;
  margin: 0.5rem 0;
  width: 100%;
  transition: border-color 0.3s;
}

.input .el-input__inner:focus {
  border-color: var(--input-focus-border);
}

@keyframes show {
  0%,
  49.99% {
    opacity: 0;
    z-index: 1;
  }

  50%,
  100% {
    opacity: 1;
    z-index: 5;
  }
}
  </style>
  
  <script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';
import { ElForm, ElFormItem, ElInput, ElButton, ElMessage } from 'element-plus';
import 'element-plus/dist/index.css';
import { ElMessageBox,} from 'element-plus';
import { getCurrentInstance } from 'vue';
  
const { proxy } = getCurrentInstance();
const dialogVisible = ref(false);

const appContainer = ref<HTMLElement | null>(null);
const adminForm = ref({
  adminAccount: '',
  adminPassword: ''
});
const ordinaryForm = ref({
  ordinaryAccount: '',
  ordinaryPassword: ''
});

const router = useRouter();


// const ordinaryLogin = () => {
//   axios.post(`${serverUrl.value}/ordinary/login`, {
//     ordinaryId: ordinaryForm.value.ordinaryAccount,
//     ordinaryPassword: ordinaryForm.value.ordinaryPassword,
//   }, {
//     headers: {
//       'Content-Type': 'application/json',
//     },
//   })
//     .then(res => {
//       console.log(res.data);
//       if (res.data.code === 200) {
//         const result = res.data.data;
//         sessionStorage.setItem('ordinaryId', result.ordinaryId);
//         sessionStorage.setItem('ordinaryInfo', JSON.stringify(result));
//         router.replace('/ordinaryindex');
//       } else {
//         ElMessage.error('账号或密码错误！');
//       }
//     })
//     .catch(err => {
//       console.error(err);
//     });
// };


const adminLogin = async () => {
   try{
    const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/employee/login`, {
    account: adminForm.value.adminAccount,
    password: adminForm.value.adminPassword,
  })
    console.log(response);
    if (response.data.code === 200) {
        router.replace('/index_admin');
    } else {
      ElMessage.error('账号或密码错误！');
    }
   } catch (error) {
    ElMessage.error('登录失败');
    console.error(error);
   }
  
    // .then(res => {
    //   console.log(res.data);
    //   if (res.data.code === 200) {
    //     const result = res.data;
    //     sessionStorage.setItem('adminId', result.adminId);
    //     sessionStorage.setItem('adminInfo', JSON.stringify(result));
    //     router.replace('/childpages/user_info');
    //   } else {
    //     ElMessage.error('账号或密码错误！');
    //   }
    // })
    // .catch(err => {
    //   console.error(err);
    // });
  
};

const switchToSignIn = () => {
  console.log("切换到登录");
  if (appContainer.value) {
    appContainer.value.classList.remove("right-panel-active");
  }
};


const switchToSignUp = () => {
  console.log("切换到注册");
  if (appContainer.value) {
    appContainer.value.classList.add("right-panel-active");
  }
};


const submitForm = (formType: 'admin' | 'ordinary') => {
  if (formType === 'admin') {
    adminLogin();
  } else if (formType === 'ordinary') {
    ordinaryLogin();
  }
};
</script>
  