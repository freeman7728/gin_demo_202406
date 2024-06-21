<template>
    <el-button type="primary" plain @click="openDialog">添加员工</el-button>

    <el-dialog v-model="dialogVisible" title="添加员工信息" width="90%" :before-close="handleClose">
      <el-form ref="employeeForm" :rules="rules" >
        <el-row class="header-row">
          <el-col :span="4"><strong>员工姓名</strong></el-col>
          <el-col :span="4"><strong>员工密码</strong></el-col>
          <el-col :span="4"><strong>员工电话</strong></el-col>
          <el-col :span="4"><strong>员工工资</strong></el-col>
          <el-col :span="4"><strong>备注</strong></el-col>
          <el-col :span="4"><strong>操作</strong></el-col>
        </el-row>
        <div v-for="(employee, index) in employees" :key="index" class="employee-row">
      <el-row :gutter="5">
        <el-col :span="4">
          <el-form-item :prop="'employees.' + index + '.name'">
            <el-input v-model="employee.name"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="4">
          <el-form-item :prop="'employees.' + index + '.password'">
            <el-input type="password" v-model="employee.password"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="4">
          <el-form-item :prop="'employees.' + index + '.tel'">
            <el-input v-model="employee.tel"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="4">
          <el-form-item :prop="'employees.' + index + '.salary'">
            <el-input v-model="employee.salary"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'employees.' + index + '.note'">
            <el-input type="textarea" v-model="employee.note"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3" class="action-col">
          <el-button type="danger"  @click="removeEmployee(index)" class="delete-button">删除</el-button>
        </el-col>
      </el-row>
    </div>
        

        <el-button type="primary" plain @click="addNewEmployee">+ 添加更多员工</el-button>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveEmployees">确定</el-button>
        </div>
      </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { ElMessageBox } from 'element-plus';
import axios from 'axios';

const dialogVisible = ref(false);
const employees = ref([
  {
    name: '',
    password: '',
    tel: '',
    salary: '',
    note: '',
  },
]);

const rules = {
  id: [{ required: true, message: '请输入员工编号', trigger: 'blur' }],
  name: [{ required: true, message: '请输入员工姓名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入员工密码', trigger: 'blur' }],
  level: [{ required: true, message: '请输入员工级别', trigger: 'blur' }],
  phone: [{ required: true, message: '请输入员工电话', trigger: 'blur' }],
  salary: [{ required: true, message: '请输入员工工资', trigger: 'blur' }],
};

const handleClose = (done: () => void) => {
    ElMessageBox.confirm('确定关闭吗？').then(() => {
      done();
    }).catch(() => {
    });
  };

const openDialog = () => {
  dialogVisible.value = true;
};

const addNewEmployee = () => {
  employees.value.push({
    name: '',
    password: '',
    tel: '',
    salary: '',
    note: '',
  });
};

const saveEmployees = () => {
  console.log('保存员工列表:', employees.value);
  resetForm();
  dialogVisible.value = false;
};

const resetForm = () => {
  employees.value = [
    {
      name: '',
      password: '',
      tel: '',
      salary: '',
      note: '',
    },
  ];
};

const removeEmployee = (index) => {
  employees.value.splice(index, 1);
};
</script>

<style scoped>
.dialog-footer {
  text-align: right;
  margin-top: 10px;
}
.employee-row {
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
