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
        <el-button @click="fetchProductInfo" style="margin-left: 80px; margin-bottom: 30px;">确定</el-button>
          <el-form-item label="商品名称" prop="name">
            <el-input v-model="ruleForm.name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="商品单价" prop="price">
            <el-input v-model="ruleForm.price" autocomplete="off" />
          </el-form-item>
          <el-form-item label="供应商编号" prop="producer_id">
            <el-input v-model="ruleForm.producer_id" autocomplete="off" />
          </el-form-item>
          <el-form-item label="商品简介" prop="introduction">
            <el-input v-model="ruleForm.introduction" type="textarea" autocomplete="off" />
          </el-form-item>
          <el-form-item label="备注" prop="note">
            <el-input v-model="ruleForm.note" type="textarea" autocomplete="off" />
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
  import { ElMessageBox, FormInstance, FormRules, ElMessage } from 'element-plus';
  import { getCurrentInstance, } from 'vue';
  const { proxy } = getCurrentInstance();
  
  const dialogVisible = ref(false);
  const ruleFormRef = ref<FormInstance>();
  
  const ruleForm = reactive({
    id: '',
    name: '',
    price: '',
    producer_id: '',
    introduction: '',
    note: '',
  });
  
  const rules = reactive<FormRules<typeof ruleForm>>({
    id: [{ required: true, message: '请输入商品编号', trigger: 'blur' }],
  });
  const fetchProductInfo = async () => {
    try {
      const response = await proxy.$axios.get(`${proxy.$serverUrl_test}/product/${ruleForm.id}`);
      console.log(response);
      if (response.data.code === 200) {
        ruleForm.name = response.data.data.name;
        ruleForm.price = response.data.data.price;
        ruleForm.producer_id = response.data.data.producer_id;
        ruleForm.introduction = response.data.data.introduction;
        ruleForm.note = response.data.data.note;
        ElMessage.success('请修改商品信息');
      } else {
        ElMessage.error('获取商品信息失败');
      }
    } catch (error) {
      ElMessage.error('获取商品信息时出错');
      console.error(error);
    }
  };

  
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
        console.log('提交商品信息:', ruleForm);
        try {
          const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/product/update`, {
            ...ruleForm,
          id: Number(ruleForm.id) // 确保 id 是整数
          });
          if (response.data.code === 200) {
            ElMessage.success('商品信息已成功更新');
            resetForm(formEl);
            dialogVisible.value = false; // 关闭对话框
          } else {
            ElMessage.error('更新商品信息失败');
          }
        } catch (error) {
          ElMessage.error('更新商品信息时出错');
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
    max-width: 600px;
  }
  </style>
  