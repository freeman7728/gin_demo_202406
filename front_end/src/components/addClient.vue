<template>
      <el-button type="primary"  @click="openDialog">添加客户</el-button>
  
      <el-dialog v-model="dialogVisible" title="添加客户信息" width="80%" :before-close="handleClose">
        <el-form ref="customerForm" :rules="rules" >
          <el-row class="header-row">
            <el-col :span="3"><strong>客户名称</strong></el-col>
            <el-col :span="3"><strong>客户简称</strong></el-col>
            <el-col :span="3"><strong>地址</strong></el-col>
            <el-col :span="2"><strong>公司电话</strong></el-col>
            <el-col :span="2"><strong>邮件</strong></el-col>
            <el-col :span="2"><strong>联系人</strong></el-col>
            <el-col :span="3"><strong>联系人电话</strong></el-col>
            <el-col :span="4"><strong>备注</strong></el-col>
            <el-col :span="2"><strong>操作</strong></el-col>
          </el-row>
  
          <div v-for="(customer, index) in list" :key="index" class="customer-row">
            <el-row :gutter="5">
              <el-col :span="3">
                <el-form-item :prop="'list.' + index + '.name'">
                  <el-input v-model="customer.name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'list.' + index + '.short_name'">
                  <el-input v-model="customer.short_name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'list.' + index + '.address'">
                  <el-input v-model="customer.address"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'list.' + index + '.tel'">
                  <el-input v-model="customer.tel"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'list.' + index + '.email'">
                  <el-input v-model="customer.email"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'list.' + index + '.contact_name'">
                  <el-input v-model="customer.contact_name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'list.' + index + '.contact_tel'">
                  <el-input v-model="customer.contact_tel"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'list.' + index + '.note'">
                  <el-input type="textarea" v-model="customer.note"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3" class="action-col">
                <el-button type="danger"  @click="removeCustomer(index)" class="delete-button">删除</el-button>
              </el-col>
            </el-row>
          </div>
  
          <el-button type="primary" plain @click="addNewCustomer">+ 添加更多客户</el-button>
        </el-form>
  
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="savelist">确定</el-button>
          </div>
        </template>
      </el-dialog>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import { ElMessageBox, ElMessage } from 'element-plus';
  import { getCurrentInstance } from 'vue';
  
  const { proxy } = getCurrentInstance();
  const dialogVisible = ref(false);
  const list = ref([
    {
      name: '',
      short_name: '',
      address: '',
      tel: '',
      email: '',
      contact_name: '',
      contact_tel: '',
      note: '',
    },
  ]);
  
  const rules = {
    name: [{ required: true, message: '请输入客户名称', trigger: 'blur' }],
    short_name: [{ required: true, message: '请输入客户简称', trigger: 'blur' }],
    address: [{ required: true, message: '请输入地址', trigger: 'blur' }],
    tel: [{ required: true, message: '请输入公司电话', trigger: 'blur' }],
    email: [{ required: true, message: '请输入邮件', trigger: 'blur' }],
    contact_name: [{ required: true, message: '请输入联系人', trigger: 'blur' }],
    contact_tel: [{ required: true, message: '请输入联系人电话', trigger: 'blur' }],
  };
  
  const handleClose = (done: () => void) => {
    ElMessageBox.confirm('确定关闭吗？').then(() => {
      resetForm();
      done();
    }).catch(() => {
      // Handle cancellation
    });
  };
  
  const openDialog = () => {
    dialogVisible.value = true;
  };
  
  const addNewCustomer = () => {
    list.value.push({
      name: '',
      short_name: '',
      address: '',
      tel: '',
      email: '',
      contact_name: '',
      contact_tel: '',
      note: '',
    });
  };
  
  const savelist = async () => {
  for (let i = 0; i < list.value.length; i++) {
    const customer = list.value[i];
    if (!customer.name || !customer.short_name || !customer.address || !customer.tel || !customer.email || !customer.contact_name || !customer.tel || !customer.note) {
      ElMessage.error('除备注外所有字段都必须填写');
      return;
    }
  }

  try {
    const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/producer/insert`, { list: list.value });
    if (response.data.code === 200) {
      console.log(list.value);
      ElMessage.success('客户信息已成功添加至数据库');
      resetForm();
      dialogVisible.value = false;
    } else {
      console.log(response.status);
      ElMessage.error('添加客户信息失败');
    }
  } catch (error) {
    ElMessage.error('添加客户信息时出错');
    console.error(error);
  }
};

  
  const resetForm = () => {
    list.value = [
      {
        name: '',
        short_name: '',
        address: '',
        tel: '',
        email: '',
        contact_name: '',
        contact_tel: '',
        note: '',
      },
    ];
  };
  
  const removeCustomer = (index: number) => {
    list.value.splice(index, 1);
  };
  </script>
  
  <style scoped>
  .dialog-footer {
    text-align: right;
    margin-top: 10px;
  }
  .customer-row {
    margin-bottom: 10px;
  }
  .header-row {
    margin-bottom: 20px;
    font-weight: bold;
  }
  .action-col {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .delete-button {
  margin-top: -30px; /* 上移 5px */
  }
  </style>
  