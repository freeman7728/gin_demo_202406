<template>
      <el-button type="primary" plain @click="openDialog">添加商品</el-button>
  
      <el-dialog v-model="dialogVisible" title="添加商品信息" width="85%" :before-close="handleClose">
        <el-form ref="productForm" :rules="rules" >
          <!-- 表头 -->
          <el-row class="header-row">
            <el-col :span="3"><strong>商品编号</strong></el-col>
            <el-col :span="3"><strong>商品名称</strong></el-col>
            <el-col :span="3"><strong>商品单价</strong></el-col>
            <el-col :span="3"><strong>客户编号</strong></el-col>
            <el-col :span="3"><strong>商品简介</strong></el-col>
            <el-col :span="3"><strong>备注</strong></el-col>
            <el-col :span="2"><strong>操作</strong></el-col>
          </el-row>
  
          <!-- 输入行 -->
          <div v-for="(product, index) in products" :key="index" class="product-row">
            <el-row :gutter="5">
              <el-col :span="3">
                <el-form-item :prop="'products.' + index + '.id'" >
                  <el-input v-model="product.id"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'products.' + index + '.name'">
                  <el-input v-model="product.name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'products.' + index + '.price'">
                  <el-input type="number" v-model="product.price"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'products.' + index + '.customerId'">
                  <el-input v-model="product.customerId"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'products.' + index + '.description'">
                  <el-input v-model="product.description"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'products.' + index + '.remark'">
                  <el-input type="textarea" v-model="product.remark"></el-input>
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
  import { ElMessageBox } from 'element-plus';
  
  const dialogVisible = ref(false);
  const products = ref([
    {
      id: '',
      name: '',
      price: '',
      customerId: '',
      description: '',
      remark: '',
    },
  ]);
  
  const rules = {
    id: [{ required: true, message: '请输入商品编号', trigger: 'blur' }],
    name: [{ required: true, message: '请输入商品名称', trigger: 'blur' }],
    price: [{ required: true, message: '请输入商品单价', trigger: 'blur' }],
    customerId: [{ required: true, message: '请输入客户编号', trigger: 'blur' }],
    description: [{ required: true, message: '请输入商品简介', trigger: 'blur' }],
  };
  
  const handleClose = (done: () => void) => {
    ElMessageBox.confirm('确定关闭吗？').then(() => {
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
      id: '',
      name: '',
      price: '',
      customerId: '',
      description: '',
      remark: '',
    });
  };
  
  const saveProducts = () => {
    console.log('保存商品列表:', products.value);
    // Reset form or submit data to backend
    resetForm();
    dialogVisible.value = false;
  };
  
  const resetForm = () => {
    products.value = [
      {
        id: '',
        name: '',
        price: '',
        customerId: '',
        description: '',
        remark: '',
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
  