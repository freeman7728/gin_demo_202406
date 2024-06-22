<template>
    <el-button type="info" @click="openSearchDialog">查找订单</el-button>
  
    <el-dialog v-model="searchDialogVisible" title="查找订单信息" width="80%" :before-close="handleSearchClose">
      <el-form ref="searchForm" :rules="searchRules">
        <el-row class="header-row">
          <el-col :span="4"><strong>采购清单号</strong></el-col>
          <el-col :span="4"><strong>采购数量</strong></el-col>
          <el-col :span="4"><strong>商品单价</strong></el-col>
          <el-col :span="4"><strong>商品总价</strong></el-col>
          <el-col :span="4"><strong>备注</strong></el-col>
        </el-row>
  
        <el-row :gutter="5" class="search-row">
          <el-col :span="4">
            <el-form-item :prop="'search.orderId'">
              <el-input v-model="searchCriteria.orderId"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.quantity'">
              <el-input v-model.number="searchCriteria.quantity"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.unitPrice'">
              <el-input v-model.number="searchCriteria.unitPrice"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.totalPrice'">
              <el-input v-model.number="searchCriteria.totalPrice"></el-input>
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
          <el-table-column prop="orderId" label="采购清单号" width="170"></el-table-column>
          <el-table-column prop="quantity" label="采购数量" width="170"></el-table-column>
          <el-table-column prop="unitPrice" label="商品单价" width="170"></el-table-column>
          <el-table-column prop="totalPrice" label="商品总价" width="170"></el-table-column>
          <el-table-column prop="remark" label="备注" width="490"></el-table-column>
        </el-table>
      </template>
      <template v-else>
        <p style="text-align: center; margin-top: 20px">无结果</p>
      </template>
  
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" plain @click="searchOrders">确定</el-button>
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
    orderId: '',
    quantity: null,
    unitPrice: null,
    totalPrice: null,
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
  
  const searchOrders = () => {
    // Simulate a search by filtering a mock order list
    const mockOrders = [
      { orderId: '001', quantity: 10, unitPrice: 100, totalPrice: 1000, remark: '备注1' },
      { orderId: '002', quantity: 5, unitPrice: 200, totalPrice: 1000, remark: '备注2' },
      // Add more mock orders as needed
    ];
  
    searchResults.value = mockOrders.filter(order => {
      return Object.keys(searchCriteria.value).every(key => {
        return searchCriteria.value[key] === '' || searchCriteria.value[key] === null || order[key].toString().includes(searchCriteria.value[key].toString());
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
  