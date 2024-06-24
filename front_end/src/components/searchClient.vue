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
            <el-form-item :prop="'search.short_name'">
              <el-input v-model="searchCriteria.short_name"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.address'">
              <el-input v-model="searchCriteria.address"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item :prop="'search.tel'">
              <el-input v-model="searchCriteria.tel"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item :prop="'search.email'">
              <el-input v-model="searchCriteria.email"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item :prop="'search.contact_name'">
              <el-input v-model="searchCriteria.contact_name"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.contact_tel'">
              <el-input v-model="searchCriteria.contact_tel"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item :prop="'search.note'">
              <el-input type="textarea" v-model="searchCriteria.note"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
  
      <template v-if="searchResults.length">
      <el-table :data="searchResults" style="width: 100%; margin-top: 20px;">
        <el-table-column prop="id" label="客户编号" width="120"></el-table-column>
        <el-table-column prop="name" label="客户名称" width="150"></el-table-column>
        <el-table-column prop="short_name" label="客户简称" width="150"></el-table-column>
        <el-table-column prop="address" label="地址" width="150"></el-table-column>
        <el-table-column prop="tel" label="公司电话" width="150"></el-table-column>
        <el-table-column prop="email" label="邮件" width="150"></el-table-column>
        <el-table-column prop="contact_name" label="联系人" width="150"></el-table-column>
        <el-table-column prop="contact_tel" label="联系人电话" width="150"></el-table-column>
        <el-table-column prop="note" label="备注" width="500"></el-table-column>
      </el-table>
    </template>
    <template v-else>
        <p style = "text-align: center; margin-top: 20px"> 无结果 </p>
    </template>
      <template #footer>
        <div class="dialog-footer">
        <el-button type="primary" plain @click="searchCustomers">确定</el-button>
        <el-button @click="exportToCSV">导出 CSV</el-button>
          <el-button @click="searchDialogVisible = false">取消</el-button>
        </div>
      </template>
    </el-dialog>
  </template>
  
  <script setup lang="ts">
  import { onMounted, ref, watch, computed } from 'vue'
  import { useRoute, useRouter} from 'vue-router';
  import axios from 'axios';
  import { ElMessage, ElMessageBox } from 'element-plus';
  import { getCurrentInstance, } from 'vue';
  const { proxy } = getCurrentInstance();
  
  const searchDialogVisible = ref(false);
  const searchCriteria = ref({
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
  const list = ref([{
    id: '',
    name: '',
    short_name: '',
    address: '',
    tel: '',
    email: '',
    contact_name: '',
    contact_tel: '',
    note: '',
  }
  ]);
  
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

  const fetchlist = async () => {
  try {
    const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/producer/select`);
    console.log("Response data:", response.data);
    const rawData = response.data.data;
    if (Array.isArray(rawData)) {
      list.value = rawData.map(item => ({
        id: item.id,
        name: item.name,
        address: item.address,
        short_name: item.short_name,
        tel: item.tel,
        email: item.email,
        contact_name: item.contact_name,
        contact_tel: item.contact_tel,
        note: item.note
      }));
    } else {
      console.error("Unexpected response format:", rawData);
    }
  } catch (error) {
    console.error("获取客户数据失败：", error);
  }
};
  
  const openSearchDialog = () => {
    searchDialogVisible.value = true;
  };
  onMounted(() => {
    fetchlist();
  });
  const searchCustomers = () => {
    // Simulate a search by filtering a mock customer list

  
    searchResults.value = list.value.filter(product => {
      return Object.keys(searchCriteria.value).every(key => {
      return searchCriteria.value[key] === '' ||
             searchCriteria.value[key] === null ||
             (product[key] !== undefined && product[key].toString().includes(searchCriteria.value[key].toString()));
    });
    });
    if (searchResults.value.length === 0) {
      console.log('无结果');
    }
  };
  const exportToCSV = () => {
  try {
    let csvContent = '客户编号,客户名称,客户简称,地址,公司电话,邮件,联系人,联系人电话,备注\n';
    searchResults.value.forEach(customer => {
      csvContent += `${customer.id},${customer.name},${customer.short_name},${customer.address},${customer.tel},${customer.email},${customer.contact_name},${customer.contact_tel},${customer.note}\n`;
    });

    const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
    const url = URL.createObjectURL(blob);

    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', 'customers.csv');
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    ElMessage.success('客户信息已成功导出为 CSV 文件');
  } catch (error) {
    ElMessage.error('导出客户信息失败');
    console.error(error);
  }
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
  