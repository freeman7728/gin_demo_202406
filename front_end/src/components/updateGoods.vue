<template>
    <div>
      <el-button type="warning" @click="openDialog">修改商品</el-button>
      <el-dialog v-model="dialogVisible" title="修改商品信息" width="600px" :before-close="handleClose">
        <el-form
          ref="ruleFormRef"
          :model="ruleForm"
          status-icon
          :rules="rules"
          label-width="auto"
          class="demo-ruleForm"
        >
          <el-form-item label="商品编号" prop="id">
            <el-input v-model="ruleForm.id" autocomplete="off" />
          </el-form-item>
          <el-form-item label="商品名称" prop="name">
            <el-input v-model="ruleForm.name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="商品单价" prop="price">
            <el-input v-model="ruleForm.price" autocomplete="off" />
          </el-form-item>
          <el-form-item label="供应商编号" prop="supplierId">
            <el-input v-model="ruleForm.supplierId" autocomplete="off" />
          </el-form-item>
          <el-form-item label="商品简介" prop="description">
            <el-input v-model="ruleForm.description" type="textarea" autocomplete="off" />
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
    price: '',
    supplierId: '',
    description: '',
    remark: '',
  });
  
  const rules = reactive<FormRules<typeof ruleForm>>({
    id: [{ required: true, message: '请输入商品编号', trigger: 'blur' }],
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
        console.log('提交商品信息:', ruleForm);
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
  