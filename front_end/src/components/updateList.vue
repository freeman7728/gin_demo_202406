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
        <el-form-item label="清单号" prop="id">
          <el-input v-model="ruleForm.id" autocomplete="off" />
        </el-form-item>
        <el-button @click="fetchListInfo" style="margin-left: 70px; margin-bottom: 30px;">确定</el-button>
          <el-form-item label="员工编号" prop="employee_id">
            <el-input v-model="ruleForm.employee_id" autocomplete="off" />
          </el-form-item>
          <el-form-item label="采购数量" prop="quantity">
            <el-input v-model="ruleForm.quantity" type="number" autocomplete="off" />
          </el-form-item>
          <el-form-item label="采购总价" prop="total_price">
            <el-input v-model="ruleForm.total_price" type="number" autocomplete="off" />
          </el-form-item>
          <el-form-item label="采购时间" prop="time">
            <el-input v-model="ruleForm.time" autocomplete="off" />
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
    employee_id: '',
    quantity: '',
    total_price: '',
    time: '',
    note: '',
  });
  
  const rules = reactive<FormRules<typeof ruleForm>>({
    id: [{ required: true, message: '请输入清单号', trigger: 'blur' }],
  });

  const fetchListInfo = async () => {
    try {
      const response = await proxy.$axios.get(`${proxy.$serverUrl_test}/order/${ruleForm.id}`);
      if (response.data.code === 200) {
        ruleForm.employee_id = response.data.data.employee_id;
        ruleForm.quantity = response.data.data.quantity;
        ruleForm.total_price = response.data.data.total_price;
        ruleForm.time = response.data.data.time;
        ruleForm.note = response.data.data.note;
      } else {
        ElMessage.error('获取清单信息失败');
      }
    } catch (error) {
      ElMessage.error('获取清单信息时出错');
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
        console.log('提交清单信息:', ruleForm);
        try {
          const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/order/update`, {
            ...ruleForm,
          id: Number(ruleForm.id),
          total_price: parseFloat(ruleForm.total_price)
          });
          console.log(ruleForm);
          console.log(response);
          if (response.data.code === 200) {
            ElMessage.success('清单信息已成功更新');
            resetForm(formEl);
            dialogVisible.value = false; // 关闭对话框
          }
          else if(response.data.code === 404) {
            ElMessage.error('数据未发生更改或数据不存在');
          }
          else {
            ElMessage.error('更新清单信息失败');
          }
        } catch (error) {
          ElMessage.error('员工不存在请检查员工编号');
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
  