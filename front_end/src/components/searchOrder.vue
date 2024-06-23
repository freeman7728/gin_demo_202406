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
            <el-form-item :prop="'search.id'">
              <el-input v-model="searchCriteria.id"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.quantity'">
              <el-input v-model.number="searchCriteria.quantity"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.price'">
              <el-input v-model.number="searchCriteria.price"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item :prop="'search.total_price'">
              <el-input v-model.number="searchCriteria.total_price"></el-input>
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
          <el-table-column prop="id" label="采购订单号" width="150"></el-table-column>
          <el-table-column prop="list_id" label="采购清单号" width="150"></el-table-column>
          <el-table-column prop="product_id" label="采购商品编号" width="150"></el-table-column>
          <el-table-column prop="quantity" label="采购数量" width="170"></el-table-column>
          <el-table-column prop="price" label="商品单价" width="170"></el-table-column>
          <el-table-column prop="total_price" label="商品总价" width="170"></el-table-column>
          <el-table-column prop="note" label="备注" width="290"></el-table-column>
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
  import { onMounted, ref, watch, computed } from 'vue'
  import { useRoute, useRouter} from 'vue-router';
  import axios from 'axios';
  import { ElMessage, ElMessageBox } from 'element-plus';
  import { getCurrentInstance, } from 'vue';
  const { proxy } = getCurrentInstance();

  
  const searchDialogVisible = ref(false);
  const searchCriteria = ref({
    id: '',
    quantity: null,
    price: null,
    total_price: null,
    note: '',
  });
  const list = ref([{
    id: '',
    quantity: null,
    price: null,
    total_price: null,
    note: '',
  }
  ]);
  const fetchlist = async () => {
  try {
    const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/detail/select`);
    console.log("Response data:", response.data);

    const rawData = response.data.data.list;
    if (Array.isArray(rawData)) {
      list.value = rawData.map(item => ({
        id: item.id,
        list_id: item.list_id,
        price: parseFloat(item.total_price) / parseInt(item.quantity),
        product_id: item.product_id,
        total_price: item.total_price,
        quantity: item.quantity,
        note: item.note,
      }));
    } else {
      console.error("Unexpected response format:", rawData);
    }
  } catch (error) {
    console.error("获取商品数据失败：", error);
  }
};
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
  onMounted(() => {
    fetchlist();
  });
  
  const openSearchDialog = () => {
    searchDialogVisible.value = true;
  };
  
  const searchOrders = () => {
    // Simulate a search by filtering a mock order list

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
  