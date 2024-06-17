package dao

import (
	"context"
	"database_lesson/app/employee/repository/db/model"
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
	sql := `INSERT INTO employee (id ,name, password, tel, salary, note, level) VALUES (?,?, ?, ?, ?, ?, ?)`
	result := dao.Exec(sql, "123", in.Name, in.Password, in.Tel, in.Salary, in.Note, in.Level).Scan(nil)
	return result
}
