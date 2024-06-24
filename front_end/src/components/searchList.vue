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
            <el-form-item :prop="'search.employee_id'">
              <el-input v-model="searchCriteria.employee_id"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item :prop="'search.time'">
              <el-date-picker v-model="searchCriteria.time" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item :prop="'search.note'">
              <el-input type="textarea" v-model="searchCriteria.note"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
  
  
      <template v-if="searchResults.length">
        <el-table :data="searchResults" style="width: 100%; margin-top: 20px;">
          <el-table-column prop="id" label="清单号" width="120"></el-table-column>
          <el-table-column prop="employee_id" label="员工编号" width="150"></el-table-column>
          <el-table-column prop="quantity" label="采购数量" width="150"></el-table-column>
          <el-table-column prop="total_price" label="采购总价" width="150"></el-table-column>
          <el-table-column prop="time" label="采购时间" width="200"></el-table-column>
          <el-table-column prop="note" label="备注" width="500"></el-table-column>
        </el-table>
      </template>
      <template v-else>
        <p style="text-align: center; margin-top: 20px;">无结果</p>
      </template>
  
      <template #footer>
        <div class="dialog-footer">
            <el-button type="primary" plain @click="searchInventories">确定</el-button>
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
    employee_id: '',
    time: '',
    note: '',
    quantity: '',
    total_price: ''
  });

  const list = ref([{
    id: '',
    employee_id: '',
    time: '',
    note: '',
    quantity: '',
    total_price: ''
  }
  ]);
  
  const searchResults = ref([]);
  
  const searchRules = {
    // Define rules if needed for validation
  };


  const fetchlist = async () => {
  try {
    const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/order/select`);
    console.log("Response data:", response.data);

    const rawData = response.data.data.list;
    if (Array.isArray(rawData)) {
      list.value = rawData.map(item => ({
        id: item.id,
        employee_id: item.employee_id,
        time: item.time,
        note: item.note,
        total_price: item.total_price,
        quantity: item.quantity
      }));
      console.log(list.value);
    } else {
      console.error("Unexpected response format:", rawData);
    }
  } catch (error) {
    console.error("获取清单数据失败：", error);
  }
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
  onMounted(() => {
    fetchlist();
  });
  const searchInventories = () => {
    // Simulate a search by filtering a mock inventory list
  
    const formatDate = (date) => {
      const d = new Date(date);
      const month = `${d.getMonth() + 1}`.padStart(2, '0');
      const day = `${d.getDate()}`.padStart(2, '0');
      return `${d.getFullYear()}-${month}-${day}`;
    };
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
  const exportToCSV = () => {
  try {
    let csvContent = '清单号,员工编号,采购数量,采购总价,采购时间,备注\n';
    searchResults.value.forEach(purchase => {
      const purchaseTime = new Date(purchase.time).toLocaleString(); 

      csvContent += `${purchase.id},${purchase.employee_id},${purchase.quantity},${purchase.total_price},${purchaseTime},${purchase.note}\n`;
    });

    const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
    const url = URL.createObjectURL(blob);

    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', 'purchase_list.csv');
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    ElMessage.success('清单信息已成功导出为 CSV 文件');
  } catch (error) {
    ElMessage.error('导出清单信息失败');
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
  