package routers

import (
	"database_lesson/controllers"
	"database_lesson/global"
	"database_lesson/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.Use(middleware.SetupCorsMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	employeeRouter := router.Group("/employee")
	{
		employeeRouter.POST("/insert", controllers.InsertEmployeeController)
		employeeRouter.POST("/login", controllers.LoginEmployeeController)
		employeeRouter.POST("/update", controllers.UpdateEmployeeController)
		employeeRouter.POST("/delete", controllers.DeleteEmployeeController)
	}
	producerRouter := router.Group("/producer")
	{
		producerRouter.POST("/insert", controllers.InsertProducerController)
		producerRouter.POST("/delete", controllers.DeleteProducerController)
		producerRouter.POST("/update", controllers.UpdateProducerController)
		producerRouter.POST("/select", controllers.SelectProducerController)
	}

	return router
}
