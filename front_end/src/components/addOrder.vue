<template>
      <el-button type="primary" plain @click="openDialog">添加订单</el-button>
  
      <el-dialog v-model="dialogVisible" title="添加订单信息" width="85%" :before-close="handleClose">
        <el-form ref="orderForm" :rules="rules" >
          <!-- 表头 -->
          <el-row class="header-row">
            <el-col :span="3"><strong>订单号</strong></el-col>
            <el-col :span="3"><strong>清单号</strong></el-col>
            <el-col :span="3"><strong>商品编号</strong></el-col>
            <el-col :span="3"><strong>采购数量</strong></el-col>
            <el-col :span="3"><strong>商品单价</strong></el-col>
            <el-col :span="3"><strong>商品总价</strong></el-col>
            <el-col :span="3"><strong>备注</strong></el-col>
            <el-col :span="3"><strong>操作</strong></el-col>
          </el-row>
  
          <!-- 输入行 -->
          <div v-for="(orderItem, index) in orderItems" :key="index" class="order-item-row">
            <el-row :gutter="5">
              <el-col :span="3">
                <el-form-item :prop="'orderItems.' + index + '.orderId'" >
                  <el-input v-model="orderItem.orderId"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'orderItems.' + index + '.listId'">
                  <el-input v-model="orderItem.listId"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'orderItems.' + index + '.productId'">
                  <el-input v-model="orderItem.productId"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'orderItems.' + index + '.quantity'">
                  <el-input type="number" v-model="orderItem.quantity"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'orderItems.' + index + '.unitPrice'">
                  <el-input v-model="orderItem.unitPrice"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'orderItems.' + index + '.totalPrice'">
                  <el-input v-model="orderItem.totalPrice" disabled></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'orderItems.' + index + '.remark'">
                  <el-input type="textarea" v-model="orderItem.remark"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3" class="action-col">
                <el-button type="danger"  @click="removeOrderItem(index)" class="delete-button">删除</el-button>
              </el-col>
            </el-row>
          </div>
  
          <el-button type="primary" plain @click="addNewOrderItem">+ 添加更多订单</el-button>
        </el-form>
  
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="saveOrderItems">确定</el-button>
          </div>
        </template>
      </el-dialog>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import { ElMessageBox } from 'element-plus';
  
  const dialogVisible = ref(false);
  const orderItems = ref([
    {
      orderId: '',
      listId: '',
      productId: '',
      quantity: '',
      unitPrice: '',
      totalPrice: '',
      remark: '',
    },
  ]);
  
  const rules = {
    orderId: [{ required: true, message: '请输入订单号', trigger: 'blur' }],
    listId: [{ required: true, message: '请输入清单号', trigger: 'blur' }],
    productId: [{ required: true, message: '请输入商品编号', trigger: 'blur' }],
    quantity: [{ required: true, message: '请输入采购数量', trigger: 'blur' }],
    unitPrice: [{ required: true, message: '请输入商品单价', trigger: 'blur' }],
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
  
  const addNewOrderItem = () => {
    orderItems.value.push({
      orderId: '',
      listId: '',
      productId: '',
      quantity: '',
      unitPrice: '',
      totalPrice: '',
      remark: '',
    });
  };
  
  const saveOrderItems = () => {
    console.log('保存订单列表:', orderItems.value);
    // Reset form or submit data to backend
    resetForm();
    dialogVisible.value = false;
  };
  
  const resetForm = () => {
    orderItems.value = [
      {
        orderId: '',
        listId: '',
        productId: '',
        quantity: '',
        unitPrice: '',
        totalPrice: '',
        remark: '',
      },
    ];
  };
  
  const removeOrderItem = (index: number) => {
    orderItems.value.splice(index, 1);
  };
  
  // 计算商品总价
  const computeTotalPrice = (item: any) => {
    if (item.quantity && item.unitPrice) {
      item.totalPrice = parseFloat(item.quantity) * parseFloat(item.unitPrice);
    } else {
      item.totalPrice = '';
    }
  };
  
  // 监听数量和单价变化，更新总价

  </script>
  
  <style scoped>
  .dialog-footer {
    text-align: right;
    margin-top: 10px;
  }
  .order-item-row {
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
  