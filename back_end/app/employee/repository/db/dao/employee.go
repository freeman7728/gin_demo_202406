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
