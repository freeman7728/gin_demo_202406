package rpc

import (
	"go-micro.dev/v4"

	"database_lesson/idl/pb"
)

var (
	UserService     pb.UserService
	EmployeeService pb.EmployeeService
	AuthService     pb.AuthService
)

func InitRPC() {
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		//micro.WrapClient(wrappers.NewUserWrapper),
	)
	employeeMicroService := micro.NewService(
		micro.Name("employeeService.client"),
	)
	authMicroService := micro.NewService(
		micro.Name("authMicroService.client"),
	)
	// 用户服务调用实例
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	employeeService := pb.NewEmployeeService("rpcEmployeeService", employeeMicroService.Client())
	authService := pb.NewAuthService("rpcAuthService", authMicroService.Client())

	UserService = userService
	EmployeeService = employeeService
	AuthService = authService
}
