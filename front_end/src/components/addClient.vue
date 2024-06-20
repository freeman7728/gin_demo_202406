<template>
      <el-button type="primary" plain @click="openDialog">添加客户</el-button>
  
      <el-dialog v-model="dialogVisible" title="添加客户信息" width="80%" :before-close="handleClose">
        <el-form ref="customerForm" :rules="rules" >
          <el-row class="header-row">
            <el-col :span="2"><strong>客户编号</strong></el-col>
            <el-col :span="3"><strong>客户名称</strong></el-col>
            <el-col :span="3"><strong>客户简称</strong></el-col>
            <el-col :span="3"><strong>地址</strong></el-col>
            <el-col :span="2"><strong>公司电话</strong></el-col>
            <el-col :span="2"><strong>邮件</strong></el-col>
            <el-col :span="2"><strong>联系人</strong></el-col>
            <el-col :span="2"><strong>联系人电话</strong></el-col>
            <el-col :span="2"><strong>备注</strong></el-col>
            <el-col :span="3"><strong>操作</strong></el-col>
          </el-row>
  
          <div v-for="(customer, index) in customers" :key="index" class="customer-row">
            <el-row :gutter="5">
              <el-col :span="2">
                <el-form-item :prop="'customers.' + index + '.id'" >
                  <el-input v-model="customer.id"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'customers.' + index + '.name'">
                  <el-input v-model="customer.name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'customers.' + index + '.shortName'">
                  <el-input v-model="customer.shortName"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'customers.' + index + '.address'">
                  <el-input v-model="customer.address"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'customers.' + index + '.companyPhone'">
                  <el-input v-model="customer.companyPhone"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'customers.' + index + '.email'">
                  <el-input v-model="customer.email"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'customers.' + index + '.contactPerson'">
                  <el-input v-model="customer.contactPerson"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'customers.' + index + '.contactPhone'">
                  <el-input v-model="customer.contactPhone"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'customers.' + index + '.remark'">
                  <el-input type="textarea" v-model="customer.remark"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2" class="action-col">
                <el-button type="danger"  @click="removeCustomer(index)" class="delete-button">删除</el-button>
              </el-col>
            </el-row>
          </div>
  
          <el-button type="primary" plain @click="addNewCustomer">+ 添加更多客户</el-button>
        </el-form>
  
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="saveCustomers">确定</el-button>
          </div>
        </template>
      </el-dialog>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import { ElMessageBox } from 'element-plus';
  
  const dialogVisible = ref(false);
  const customers = ref([
    {
      id: '',
      name: '',
      shortName: '',
      address: '',
      companyPhone: '',
      email: '',
      contactPerson: '',
      contactPhone: '',
      remark: '',
    },
  ]);
  
  const rules = {
    id: [{ required: true, message: '请输入客户编号', trigger: 'blur' }],
    name: [{ required: true, message: '请输入客户名称', trigger: 'blur' }],
    shortName: [{ required: true, message: '请输入客户简称', trigger: 'blur' }],
    address: [{ required: true, message: '请输入地址', trigger: 'blur' }],
    companyPhone: [{ required: true, message: '请输入公司电话', trigger: 'blur' }],
    email: [{ required: true, message: '请输入邮件', trigger: 'blur' }],
    contactPerson: [{ required: true, message: '请输入联系人', trigger: 'blur' }],
    contactPhone: [{ required: true, message: '请输入联系人电话', trigger: 'blur' }],
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
  
  const addNewCustomer = () => {
    customers.value.push({
      id: '',
      name: '',
      shortName: '',
      address: '',
      companyPhone: '',
      email: '',
      contactPerson: '',
      contactPhone: '',
      remark: '',
    });
  };
  
  const saveCustomers = () => {
    console.log('保存客户列表:', customers.value);
    // Reset form or submit data to backend
    resetForm();
    dialogVisible.value = false;
  };
  
  const resetForm = () => {
    customers.value = [
      {
        id: '',
        name: '',
        shortName: '',
        address: '',
        companyPhone: '',
        email: '',
        contactPerson: '',
        contactPhone: '',
        remark: '',
      },
    ];
  };
  
  const removeCustomer = (index: number) => {
    customers.value.splice(index, 1);
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
  