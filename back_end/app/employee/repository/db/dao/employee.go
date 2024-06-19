package dao

import (
	"context"
	"database_lesson/app/employee/repository/db/model"
	"database_lesson/idl/pb"
	"errors"
	"gorm.io/gorm"
)

type EmployeeDao struct {
	*gorm.DB
}

func NewEmployeeDao(ctx context.Context) *EmployeeDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &EmployeeDao{NewDBClient(ctx)}
}

func (dao *EmployeeDao) CreateEmployee(in *model.Employee) (tx *gorm.DB) {
	sql := `CALL CreateEmployeeWithAccount(?, ?, ?, ?, ?, ?,?)`
	if len(in.Tel) < 6 {
		return &gorm.DB{
			Error: errors.New("电话长度过小"),
		}
	}
	tx = dao.Exec(sql, in.Name, in.Tel, in.Tel[len(in.Tel)-6:], in.Salary, in.Note, in.Level, "@p_account")
	return tx
}
func (dao *EmployeeDao) LoginEmployee(in *model.Login, out *model.Login) (tx *gorm.DB) {
	sql := `SELECT id FROM account WHERE account = ? AND password = ?`
	result := dao.Raw(sql, in.Account, in.Password).Scan(out)
	return result
}
func (dao *EmployeeDao) GetEmployeeByTel(successModel *pb.EmployeeSignupSuccessModel) (tx *gorm.DB) {
	sql := `SELECT employee.id,account,password,name,tel,salary,note,level
			FROM employee
			LEFT JOIN account
			ON employee.id = account.id
			WHERE employee.tel = ?`
	result := dao.Raw(sql, successModel.Tel).Scan(&successModel)
	return result
}

func (dao *EmployeeDao) EmployeeUpdate(in *pb.EmployeeUpdateRequest) (tx *gorm.DB) {
	//sql := `update employee set name = ?, tel = ?, level = ?,password = ? where account = ? and password = ?`
	sql := `UPDATE employee
			LEFT JOIN account
			ON employee.id = account.id
			SET employee.name=?,employee.tel=?,employee.salary=?,employee.note=?,account.password=?
			WHERE account.account=? AND account.password=?`
	result := dao.Exec(sql, in.Name, in.Tel, in.Salary, in.Note, in.NewPassword, in.Account, in.Password)
	return result
}

func (dao *EmployeeDao) EmployeeDelete(in *pb.EmployeeDeleteRequest) (tx *gorm.DB) {
	//sql := `delete from employee where password = ? and account = ?`
	dao.Exec("SET FOREIGN_KEY_CHECKS = 0")
	sql := `
			DELETE account,employee
			FROM employee
			LEFT JOIN account
			ON employee.id=account.id
			WHERE account.account=? AND account.password=?
			`
	result := dao.Exec(sql, in.Account, in.Password)
	dao.Exec("SET FOREIGN_KEY_CHECKS = 1")
	return result
}
