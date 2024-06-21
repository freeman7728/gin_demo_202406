package routers

import (
	"database_lesson/global"

	"github.com/gin-gonic/gin"
)

type User struct {
	USER_ID       int
	USER_NAME     string
	USER_PASSWORD string
	USER_KEY      string
}

func InitRouter() *gin.Engine {

	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	return router
}
