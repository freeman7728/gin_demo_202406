package rpc

import (
	"go-micro.dev/v4"

	"database_lesson/app/gateway/wrappers"
	"database_lesson/idl/pb"
)

var (
	UserService pb.UserService
)

func InitRPC() {
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	// 用户服务调用实例
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	// task
	UserService = userService

}
