<template>
      <el-button type="primary"  @click="openDialog">添加订单</el-button>
  
      <el-dialog v-model="dialogVisible" title="添加订单信息" width="85%" :before-close="handleClose">
        <el-form ref="orderForm" :rules="rules" >
          <!-- 表头 -->
          <el-row class="header-row">
            <el-col :span="4"><strong>清单号</strong></el-col>
            <el-col :span="4"><strong>商品编号</strong></el-col>
            <el-col :span="4"><strong>采购数量</strong></el-col>
            <el-col :span="6"><strong>备注</strong></el-col>
            <el-col :span="4"><strong>操作</strong></el-col>
          </el-row>
  
          <!-- 输入行 -->
          <div v-for="(orderItem, index) in list" :key="index" class="order-item-row">
            <el-row :gutter="5">
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.list_id'">
                  <el-input v-model="orderItem.list_id"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.product_id'">
                  <el-input v-model="orderItem.product_id"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.quantity'">
                  <el-input type="number" v-model="orderItem.quantity"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.note'">
                  <el-input type="textarea" v-model="orderItem.note"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="5" class="action-col">
                <el-button type="danger"  @click="removeOrderItem(index)" class="delete-button">删除</el-button>
              </el-col>
            </el-row>
          </div>
  
          <el-button type="primary" plain @click="addNewOrderItem">+ 添加更多订单</el-button>
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
      list_id: '',
      product_id: '',
      quantity: '',
      note: '',
    },
  ]);
  
  const rules = {
    orderId: [{ required: true, message: '请输入订单号', trigger: 'blur' }],
    list_id: [{ required: true, message: '请输入清单号', trigger: 'blur' }],
    product_id: [{ required: true, message: '请输入商品编号', trigger: 'blur' }],
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
    list.value.push({
      list_id: '',
      product_id: '',
      quantity: '',
      note: '',
    });
  };
  
  const savelist = async () => {
    const listToSave = list.value.map(list => ({
      ...list,
      quantity: parseInt(list.quantity),
      product_id: parseFloat(list.product_id),
      list_id: parseFloat(list.list_id)
    }));

    for (let i = 0; i < list.value.length; i ++) { 
      const order = list.value[i];
      console.log(order);
      if (!order.list_id || !order.quantity || !order.product_id) { 
        ElMessage.error('除备注外所有字段必须填写');
        return;
      }
      if (parseInt(order.quantity) <= 0) {
        ElMessage.error('采购数量必须为正数');
        return;
      }
    }
    try {
      const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/detail/insert`, {list: listToSave})
      console.log(response);
      if (response.data.code === 200) {
      console.log(list.value);
      ElMessage.success('订单信息已成功添加至数据库');
      resetForm();
      dialogVisible.value = false;
      location.reload();
    } else {
      console.log(response.status);
      ElMessage.error('添加订单信息失败');
    }
  } catch (error) {
    ElMessage.error('添加订单信息时出错');
    console.error(error);
  }
  };
  
  const resetForm = () => {
    list.value = [
      {
        list_id: '',
        product_id: '',
        quantity: '',
        note: '',
      },
    ];
  };
  
  const removeOrderItem = (index: number) => {
    list.value.splice(index, 1);
  };
  

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
  