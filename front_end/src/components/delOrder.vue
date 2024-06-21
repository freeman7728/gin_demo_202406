<template>
    <div>
      <el-button type="danger" @click="openDialog">删除订单</el-button>
      <el-dialog v-model="dialogVisible" title="删除订单" width="400px" :before-close="handleClose">
        <el-form
          ref="ruleFormRef"
          :model="ruleForm"
          status-icon
          :rules="rules"
          label-width="auto"
          class="demo-ruleForm"
        >
          <el-form-item label="订单号" prop="orderId">
            <el-input v-model="ruleForm.orderId" autocomplete="off" />
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
  import { ElMessageBox, FormInstance, FormRules } from 'element-plus';
  
  const dialogVisible = ref(false);
  const ruleFormRef = ref<FormInstance>();
  
  const ruleForm = reactive({
    orderId: '',
  });
  
  const rules = reactive<FormRules<typeof ruleForm>>({
    orderId: [{ required: true, message: '请输入订单号', trigger: 'blur' }],
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
  
  const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return;
    formEl.validate((valid) => {
      if (valid) {
        console.log('删除订单号:', ruleForm.orderId);
        // 在这里添加删除订单的逻辑，例如调用API删除订单
        dialogVisible.value = false; // 关闭对话框
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
  