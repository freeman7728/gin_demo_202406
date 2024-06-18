package service

import (
	"context"
	"database_lesson/app/employee/repository/db/dao"
	"database_lesson/app/employee/repository/db/model"
	"database_lesson/idl/pb"
	"database_lesson/pkg/e"
	"fmt"
	"sync"
)

var EmployeeSrvIns *EmployeeSrv
var EmployeeSrvOnce sync.Once

type EmployeeSrv struct {
}

func GetEmployeeSrv() *EmployeeSrv {
	EmployeeSrvOnce.Do(func() {
		EmployeeSrvIns = &EmployeeSrv{}
	})
	return EmployeeSrvIns
}

func (r *EmployeeSrv) EmployeeSignup(ctx context.Context, req *pb.EmployeeRequest, resp *pb.EmployeeResponse) (err error) {
	resp.Code = e.SUCCESS
	ed := dao.NewEmployeeDao(ctx)
	for _, item := range req.EmployeeList {
		employee := &model.Employee{
			Name:     item.Name,
			Password: item.Password,
			Tel:      item.Tel,
			Salary:   item.Salary,
			Note:     item.Note,
			Level:    item.Level,
		}
		fmt.Println(employee)
		res := ed.CreateEmployee(employee)
		if res.Error != nil {
			resp.Code = e.ERROR
			return
		}
		if res.RowsAffected == 1 {
			resp.SuccessList = append(resp.SuccessList, item)
		} else {
			resp.FailList = append(resp.FailList, item)
		}
	}
	return
}
func (r *EmployeeSrv) EmployeeLogin(ctx context.Context, in *pb.EmployeeLoginRequest, out *pb.EmployeeLoginResponse) (err error) {
	out.Code = e.SUCCESS
	daoIns := dao.NewEmployeeDao(ctx)
	accountIns := &model.Login{
		Account:  in.Account,
		Password: in.Password,
	}
	outIns := &model.Login{}
	err = daoIns.LoginEmployee(accountIns, outIns).Error

	if err != nil {
		out.Code = e.ERROR
		return
	}
	if outIns.Id == 0 {
		out.Code = e.WrongPassword
		return
	}
	//out.Token = outIns.Token
	out.Id = int32(outIns.Id)
	return
}
