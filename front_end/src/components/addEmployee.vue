<template>
    <el-button type="primary" plain @click="openDialog">添加员工</el-button>

    <el-dialog v-model="dialogVisible" title="添加员工信息" width="85%" :before-close="handleClose">
      <el-form ref="employeeForm" :rules="rules" >
        <el-row class="header-row">
          <el-col :span="3"><strong>员工编号</strong></el-col>
          <el-col :span="3"><strong>员工姓名</strong></el-col>
          <el-col :span="3"><strong>员工密码</strong></el-col>
          <el-col :span="3"><strong>员工级别</strong></el-col>
          <el-col :span="3"><strong>员工电话</strong></el-col>
          <el-col :span="3"><strong>员工工资</strong></el-col>
          <el-col :span="3"><strong>备注</strong></el-col>
          <el-col :span="3"><strong>操作</strong></el-col>
        </el-row>
        <div v-for="(employee, index) in employees" :key="index" class="employee-row">
      <el-row :gutter="5">
        <el-col :span="3">
          <el-form-item :prop="'employees.' + index + '.id'" >
            <el-input v-model="employee.id"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'employees.' + index + '.name'">
            <el-input v-model="employee.name"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'employees.' + index + '.password'">
            <el-input type="password" v-model="employee.password"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'employees.' + index + '.level'">
            <el-input v-model="employee.level"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'employees.' + index + '.phone'">
            <el-input v-model="employee.phone"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'employees.' + index + '.salary'">
            <el-input v-model="employee.salary"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item :prop="'employees.' + index + '.remark'">
            <el-input type="textarea" v-model="employee.remark"></el-input>
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

const dialogVisible = ref(false);
const employees = ref([
  {
    id: '',
    name: '',
    password: '',
    level: '',
    phone: '',
    salary: '',
    remark: '',
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
      // Handle cancellation
    });
  };

const openDialog = () => {
  dialogVisible.value = true;
};

const addNewEmployee = () => {
  employees.value.push({
    id: '',
    name: '',
    password: '',
    level: '',
    phone: '',
    salary: '',
    remark: '',
  });
};

const saveEmployees = () => {
  console.log('保存员工列表:', employees.value);
  // Reset form or submit data to backend
  resetForm();
  dialogVisible.value = false;
};

const resetForm = () => {
  employees.value = [
    {
      id: '',
      name: '',
      password: '',
      level: '',
      phone: '',
      salary: '',
      remark: '',
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
