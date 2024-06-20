package main

import (
	"database_lesson/app/product/repository/db/dao"
	"database_lesson/app/product/service"
	"database_lesson/config"
	"database_lesson/idl/pb"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitDB()
	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcProductService"), // 微服务名字
		micro.Address(config.ProductServiceAddress),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterProductServiceHandler(microService.Server(), service.GetProductSrv())
	// 启动微服务
	_ = microService.Run()
}
