<template>
    <div>
      <el-button type="warning" @click="openDialog">修改订单</el-button>
      <el-dialog v-model="dialogVisible" title="修改订单信息" width="600px" :before-close="handleClose">
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
          <el-form-item label="清单号" prop="listId">
            <el-input v-model="ruleForm.listId" autocomplete="off" />
          </el-form-item>
          <el-form-item label="商品号" prop="productId">
            <el-input v-model="ruleForm.productId" autocomplete="off" />
          </el-form-item>
          <el-form-item label="采购数量" prop="purchaseQuantity">
            <el-input v-model="ruleForm.purchaseQuantity" type="number" autocomplete="off" />
          </el-form-item>
          <el-form-item label="商品单价" prop="productPrice">
            <el-input v-model="ruleForm.productPrice" type="number" autocomplete="off" />
          </el-form-item>
          <el-form-item label="商品总价" prop="totalPrice">
            <el-input v-model="ruleForm.totalPrice" type="number" autocomplete="off" />
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
    orderId: '',
    listId: '',
    productId: '',
    purchaseQuantity: '',
    productPrice: '',
    totalPrice: '',
    remark: '',
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
        console.log('提交订单信息:', ruleForm);
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
  