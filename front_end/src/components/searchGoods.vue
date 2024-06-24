<template>
    <el-button type="info"  @click="openSearchDialog">查找商品</el-button>
  
    <el-dialog v-model="searchDialogVisible" title="查找商品信息" width="80%" :before-close="handleSearchClose">
      <el-form ref="searchForm" :rules="searchRules">
        <el-row class="header-row">
          <el-col :span="2"><strong>商品编号</strong></el-col>
          <el-col :span="3"><strong>商品名称</strong></el-col>
          <el-col :span="3"><strong>供应商编号</strong></el-col>
          <el-col :span="6"><strong>商品简介</strong></el-col>
          <el-col :span="6"><strong>备注</strong></el-col>
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
            <el-form-item :prop="'search.producer_id'">
              <el-input v-model="searchCriteria.producer_id"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item :prop="'search.introduction'">
              <el-input type="textarea" v-model="searchCriteria.introduction"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item :prop="'search.note'">
              <el-input type="textarea" v-model="searchCriteria.note"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
  
      <template v-if="searchResults.length">
        <el-table :data="searchResults" style="width: 100%; margin-top: 20px;">
          <el-table-column prop="id" label="商品编号" width="120"></el-table-column>
          <el-table-column prop="name" label="商品名称" width="150"></el-table-column>
          <el-table-column prop="producer_id" label="供应商编号" width="120"></el-table-column>
          <el-table-column prop="price" label="商品单价" width="120"></el-table-column>
          <el-table-column prop="introduction" label="商品简介" width="200"></el-table-column>
          <el-table-column prop="note" label="备注" width="500"></el-table-column>
        </el-table>
      </template>
      <template v-else>
        <p style="text-align: center; margin-top: 20px">无结果</p>
      </template>
  
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" plain @click="searchProducts">确定</el-button>
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
    producer_id: '',
    introduction: '',
    note: '',
    price: ''
  });
  
  const searchResults = ref([]);
  
  const searchRules = {
    // Define rules if needed for validation
  };


  const list = ref([{
    id: '',
    name: '',
    producer_id: '',
    introduction: '',
    note: '',
    price: ''
  }
  ]);
  const fetchlist = async () => {
  try {
    const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/product/select`);
    console.log("Response data:", response.data);

    const rawData = response.data.data.list;
    if (Array.isArray(rawData)) {
      list.value = rawData.map(item => ({
        id: item.id,
        name: item.name,
        price: parseFloat(item.price),
        producer_id: item.producer_id,
        introduction: item.introduction,
        note: item.note,
      }));
      console.log(123);
      console.log(list.value);
    } else {
      console.error("Unexpected response format:", rawData);
    }
  } catch (error) {
    console.error("获取商品数据失败：", error);
  }
};
  
  const handleSearchClose = (done: () => void) => {
    ElMessageBox.confirm('确定关闭吗？')
      .then(() => {
        done();
      })
      .catch(() => {
        // Handle cancellation
      });
  };
  
  const openSearchDialog = () => {
    searchDialogVisible.value = true;
  };
  onMounted(() => {
    fetchlist();
  });
  
  const searchProducts = () => {
    // Simulate a search by filtering a mock product list
    console.log(list.value);
    searchResults.value = list.value.filter(product => {
    return Object.keys(searchCriteria.value).every(key => {
      // 检查 product[key] 是否存在并且定义，然后再尝试转换为字符串
      return searchCriteria.value[key] === '' ||
             searchCriteria.value[key] === null ||
             (product[key] !== undefined && product[key].toString().includes(searchCriteria.value[key].toString()));
    });
  });
  };
  const exportToCSV = () => {
  try {
    let csvContent = '商品编号,商品名称,商品单价,供应商编号,商品简介,备注\n';
    searchResults.value.forEach(product => {
      csvContent += `${product.id},${product.name},${product.price},${product.producer_id},${product.introduction},${product.note}\n`;
    });

    const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', 'products.csv');
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    ElMessage.success('商品信息已成功导出为 CSV 文件');
  } catch (error) {
    ElMessage.error('导出商品信息失败');
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
  