<template>
    <div>
      <el-button type="danger" @click="openDialog">删除客户</el-button>
      <el-dialog v-model="dialogVisible" title="删除客户" width="400px" :before-close="handleClose">
        <el-form
          ref="ruleFormRef"
          :model="ruleForm"
          status-icon
          :rules="rules"
          label-width="auto"
          class="demo-ruleForm"
        >
          <el-form-item label="客户编号" prop="id">
            <el-input v-model="ruleForm.id" autocomplete="off" />
          </el-form-item>
          <el-form-item>
            <el-button type="danger" @click="submitForm(ruleFormRef)">
              删除
            </el-button>
            <el-button @click="resetForm(ruleFormRef)">重置</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </template>
  
  <script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import { ElMessageBox, FormInstance, FormRules, ElMessage } from 'element-plus';
  import axios from 'axios';
  import { getCurrentInstance, } from 'vue';
  const { proxy } = getCurrentInstance();
  
  const dialogVisible = ref(false);
  const ruleFormRef = ref<FormInstance>();
  
  const ruleForm = reactive({
    id: '',
  });
  
  const rules = reactive<FormRules<typeof ruleForm>>({
    id: [{ required: true, message: '请输入客户编号', trigger: 'blur' }],
  });
  
  const handleClose = (done: () => void) => {
    ElMessageBox.confirm('确定关闭吗？').then(() => {
      done();
    }).catch(() => {
      // Handle cancellation
    });
  };
  
  const openDialog = () => {
    dialogVisible.value = true;
  };
  
  const submitForm = async (formEl: FormInstance | undefined) => {
    if (!formEl) return;
    formEl.validate(async (valid) => {
      if (valid) {
        console.log('删除客户编号:', ruleForm.id);
        try {
          const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/producer/delete`, {  id: parseInt(ruleForm.id ) });
          console.log(response);
          if (response.data.code === 200) {
            ElMessage.success('客户已成功删除');
            resetForm(formEl);
            dialogVisible.value = false; // 关闭对话框
            location.reload();
          } else {
            ElMessage.error('删除客户失败');
          }
        } catch (error) {
          ElMessage.error('删除客户时出错');
          console.error(error);
        }
      } else {
        console.log('提交出错!');
      }
    });
  };
  
  const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return;
    formEl.resetFields();
  };
  </script>
  
  <style scoped>
  .demo-ruleForm {
    max-width: 400px;
  }
  </style>
  