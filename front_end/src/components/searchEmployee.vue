<template>
    <el-button type="info"  @click="openSearchDialog">查找员工</el-button>
  
    <el-dialog v-model="searchDialogVisible" title="查找员工信息" width="80%" :before-close="handleSearchClose">
      <el-form ref="searchForm" :rules="searchRules">
        <el-row class="header-row">
          <el-col :span="2"><strong>员工编号</strong></el-col>
          <el-col :span="3"><strong>员工姓名</strong></el-col>
          <el-col :span="3"><strong>员工级别</strong></el-col>
          <el-col :span="4"><strong>员工电话</strong></el-col>
          <el-col :span="3"><strong>员工工资</strong></el-col>
          <el-col :span="4"><strong>备注</strong></el-col>
        </el-row>
  
        <el-row :gutter="10" class="search-row">
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
            <el-form-item :prop="'search.level'">
              <el-select v-model="searchCriteria.level" placeholder="请选择">
                <el-option label="root" value="root"></el-option>
                <el-option label="admin" value="admin"></el-option>
                <el-option label="ordinary" value="ordinary"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.phone'">
              <el-input v-model="searchCriteria.phone"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.salary'">
              <el-input v-model="searchCriteria.salary"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.remark'">
              <el-input type="textarea" v-model="searchCriteria.remark"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
  
      
  
      <template v-if="searchResults.length">
      <el-table :data="searchResults" style="width: 100%; margin-top: 20px;">
        <el-table-column prop="id" label="员工编号" width="120"></el-table-column>
        <el-table-column prop="name" label="员工姓名" width="150"></el-table-column>
        <el-table-column prop="password" label="员工密码" width="150"></el-table-column>
        <el-table-column prop="level" label="员工级别" width="150"></el-table-column>
        <el-table-column prop="phone" label="员工电话" width="150"></el-table-column>
        <el-table-column prop="salary" label="员工工资" width="150"></el-table-column>
        <el-table-column prop="remark" label="备注" width="500"></el-table-column>
      </el-table>
    </template>
    <template v-else>
      <p style="text-align: center; margin-top: 20px;">无结果</p>
    </template>
  
      <template #footer>
        <div class="dialog-footer">
        <el-button type="primary" plain @click="searchEmployees">确定</el-button>
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
    level: '',
    phone: '',
    salary: '',
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
  
  const searchEmployees = () => {
    // Simulate a search by filtering a mock employee list
    const mockEmployees = [
      { id: '1', name: '员工A', password: 'pass123', level: 'root', phone: '123456', salary: '10000', remark: '备注1' },
      { id: '2', name: '员工B', password: 'pass456', level: 'admin', phone: '234567', salary: '8000', remark: '备注2' },
      // Add more mock employees as needed
    ];
  
    searchResults.value = mockEmployees.filter(employee => {
      return Object.keys(searchCriteria.value).every(key => {
        return searchCriteria.value[key] === '' || employee[key].includes(searchCriteria.value[key]);
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
  