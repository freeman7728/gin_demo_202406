package main

import (
	"database_lesson/core"
	"database_lesson/global"
	"database_lesson/routers"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.GormInit()

	router := routers.InitRouter()
	addr := global.Config.System.Addr()

	global.Log.Infof("服务运行在%s", addr)
	_ = router.Run(addr)
}
