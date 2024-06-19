package router

import (
	"database_lesson/app/gateway/http"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	//ginRouter.Use(middleware.Cors())
	//store := cookie.NewStore([]byte("something-very-secret"))
	//ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/test")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		v1.POST("/user/register", http.UserRegisterHandler)
		//v1.POST("/user/login", http.UserLoginHandler)
	}
	employeeRouter := ginRouter.Group("/employee")
	{
		employeeRouter.POST("/signup", http.EmployeeSignupHandler)
		employeeRouter.POST("/login", http.EmployeeLoginHandler)
		employeeRouter.POST("/update", http.EmployeeUpdateHandler)
		employeeRouter.POST("/delete", http.EmployeeDeleteHandler)
	}
	return ginRouter
}
