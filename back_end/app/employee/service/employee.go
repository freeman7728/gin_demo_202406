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

func (r *EmployeeSrv) EmployeeSignup(ctx context.Context, req *pb.EmployeeSignupRequest, resp *pb.EmployeeSignupResponse) (err error) {
	resp.Code = e.SUCCESS
	ed := dao.NewEmployeeDao(ctx)
	for _, item := range req.EmployeeList {
		employee := &model.Employee{
			Name:   item.Name,
			Tel:    item.Tel,
			Salary: item.Salary,
			Note:   item.Note,
			Level:  item.Level,
		}
		fmt.Println(item)
		res := ed.CreateEmployee(employee)
		if res.Error == nil {
			successIns := &pb.EmployeeSignupSuccessModel{
				Tel: item.Tel,
			}
			_ = ed.GetEmployeeByTel(successIns)
			resp.SuccessList = append(resp.SuccessList, successIns)
		} else {
			failIns := &pb.EmployeeSignupFailModel{
				Name:         item.Name,
				Tel:          item.Tel,
				Salary:       item.Salary,
				Note:         item.Note,
				Level:        item.Level,
				WrongMessage: res.Error.Error(),
			}
			resp.FailList = append(resp.FailList, failIns)
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
