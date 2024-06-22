<template>
    <el-button type="info"  @click="openSearchDialog">查找清单</el-button>
  
    <el-dialog v-model="searchDialogVisible" title="查找清单信息" width="80%" :before-close="handleSearchClose">
      <el-form ref="searchForm" :rules="searchRules">
        <el-row class="header-row">
          <el-col :span="4"><strong>清单号</strong></el-col>
          <el-col :span="4"><strong>员工编号</strong></el-col>
          <el-col :span="5"><strong>采购时间</strong></el-col>
          <el-col :span="5"><strong>备注</strong></el-col>
        </el-row>
  
        <el-row :gutter="10"class="search-row">
          <el-col :span="4">
            <el-form-item :prop="'search.id'">
              <el-input v-model="searchCriteria.id"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.employeeId'">
              <el-input v-model="searchCriteria.employeeId"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item :prop="'search.purchaseTime'">
              <el-date-picker v-model="searchCriteria.purchaseTime" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item :prop="'search.remark'">
              <el-input type="textarea" v-model="searchCriteria.remark"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
  
  
      <template v-if="searchResults.length">
        <el-table :data="searchResults" style="width: 100%; margin-top: 20px;">
          <el-table-column prop="id" label="清单号" width="120"></el-table-column>
          <el-table-column prop="employeeId" label="员工编号" width="150"></el-table-column>
          <el-table-column prop="quantity" label="采购数量" width="150"></el-table-column>
          <el-table-column prop="totalPrice" label="采购总价" width="150"></el-table-column>
          <el-table-column prop="purchaseTime" label="采购时间" width="200"></el-table-column>
          <el-table-column prop="remark" label="备注" width="500"></el-table-column>
        </el-table>
      </template>
      <template v-else>
        <p style="text-align: center; margin-top: 20px;">无结果</p>
      </template>
  
      <template #footer>
        <div class="dialog-footer">
            <el-button type="primary" plain @click="searchInventories">确定</el-button>
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
    employeeId: '',
    purchaseTime: '',
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
  
  const searchInventories = () => {
    // Simulate a search by filtering a mock inventory list
    const mockInventories = [
      { id: '1', employeeId: '101', quantity: '50', totalPrice: '5000', purchaseTime: '2024-01-01', remark: '备注1' },
      { id: '2', employeeId: '102', quantity: '30', totalPrice: '3000', purchaseTime: '2024-02-01', remark: '备注2' },
      // Add more mock inventories as needed
    ];
  
    const formatDate = (date) => {
      const d = new Date(date);
      const month = `${d.getMonth() + 1}`.padStart(2, '0');
      const day = `${d.getDate()}`.padStart(2, '0');
      return `${d.getFullYear()}-${month}-${day}`;
    };
  
    searchResults.value = mockInventories.filter(inventory => {
      return Object.keys(searchCriteria.value).every(key => {
        if (key === 'purchaseTime') {
          return searchCriteria.value[key] === '' || formatDate(inventory[key]) === formatDate(searchCriteria.value[key]);
        } else {
          return searchCriteria.value[key] === '' || inventory[key].includes(searchCriteria.value[key]);
        }
      });
    });
  
    if (searchResults.value.length === 0) {
      console.log('无结果');
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
  