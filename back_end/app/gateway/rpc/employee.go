package rpc

import (
	"context"
	"database_lesson/idl/pb"
)

func EmployeeSignup(ctx context.Context, req *pb.EmployeeRequest) (resp *pb.EmployeeResponse, err error) {
	resp, err = EmployeeService.EmployeeSignup(ctx, req)
	if err != nil {
		return
	}
	return
}
