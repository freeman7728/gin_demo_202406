<template>
      <el-button type="primary" plain @click="openDialog">添加清单</el-button>
  
      <el-dialog v-model="dialogVisible" title="添加清单信息" width="85%" :before-close="handleClose">
        <el-form ref="listForm" :rules="rules" >
          <!-- 表头 -->
          <el-row class="header-row">
            <el-col :span="3"><strong>清单号</strong></el-col>
            <el-col :span="3"><strong>员工编号</strong></el-col>
            <el-col :span="3"><strong>采购数量</strong></el-col>
            <el-col :span="3"><strong>采购商品编号</strong></el-col>
            <el-col :span="3"><strong>采购事件</strong></el-col>
            <el-col :span="3"><strong>备注</strong></el-col>
            <el-col :span="3"><strong>操作</strong></el-col>
          </el-row>
  
          <!-- 输入行 -->
          <div v-for="(listItem, index) in listItems" :key="index" class="list-item-row">
            <el-row :gutter="5">
              <el-col :span="3">
                <el-form-item :prop="'listItems.' + index + '.id'" >
                  <el-input v-model="listItem.id"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'listItems.' + index + '.employeeId'">
                  <el-input v-model="listItem.employeeId"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'listItems.' + index + '.quantity'">
                  <el-input type="number" v-model="listItem.quantity"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'listItems.' + index + '.productId'">
                  <el-input v-model="listItem.productId"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :prop="'listItems.' + index + '.event'">
                  <el-input v-model="listItem.event"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item :prop="'listItems.' + index + '.remark'">
                  <el-input type="textarea" v-model="listItem.remark"></el-input>
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
            <el-button type="primary" @click="saveListItems">确定</el-button>
          </div>
        </template>
      </el-dialog>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import { ElMessageBox } from 'element-plus';
  
  const dialogVisible = ref(false);
  const listItems = ref([
    {
      id: '',
      employeeId: '',
      quantity: '',
      productId: '',
      event: '',
      remark: '',
    },
  ]);
  
  const rules = {
    id: [{ required: true, message: '请输入清单号', trigger: 'blur' }],
    employeeId: [{ required: true, message: '请输入员工编号', trigger: 'blur' }],
    quantity: [{ required: true, message: '请输入采购数量', trigger: 'blur' }],
    productId: [{ required: true, message: '请输入采购商品编号', trigger: 'blur' }],
    event: [{ required: true, message: '请输入采购事件', trigger: 'blur' }],
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
    listItems.value.push({
      id: '',
      employeeId: '',
      quantity: '',
      productId: '',
      event: '',
      remark: '',
    });
  };
  
  const saveListItems = () => {
    console.log('保存清单列表:', listItems.value);
    // Reset form or submit data to backend
    resetForm();
    dialogVisible.value = false;
  };
  
  const resetForm = () => {
    listItems.value = [
      {
        id: '',
        employeeId: '',
        quantity: '',
        productId: '',
        event: '',
        remark: '',
      },
    ];
  };
  
  const removeListItem = (index: number) => {
    listItems.value.splice(index, 1);
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
  