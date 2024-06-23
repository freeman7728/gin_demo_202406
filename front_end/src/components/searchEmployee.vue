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
                <el-option label="root" value="0"></el-option>
                <el-option label="admin" value="1"></el-option>
                <el-option label="ordinary" value="2"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.tel'">
              <el-input v-model="searchCriteria.tel"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.salary'">
              <el-input v-model="searchCriteria.salary"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.note'">
              <el-input type="textarea" v-model="searchCriteria.note"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
  
      
  
      <template v-if="searchResults.length">
      <el-table :data="searchResults" style="width: 100%; margin-top: 20px;">
        <el-table-column prop="id" label="员工编号" width="120"></el-table-column>
        <el-table-column prop="name" label="员工姓名" width="150"></el-table-column>
        <el-table-column prop="level" label="员工级别" width="150">
          <template #default="{ row }">
            {{ getLevelText(row.level) }}
          </template>
        </el-table-column>
        <el-table-column prop="tel" label="员工电话" width="150"></el-table-column>
        <el-table-column prop="salary" label="员工工资" width="150"></el-table-column>
        <el-table-column prop="note" label="备注" width="500"></el-table-column>
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
    level: '',
    tel: '',
    salary: '',
    note: '',
  });
  const list = ref([{
    id: '',
    name: '',
    level: '',
    tel: '',
    salary: '',
    note: '',
  }
  ]);
  
  const searchResults = ref([]);
  
  const searchRules = {
    // Define rules if needed for validation
  };
  const getLevelText = (level) => {
  if (level === 0) {
    return 'root';
  } else if (level === 1) {
    return 'admin';
  } else {
    return 'ordinary';

  }

};
  const fetchlist = async () => {
  try {
    const response = await proxy.$axios.get(`${proxy.$serverUrl_test}/employee/getAll`);
    console.log("Response data:", response.data);

    const rawData = response.data.data.list;
    if (Array.isArray(rawData)) {
      list.value = rawData.map(item => ({
        id: item.id,
        name: item.name,
        account: item.account,
        level: item.level,
        tel: item.tel,
        email: item.tel,
        salary: item.salary,
        note: item.note,
      }));
    } else {
      console.error("Unexpected response format:", rawData);
    }
  } catch (error) {
    console.error("获取员工数据失败：", error);
  }
};  
onMounted(() => {
    fetchlist();
  });
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
    console.log(list.value);
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
  