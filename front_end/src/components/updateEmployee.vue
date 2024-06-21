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
import { ElMessageBox, FormInstance, FormRules } from 'element-plus';

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
      console.log('提交员工信息:', ruleForm);
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

