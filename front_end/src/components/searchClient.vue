<template>
    <el-button type="info"  @click="openSearchDialog">查找客户</el-button>
  
    <el-dialog v-model="searchDialogVisible" title="查找客户信息" width="80%" :before-close="handleSearchClose">
      <el-form ref="searchForm" :rules="searchRules">
        <el-row class="header-row">
          <el-col :span="2"><strong>客户编号</strong></el-col>
          <el-col :span="3"><strong>客户名称</strong></el-col>
          <el-col :span="3"><strong>客户简称</strong></el-col>
          <el-col :span="3"><strong>地址</strong></el-col>
          <el-col :span="2"><strong>公司电话</strong></el-col>
          <el-col :span="2"><strong>邮件</strong></el-col>
          <el-col :span="2"><strong>联系人</strong></el-col>
          <el-col :span="3"><strong>联系人电话</strong></el-col>
          <el-col :span="2"><strong>备注</strong></el-col>
        </el-row>
  
        <el-row :gutter="5" class="search-row">
          <el-col :span="2">
            <el-form-item :prop="'search.id'">
              <el-input v-model="searchCriteria.id"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.name'">
              <el-input v-model="searchCriteria.name"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.shortName'">
              <el-input v-model="searchCriteria.shortName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.address'">
              <el-input v-model="searchCriteria.address"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item :prop="'search.companyPhone'">
              <el-input v-model="searchCriteria.companyPhone"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item :prop="'search.email'">
              <el-input v-model="searchCriteria.email"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item :prop="'search.contactPerson'">
              <el-input v-model="searchCriteria.contactPerson"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.contactPhone'">
              <el-input v-model="searchCriteria.contactPhone"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item :prop="'search.remark'">
              <el-input type="textarea" v-model="searchCriteria.remark"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
  
      <template v-if="searchResults.length">
      <el-table :data="searchResults" style="width: 100%; margin-top: 20px;">
        <el-table-column prop="id" label="客户编号" width="120"></el-table-column>
        <el-table-column prop="name" label="客户名称" width="150"></el-table-column>
        <el-table-column prop="shortName" label="客户简称" width="150"></el-table-column>
        <el-table-column prop="address" label="地址" width="150"></el-table-column>
        <el-table-column prop="companyPhone" label="公司电话" width="150"></el-table-column>
        <el-table-column prop="email" label="邮件" width="150"></el-table-column>
        <el-table-column prop="contactPerson" label="联系人" width="150"></el-table-column>
        <el-table-column prop="contactPhone" label="联系人电话" width="150"></el-table-column>
        <el-table-column prop="remark" label="备注" width="500"></el-table-column>
      </el-table>
    </template>
    <template v-else>
        <p style = "text-align: center; margin-top: 20px"> 无结果 </p>
    </template>
      <template #footer>
        <div class="dialog-footer">
        <el-button type="primary" plain @click="searchCustomers">确定</el-button>
          <el-button @click="searchDialogVisible = false">取消</el-button>
        </div>
      </template>
    </el-dialog>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import { ElMessageBox } from 'element-plus';
  
  const searchDialogVisible = ref(false);
  const searchCriteria = ref({
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
  
  const searchResults = ref([]);
  
  const searchRules = {
    // Define rules if needed for validation
  };
  
  const handleSearchClose = (done: () => void) => {
    ElMessageBox.confirm('确定关闭吗？').then(() => {
      done();
    }).catch(() => {
      // Handle cancellation
    });
  };
  
  const openSearchDialog = () => {
    searchDialogVisible.value = true;
  };
  
  const searchCustomers = () => {
    // Simulate a search by filtering a mock customer list
    const mockCustomers = [
      { id: '1', name: '客户A', shortName: 'A', address: '地址1', companyPhone: '123456', email: 'a@test.com', contactPerson: '联系人A', contactPhone: '1234567890', remark: '备注1' },
      { id: '2', name: '客户B', shortName: 'B', address: '地址2', companyPhone: '234567', email: 'b@test.com', contactPerson: '联系人B', contactPhone: '0987654321', remark: '备注2' },
      // Add more mock customers as needed
    ];
  
    searchResults.value = mockCustomers.filter(customer => {
      return Object.keys(searchCriteria.value).every(key => {
        return searchCriteria.value[key] === '' || customer[key].includes(searchCriteria.value[key]);
      });
    });
  };
  </script>
  
  <style scoped>
  .dialog-footer {
    text-align: right;
    margin-top: 10px;
  }
  .search-row {
    margin-bottom: 20px;
  }
  .header-row {
    margin-bottom: 20px;
    font-weight: bold;
  }
  </style>
  