<template>
      <el-button type="primary" plain @click="openDialog">添加商品</el-button>
  
      <el-dialog v-model="dialogVisible" title="添加商品信息" width="85%" :before-close="handleClose">
        <el-form ref="productForm" :rules="rules" >
          <!-- 表头 -->
          <el-row class="header-row">
            <el-col :span="4"><strong>商品名称</strong></el-col>
            <el-col :span="4"><strong>商品单价</strong></el-col>
            <el-col :span="4"><strong>客户编号</strong></el-col>
            <el-col :span="4"><strong>商品简介</strong></el-col>
            <el-col :span="3"><strong>备注</strong></el-col>
            <el-col :span="2"><strong>操作</strong></el-col>
          </el-row>
  
          <!-- 输入行 -->
          <div v-for="(product, index) in products" :key="index" class="product-row">
            <el-row :gutter="5">
              <el-col :span="4">
                <el-form-item :prop="'products.' + index + '.name'">
                  <el-input v-model="product.name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'products.' + index + '.price'">
                  <el-input type="number" v-model="product.price"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'products.' + index + '.producer_id'">
                  <el-input v-model="product.producer_id"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'products.' + index + '.introduction'">
                  <el-input v-model="product.introduction"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'products.' + index + '.note'">
                  <el-input type="textarea" v-model="product.note"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3" class="action-col">
                <el-button type="danger"  @click="removeProduct(index)" class="delete-button">删除</el-button>
              </el-col>
            </el-row>
          </div>
  
          <el-button type="primary" plain @click="addNewProduct">+ 添加更多商品</el-button>
        </el-form>
  
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="saveProducts">确定</el-button>
          </div>
        </template>
      </el-dialog>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import { ElMessageBox, ElMessage} from 'element-plus';
  import { getCurrentInstance, } from 'vue';
  
  const { proxy } = getCurrentInstance();
  const dialogVisible = ref(false);
  const products = ref([
    {
      name: '',
      price: '',
      producer_id: '',
      introduction: '',
      note: '',
    },
  ]);
  
  const rules = {
    name: [{ required: true, message: '请输入商品名称', trigger: 'blur' }],
    price: [{ required: true, message: '请输入商品单价', trigger: 'blur' }],
    producer_id: [{ required: true, message: '请输入客户编号', trigger: 'blur' }],
    introduction: [{ required: true, message: '请输入商品简介', trigger: 'blur' }],
  };
  
  const handleClose = (done: () => void) => {
    ElMessageBox.confirm('数据未提交，确定关闭吗？').then(() => {
      resetForm();
      done();
    }).catch(() => {
      // Handle cancellation
    });
  };
  
  const openDialog = () => {
    dialogVisible.value = true;
  };
  
  const addNewProduct = () => {
    products.value.push({
      name: '',
      price: '',
      producer_id: '',
      introduction: '',
      note: '',
    });
  };
  
  const saveProducts = async () => {
    for (let i = 0; i < products.value.length; i++) {
      const product = products.value[i];
      if (!product.name || !product.price || !product.producer_id || !product.introduction) {
        ElMessage.error('除备注外所有字段都必须填写');
        return;
      }
    }
  try {
    const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/product/insert`, { products: products.value });
    if (response.status === 200) {
      console.log(products.value);
      ElMessage.success('商品信息已成功添加至数据库');
      resetForm();
      dialogVisible.value = false;
    } else {
      console.log(response.status);
      ElMessage.error('添加商品信息失败');
    }
  } catch (error) {
    ElMessage.error('添加商品信息时出错');
    console.error(error);
  }
};
  
  const resetForm = () => {
    products.value = [
      {
        name: '',
        price: '',
        producer_id: '',
        introduction: '',
        note: '',
      },
    ];
  };
  
  const removeProduct = (index: number) => {
    products.value.splice(index, 1);
  };
  </script>
  
  <style scoped>
  .dialog-footer {
    text-align: right;
    margin-top: 10px;
  }
  .product-row {
    margin-bottom: 10px;
  }
  .header-row {
    margin-bottom: 20px;
    font-weight: bold;
  }
  .action-col {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .delete-button {
  margin-top: -30px; /* 上移 5px */
  }
  </style>
  