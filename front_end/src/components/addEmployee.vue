<template>
    <el-button type="primary"  @click="openDialog">添加员工</el-button>

    <el-dialog v-model="dialogVisible" title="添加员工信息" width="90%" :before-close="handleClose">
      <el-form ref="employeeForm" :rules="rules" >
        <el-row class="header-row">
          <el-col :span="3"><strong>员工姓名</strong></el-col>
          <el-col :span="4"><strong>员工邮箱</strong></el-col>
          <el-col :span="3"><strong>员工电话</strong></el-col>
          <el-col :span="3"><strong>员工等级</strong></el-col>
          <el-col :span="3"><strong>员工工资</strong></el-col>
          <el-col :span="4"><strong>备注</strong></el-col>
          <el-col :span="4"><strong>操作</strong></el-col>
        </el-row>
        <div v-for="(list, index) in list" :key="index" class="employee-row">
      <el-row :gutter="5">
        <el-col :span="3">
          <el-form-item :prop="'list.' + index + '.name'">
            <el-input v-model="list.name"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="4">
          <el-form-item :prop="'list.' + index + '.email'">
            <el-input type="email" v-model="list.email"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'list.' + index + '.tel'">
            <el-input v-model="list.tel"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'list.' + index + '.level'">
            <el-input v-model="list.level"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'list.' + index + '.salary'">
            <el-input v-model="list.salary"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="3">
          <el-form-item :prop="'list.' + index + '.note'">
            <el-input type="textarea" v-model="list.note"></el-input>
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
import { ElMessage } from 'element-plus';
import { getCurrentInstance } from 'vue';
const { proxy } = getCurrentInstance();

const dialogVisible = ref(false);
const list = ref([
  {
    name: '',
    email: '',
    tel: '',
    salary: '',
    level:'',
    note: '',
  },
]);

const rules = {
  id: [{ required: true, message: '请输入员工编号', trigger: 'blur' }],
  name: [{ required: true, message: '请输入员工姓名', trigger: 'blur' }],
  email: [{ required: true, message: '请输入员工密码', trigger: 'blur' }],
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
  list.value.push({
    name: '',
    email: '',
    tel: '',
    salary: '',
    note: '',
    level:''
  });
};

const saveEmployees = async () => {
  const employeesToSave = list.value.map(employee => ({
    ...employee,
    salary: parseFloat(employee.salary),
    level: parseInt(employee.level)
  }));
  try {
    const response = await proxy.$axios.post(`${proxy.$serverUrl_test}/employee/insert`, {
      list: employeesToSave,
    });
    console.log(response); 
    if(response.data.code === 200) {
      console.log('保存员工列表成功:', response.data);
    ElMessageBox.alert('员工信息保存成功', '成功', {
      confirmButtonText: '确定',
    });
    resetForm();
    dialogVisible.value = false;
    location.reload();
    }
  } catch (error) {
    console.error('保存员工列表失败:', error);
    ElMessageBox.alert('员工信息保存失败，请重试', '错误', {
      confirmButtonText: '确定',
    });
  }
};


const resetForm = () => {
  list.value = [
    {
      name: '',
      email: '',
      tel: '',
      salary: '',
      level: '',
      note: '',
    },
  ];
};

const removeEmployee = (index) => {
  list.value.splice(index, 1);
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
