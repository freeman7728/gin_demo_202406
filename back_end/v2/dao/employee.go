package dao

import (
	"database_lesson/dto"
	"database_lesson/global"
	"database_lesson/models"
	"errors"
	"gorm.io/gorm"
)

func InsertEmployeeDao(employee models.Employee) (tx *gorm.DB) {
	sql := `CALL CreateEmployeeWithAccount(?, ?, ?, ?, ?, ?,?)`
	if len(employee.Tel) < 6 {
		return &gorm.DB{
			Error: errors.New("电话长度过小"),
		}
	}
	tx = global.DB.Exec(sql, employee.Name, employee.Tel, employee.Tel[len(employee.Tel)-6:], employee.Salary, employee.Note, employee.Level, "@p_account")
	return tx
}

func LoginEmployee(in *models.Employee, out *models.Employee) (tx *gorm.DB) {
	sql := `SELECT id FROM account WHERE account = ? AND password = ?`
	tx = global.DB.Raw(sql, in.Account, in.Password).Scan(out)
	return tx
}

func EmployeeUpdate(in *dto.EmployeeUpdateRequest) (tx *gorm.DB) {
	sql := `UPDATE employee
			LEFT JOIN account
			ON employee.id = account.id
			SET employee.name=?,employee.tel=?,employee.salary=?,employee.note=?,account.password=?
			WHERE account.account=? AND account.password=?`
	result := global.DB.Exec(sql, in.Name, in.Tel, in.Salary, in.Note, in.NewPassword, in.Account, in.Password)
	return result
}

func EmployeeDelete(in *models.Employee) (tx *gorm.DB) {
	//sql := `delete from employee where password = ? and account = ?`
	global.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	sql := `
			DELETE account,employee
			FROM employee
			LEFT JOIN account
			ON employee.id=account.id
			WHERE account.account=? AND account.password=?
			`
	result := global.DB.Exec(sql, in.Account, in.Password)
	global.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
	return result
}

func SelectEmployeeByAccount(model *models.Employee, res *models.Employee) (tx *gorm.DB) {
	sql := `SELECT id,name,tel,salary,note,level FROM employee WHERE id = ?`
	return global.DB.Raw(sql, model.ID).Scan(res)
}
