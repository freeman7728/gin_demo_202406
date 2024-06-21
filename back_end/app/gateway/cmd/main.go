package main

import (
	"database_lesson/app/gateway/rpc"
	"time"

	"go-micro.dev/v4/web"

	"database_lesson/app/gateway/router"
	"database_lesson/config"
	_ "database_lesson/docs"
)

// @title           超市进销存系统--数据库课设接口文档
// @version         1.0
// @description     没有描述
// @host      localhost:4000
// @BasePath  /
func main() {
	config.Init()
	rpc.InitRPC()
	//etcdReg := etcd.NewRegistry(
	//	registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	//)

	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		// 将服务调用实例使用gin处理
		web.Handler(router.NewRouter()),
		//web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	// 接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
