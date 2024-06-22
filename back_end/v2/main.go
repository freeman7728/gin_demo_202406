package main

import (
	"database_lesson/core"
	_ "database_lesson/docs"
	"database_lesson/global"
	"database_lesson/routers"
)

// @title           超市进销存系统--数据库课设接口文档
// @version         1.0
// @description     没有描述
// @host      localhost:4000
// @BasePath  /
func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.GormInit()
	global.RedisClient = core.RedisInit()

	router := routers.InitRouter()
	addr := global.Config.System.Addr()

	global.Log.Infof("服务运行在%s", addr)
	_ = router.Run(addr)
}
