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
          <el-button @click="fetchCustomerInfo" style=" margin-left: 80px; margin-bottom: 10px;">确定</el-button>
        
          <el-form-item label="客户名称" prop="name">
            <el-input v-model="ruleForm.name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="客户简称" prop="short_name">
            <el-input v-model="ruleForm.short_name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="地址" prop="address">
            <el-input v-model="ruleForm.address" autocomplete="off" />
          </el-form-item>
          <el-form-item label="公司电话" prop="tel">
            <el-input v-model="ruleForm.tel" autocomplete="off" />
          </el-form-item>
          <el-form-item label="邮件" prop="email">
            <el-input v-model="ruleForm.email" autocomplete="off" />
          </el-form-item>
          <el-form-item label="联系人" prop="contact_name">
            <el-input v-model="ruleForm.contact_name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="联系人电话" prop="contact_tel">
            <el-input v-model="ruleForm.contact_tel" autocomplete="off" />
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
    short_name: '',
    address: '',
    tel: '',
    email: '',
    contact_name: '',
    contact_tel: '',
    note: '',
  });
  
  const fetchCustomerInfo = async () => {
    try {
      const response = await proxy.$axios.get(`${proxy.$serverUrl_test}/customer/${ruleForm.id}`);
      if (response.status === 200) {
        ruleForm.name = response.data.name;
        ruleForm.short_name = response.data.short_name;
        ruleForm.address = response.data.address;
        ruleForm.tel = response.data.tel;
        ruleForm.email = response.data.email;
        ruleForm.contact_name = response.data.contact_name;
        ruleForm.contact_tel = response.data.contact_tel;
        ruleForm.note = response.data.note;
      } else {
        ElMessage.error('获取客户信息失败');
      }
    } catch (error) {
      ElMessage.error('获取客户信息时出错');
      console.error(error);
    }
  };
  const rules = reactive<FormRules<typeof ruleForm>>({
    id: [{ required: true, message: '请输入需要修改的客户的编号', trigger: 'blur' }],
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
        console.log('提交客户信息:', ruleForm);
        try {
          const response = await proxy.$axios.put(`${proxy.$serverUrl_test}/producer/update`, ruleForm);
          if (response.status === 200) {
            ElMessage.success('客户信息已成功更新');
            resetForm(formEl);
            dialogVisible.value = false; // 关闭对话框
          } else {
            ElMessage.error('更新客户信息失败');
          }
        } catch (error) {
          ElMessage.error('更新客户信息时出错');
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
  