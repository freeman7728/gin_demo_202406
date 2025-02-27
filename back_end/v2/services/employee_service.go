package services

import (
	"database_lesson/dao"
	"database_lesson/dto"
	"database_lesson/middleware"
	"database_lesson/models"
	"fmt"
	"net/http"
)

func InsertEmployeeService(employee *dto.EmployeeList) (err error, resp dto.Response) {
	var res dto.EmployeeInsertResult
	resp.Code = http.StatusOK
	for _, model := range employee.EmployeeList {
		err := dao.InsertEmployeeDao(model).Error
		if err != nil {
			resp.Code = 500
			model.Msg = err.Error()
			res.Failed = append(res.Failed, model)
		} else {
			res.Successful = append(res.Successful, model)
		}
	}
	resp.Data = res
	return
}

func LoginEmployeeService(employee *models.Employee) (err error, resp dto.Response) {
	var temp models.Employee
	resp.Code = http.StatusOK
	err = dao.LoginEmployee(employee, &temp).Error
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		return
	}
	if temp.ID == 0 {
		resp.Code = 500
		resp.Message = "账号或者密码错误"
		return
	}
	var res models.Employee
	err = dao.SelectEmployeeByAccount(&temp, &res).Error
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		return
	}
	token, err := middleware.ReleaseToken(res)
	res.Account = employee.Account
	resp.Data = dto.AddToken(&res, token)
	resp.Message = "登录成功"
	return
}

func UpdateEmployeeService(employee *dto.EmployeeUpdateRequest) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	tx := dao.EmployeeUpdate(employee)
	if tx.RowsAffected == 0 {
		resp.Code = 500
		resp.Message = "账号或密码错误"
		return
	}
	if tx.Error != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = tx.Error.Error()
		err = tx.Error
		return
	}
	resp.Message = "修改成功"
	return
}

func DeleteEmployeeService(employee *models.Employee) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	tx := dao.EmployeeDelete(employee)
	if tx.RowsAffected == 0 {
		resp.Code = 500
		resp.Message = "账号或密码错误"
		return
	}
	if tx.Error != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = tx.Error.Error()
		err = tx.Error
		return
	}
	resp.Message = "删除成功"
	return
}
func SelectEmployeeById(employee *models.Employee) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.SelectEmployeeById(employee)
	fmt.Println(employee)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	if res.RowsAffected == 0 {
		resp.Message = "员工记录不存在"
		resp.Code = http.StatusBadRequest
		return
	}
	resp.Data = employee
	resp.Message = "查询成功"
	return
}

func DeleteEmployeeByIdWithoutPassword(model *dto.DeleteEmployeeByAdminRequest) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.SelectEmployeeLevelById(&model.Operator)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	if res.RowsAffected == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "操作者不存在"
		return
	}
	r := dao.CheckOrderDependency(model.Target)
	if r != 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "存在订单依赖目标用户"
		return
	}
	res = dao.DeleteEmployeeByIdWithoutPasswordButCheckLevel(&model.Operator, &model.Target)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	if res.RowsAffected == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "操作对象不存在或权限不足"
		return
	}
	resp.Message = "删除成功"
	return
}

func SelectAllEmployee() (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	var list dto.EmployeeList
	res := dao.SelectAllEmployee(&list)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	resp.Data = list
	resp.Message = "获取成功"
	return
}

func UpdateEmployeeById(employee *models.Employee) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.UpdateEmployeeById(employee)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	resp.Message = "更改成功"
	return
}
