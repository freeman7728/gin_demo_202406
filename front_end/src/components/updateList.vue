<template>
    <div>
      <el-button type="warning" @click="openDialog">修改清单</el-button>
      <el-dialog v-model="dialogVisible" title="修改清单信息" width="600px" :before-close="handleClose">
        <el-form
          ref="ruleFormRef"
          :model="ruleForm"
          status-icon
          :rules="rules"
          label-width="auto"
          class="demo-ruleForm"
        >
          <el-form-item label="清单号" prop="listId">
            <el-input v-model="ruleForm.listId" autocomplete="off" />
          </el-form-item>
          <el-form-item label="员工编号" prop="employeeId">
            <el-input v-model="ruleForm.employeeId" autocomplete="off" />
          </el-form-item>
          <el-form-item label="采购数量" prop="purchaseQuantity">
            <el-input v-model="ruleForm.purchaseQuantity" type="number" autocomplete="off" />
          </el-form-item>
          <el-form-item label="采购总价" prop="purchaseTotal">
            <el-input v-model="ruleForm.purchaseTotal" type="number" autocomplete="off" />
          </el-form-item>
          <el-form-item label="采购事件" prop="purchaseEvent">
            <el-input v-model="ruleForm.purchaseEvent" autocomplete="off" />
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
    listId: '',
    employeeId: '',
    purchaseQuantity: '',
    purchaseTotal: '',
    purchaseEvent: '',
    remark: '',
  });
  
  const rules = reactive<FormRules<typeof ruleForm>>({
    listId: [{ required: true, message: '请输入清单号', trigger: 'blur' }],
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
        console.log('提交清单信息:', ruleForm);
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
  