package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCorsMiddleware() gin.HandlerFunc {
	// 使用默认配置
	// 允许所有源，允许所有方法（GET, POST等），允许所有标头，允许凭据
	//config := cors.DefaultConfig()

	// 或者使用自定义配置
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173",
		"http://8.146.198.97:81"} // 允许特定源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许特定方法
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // 允许特定标头

	return cors.New(config)
}
