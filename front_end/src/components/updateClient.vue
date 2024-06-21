<template>
    <div>
      <el-button type="warning" @click="openDialog">修改客户</el-button>
      <el-dialog v-model="dialogVisible" title="修改客户信息" width="600px" :before-close="handleClose">
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
          <el-form-item label="客户名称" prop="name">
            <el-input v-model="ruleForm.name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="客户简称" prop="shortName">
            <el-input v-model="ruleForm.shortName" autocomplete="off" />
          </el-form-item>
          <el-form-item label="地址" prop="address">
            <el-input v-model="ruleForm.address" autocomplete="off" />
          </el-form-item>
          <el-form-item label="公司电话" prop="phone">
            <el-input v-model="ruleForm.phone" autocomplete="off" />
          </el-form-item>
          <el-form-item label="邮件" prop="email">
            <el-input v-model="ruleForm.email" autocomplete="off" />
          </el-form-item>
          <el-form-item label="联系人" prop="contact">
            <el-input v-model="ruleForm.contact" autocomplete="off" />
          </el-form-item>
          <el-form-item label="联系人电话" prop="contactPhone">
            <el-input v-model="ruleForm.contactPhone" autocomplete="off" />
          </el-form-item>
          <el-form-item label="备注" prop="remark">
            <el-input v-model="ruleForm.remark" type="textarea" autocomplete="off" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm(ruleFormRef)">
              修改
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
    id: '',
    name: '',
    shortName: '',
    address: '',
    phone: '',
    email: '',
    contact: '',
    contactPhone: '',
    remark: '',
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
  
  const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return;
    formEl.validate((valid) => {
      if (valid) {
        console.log('提交客户信息:', ruleForm);
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
    max-width: 600px;
  }
  </style>
  