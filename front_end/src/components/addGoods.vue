<template>
      <el-button type="primary"  @click="openDialog">添加商品</el-button>
  
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
          <div v-for="(product, index) in list" :key="index" class="product-row">
            <el-row :gutter="5">
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.name'">
                  <el-input v-model="product.name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.price'">
                  <el-input type="number" v-model="product.price"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.producer_id'">
                  <el-input v-model="product.producer_id"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.introduction'">
                  <el-input v-model="product.introduction"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'list.' + index + '.note'">
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
            <el-button type="primary" @click="savelist">确定</el-button>
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
  const list = ref([
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
    list.value.push({
      name: '',
      price: '',
      producer_id: '',
      introduction: '',
      note: '',
    });
  };
  
  const savelist = async () => {
    const productToSave = list.value.map(product => ({
    ...product,
    price: parseFloat(product.price),
    producer_id: parseFloat(product.producer_id)
  }));

    for (let i = 0; i < list.value.length; i++) {
      const product = list.value[i];
      if (!product.name || !product.price || !product.producer_id || !product.introduction) {
        ElMessage.error('除备注外所有字段都必须填写');
        return;
      }
    }
  try {
    const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/product/insert`, { list: productToSave });
    console.log(response);
    if (response.data.code === 200) {
      console.log(list.value);
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
    list.value = [
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
    list.value.splice(index, 1);
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
  