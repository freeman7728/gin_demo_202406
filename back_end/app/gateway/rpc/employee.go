package rpc

import (
	"context"
	"database_lesson/idl/pb"
)

func EmployeeSignup(ctx context.Context, req *pb.EmployeeSignupRequest) (resp *pb.EmployeeSignupResponse, err error) {
	resp, err = EmployeeService.EmployeeSignup(ctx, req)
	if err != nil {
		return
	}
	return
}
func EmployeeLogin(ctx context.Context, req *pb.EmployeeLoginRequest) (resp *pb.EmployeeLoginResponse, err error) {
	resp, err = EmployeeService.EmployeeLogin(ctx, req)
	if err != nil {
		return
	}
	return
}
