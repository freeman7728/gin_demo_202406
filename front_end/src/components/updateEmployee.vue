<template>
    <div>
      <el-button type="warning" @click="openDialog">修改员工</el-button>
      <el-dialog v-model="dialogVisible" title="修改员工信息" width="600px" :before-close="handleClose">
        <el-form
          ref="ruleFormRef"
          :model="ruleForm"
          status-icon
          :rules="rules"
          label-width="auto"
          class="demo-ruleForm"
        >
        <el-form-item label="员工编号" prop="id">
          <el-input v-model="ruleForm.id" autocomplete="off" />
        </el-form-item>
        <el-button @click="fetchEmployeeInfo" style="margin-left:80px; margin-bottom: 30px;">确定</el-button>
          <el-form-item label="员工姓名" prop="name">
            <el-input v-model="ruleForm.name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="员工密码" prop="password">
            <el-input v-model="ruleForm.password" type="password" autocomplete="off" />
          </el-form-item>
          <el-form-item label="员工级别" prop="level">
            <el-input v-model="ruleForm.level" autocomplete="off" />
          </el-form-item>
          <el-form-item label="员工电话" prop="phone">
            <el-input v-model="ruleForm.phone" autocomplete="off" />
          </el-form-item>
          <el-form-item label="员工工资" prop="salary">
            <el-input v-model="ruleForm.salary" autocomplete="off" />
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
  import { ElMessageBox, FormInstance, FormRules, ElMessage } from 'element-plus';
  import { getCurrentInstance, } from 'vue';
  const { proxy } = getCurrentInstance();

const dialogVisible = ref(false);
const ruleFormRef = ref<FormInstance>();

const ruleForm = reactive({
  id: '',
  name: '',
  password: '',
  level: '',
  phone: '',
  salary: '',
  remark: '',
});

const rules = reactive<FormRules<typeof ruleForm>>({
  id: [{ required: true, message: '请输入员工编号', trigger: 'blur' }],

});
const fetchEmployeeInfo = async () => {
    try {
      const response = await proxy.$axios.get(`${proxy.$serverUrl_test}/employee/${ruleForm.id}`);
      if (response.status === 200) {
        ruleForm.name = response.data.name;
        ruleForm.password = response.data.password;
        ruleForm.level = response.data.level;
        ruleForm.phone = response.data.phone;
        ruleForm.salary = response.data.salary;
        ruleForm.remark = response.data.remark;
      } else {
        ElMessage.error('获取员工信息失败');
      }
    } catch (error) {
      ElMessage.error('获取员工信息时出错');
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
        console.log('提交员工信息:', ruleForm);
        try {
          const response = await proxy.$axios.put(`${proxy.$serverUrl_test}/employee/update`, ruleForm);
          if (response.status === 200) {
            ElMessage.success('员工信息已成功更新');
            resetForm(formEl);
            dialogVisible.value = false; // 关闭对话框
          } else {
            ElMessage.error('更新员工信息失败');
          }
        } catch (error) {
          ElMessage.error('更新员工信息时出错');
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

