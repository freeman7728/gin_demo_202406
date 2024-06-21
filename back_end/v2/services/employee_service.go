package services

import (
	"database_lesson/dao"
	"database_lesson/dto"
	"database_lesson/models"
	"net/http"
)

func InsertEmployeeService(employee *dto.EmployeeList) (err error, resp dto.Response) {
	var res dto.EmployeeInsertResult
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
	var res models.Employee
	resp.Code = http.StatusOK
	err = dao.LoginEmployee(employee, &res).Error
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
	}
	if res.ID == 0 {
		resp.Message = "账号或者密码错误"
		return
	}
	resp.Data = res
	resp.Message = "登录成功"
	return
}

func UpdateEmployeeService(employee *dto.EmployeeUpdateRequest) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	tx := dao.EmployeeUpdate(employee)
	if tx.RowsAffected == 0 {
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
