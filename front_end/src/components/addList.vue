<template>
      <el-button type="primary"  @click="openDialog">添加清单</el-button>
  
      <el-dialog v-model="dialogVisible" title="添加清单信息" width="85%" :before-close="handleClose">
        <el-form ref="listForm" :rules="rules" >
          <!-- 表头 -->
          <el-row class="header-row">
            <el-col :span="4"><strong>员工编号</strong></el-col>
            <el-col :span="4"><strong>采购数量</strong></el-col>
            <el-col :span="4"><strong>采购时间</strong></el-col>
            <el-col :span="4"><strong>采购总价</strong></el-col>
            <el-col :span="4"><strong>备注</strong></el-col>
            <el-col :span="3"><strong>操作</strong></el-col>
          </el-row>
  
          <!-- 输入行 -->
          <div v-for="(listItem, index) in list" :key="index" class="list-item-row">
            <el-row :gutter="5">
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.employee_id'">
                  <el-input v-model="listItem.employee_id"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.quantity'">
                  <el-input type="number" v-model="listItem.quantity"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.time'">
                  <el-input v-model="listItem.time"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="4">
                <el-form-item :prop="'list.' + index + '.total_price'">
                  <el-input v-model="listItem.total_price"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'list.' + index + '.note'">
                  <el-input type="textarea" v-model="listItem.note"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3" class="action-col">
                <el-button type="danger"  @click="removeListItem(index)" class="delete-button">删除</el-button>
              </el-col>
            </el-row>
          </div>
  
          <el-button type="primary" plain @click="addNewListItem">+ 添加更多清单</el-button>
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
      employee_id: '',
      quantity: '',
      time: '',
      note: '',
      total_price: ''
    },
  ]);
  
  const rules = {
    id: [{ required: true, message: '请输入清单号', trigger: 'blur' }],
    employee_id: [{ required: true, message: '请输入员工编号', trigger: 'blur' }],
    productId: [{ required: true, message: '请输入采购商品编号', trigger: 'blur' }],
    time: [{ required: true, message: '请输入采购时间', trigger: 'blur' }],
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
  
  const addNewListItem = () => {
    list.value.push({
      employee_id: '',
      quantity: '',
      time: '',
      note: '',
      total_price: ''
    });
  };
  
  const savelist = async () => {
    const listToSave = list.value.map(list => ({
      ...list,
      quantity: parseInt(list.quantity),
      employee_id: parseFloat(list.employee_id),
      total_price: parseFloat(list.total_price)
    }));

    for (let i = 0; i < list.value.length; i ++) { 
      const order = list.value[i];
      if (!order.employee_id || !order.quantity || !order.time || !order.total_price) { 
        ElMessage.error('除备注外所有字段必须填写');
        return;
      }
    }
    try {
      const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/order/insert`, {list: listToSave})
      console.log(response);
      if (response.data.code === 200) {
      console.log(list.value);
      ElMessage.success('清单信息已成功添加至数据库');
      resetForm();
      dialogVisible.value = false;
    } else {
      console.log(response.status);
      ElMessage.error('添加清单信息失败');
    }
  } catch (error) {
    ElMessage.error('添加清单信息时出错');
    console.error(error);
  }
  };
  
  const resetForm = () => {
    list.value = [
      {
        id: '',
        employee_id: '',
        quantity: '',
        productId: '',
        time: '',
        note: '',
      },
    ];
  };
  
  const removeListItem = (index: number) => {
    list.value.splice(index, 1);
  };
  </script>
  
  <style scoped>
  .dialog-footer {
    text-align: right;
    margin-top: 10px;
  }
  .list-item-row {
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
  