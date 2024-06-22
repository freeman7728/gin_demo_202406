<template>
    <el-button type="info"  @click="openSearchDialog">查找商品</el-button>
  
    <el-dialog v-model="searchDialogVisible" title="查找商品信息" width="80%" :before-close="handleSearchClose">
      <el-form ref="searchForm" :rules="searchRules">
        <el-row class="header-row">
          <el-col :span="2"><strong>商品编号</strong></el-col>
          <el-col :span="3"><strong>商品名称</strong></el-col>
          <el-col :span="3"><strong>商品单价</strong></el-col>
          <el-col :span="3"><strong>供应商编号</strong></el-col>
          <el-col :span="5"><strong>商品简介</strong></el-col>
          <el-col :span="5"><strong>备注</strong></el-col>
        </el-row>
  
        <el-row :gutter="5" class="search-row">
          <el-col :span="2">
            <el-form-item :prop="'search.productId'">
              <el-input v-model="searchCriteria.productId"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.productName'">
              <el-input v-model="searchCriteria.productName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.unitPrice'">
              <el-input v-model.number="searchCriteria.unitPrice"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item :prop="'search.supplierId'">
              <el-input v-model="searchCriteria.supplierId"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item :prop="'search.description'">
              <el-input type="textarea" v-model="searchCriteria.description"></el-input>
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
          <el-table-column prop="productId" label="商品编号" width="120"></el-table-column>
          <el-table-column prop="productName" label="商品名称" width="150"></el-table-column>
          <el-table-column prop="unitPrice" label="商品单价" width="120"></el-table-column>
          <el-table-column prop="supplierId" label="供应商编号" width="120"></el-table-column>
          <el-table-column prop="description" label="商品简介" width="200"></el-table-column>
          <el-table-column prop="remark" label="备注" width="500"></el-table-column>
        </el-table>
      </template>
      <template v-else>
        <p style="text-align: center; margin-top: 20px">无结果</p>
      </template>
  
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" plain @click="searchProducts">确定</el-button>
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
    productId: '',
    productName: '',
    unitPrice: null,
    supplierId: '',
    description: '',
    remark: '',
  });
  
  const searchResults = ref([]);
  
  const searchRules = {
    // Define rules if needed for validation
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
  
  const searchProducts = () => {
    // Simulate a search by filtering a mock product list
    const mockProducts = [
      { productId: '1', productName: '商品A', unitPrice: 100, supplierId: 'S1', description: '商品A简介', remark: '备注A' },
      { productId: '2', productName: '商品B', unitPrice: 200, supplierId: 'S2', description: '商品B简介', remark: '备注B' },
      // Add more mock products as needed
    ];
  
    searchResults.value = mockProducts.filter(product => {
      return Object.keys(searchCriteria.value).every(key => {
        return searchCriteria.value[key] === '' || searchCriteria.value[key] === null || product[key].toString().includes(searchCriteria.value[key].toString());
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
  