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
		employeeRouter.GET("/:id", controllers.SelectEmployeeByIdController)
		employeeRouter.POST("/deleteByAdmin", controllers.DeleteEmployeeByAdminController)
	}
	producerRouter := router.Group("/producer")
	{
		producerRouter.POST("/insert", controllers.InsertProducerController)
		producerRouter.POST("/delete", controllers.DeleteProducerController)
		producerRouter.POST("/update", controllers.UpdateProducerController)
		producerRouter.POST("/select", controllers.SelectProducerController)
		producerRouter.GET("/:id", controllers.SelectProducerByIdController)
	}
	productRouter := router.Group("/product")
	{
		productRouter.POST("/insert", controllers.InsertProductController)
		productRouter.POST("/delete", controllers.DeleteProductController)
		productRouter.POST("/update", controllers.UpdateProductController)
		productRouter.POST("/select", controllers.SelectProductController)
		productRouter.GET("/:id", controllers.SelectProductByIdController)
	}
	orderRouter := router.Group("/order")
	{
		orderRouter.POST("/insert", controllers.InsertOrderController)
		orderRouter.POST("/delete", controllers.DeleteOrderController)
		orderRouter.POST("/update", controllers.UpdateOrderController)
		orderRouter.POST("/select", controllers.SelectOrderController)
		orderRouter.GET("/:id", controllers.SelectOrderByIdController)
	}
	detailRouter := router.Group("/detail")
	{
		detailRouter.POST("/insert", controllers.InsertDetailController)
		detailRouter.POST("/delete", controllers.DeleteDetailController)
		detailRouter.POST("/update", controllers.UpdateDetailController)
		detailRouter.POST("/select", controllers.SelectDetailController)
		detailRouter.GET("/:id", controllers.SelectDetailByIdController)
	}

	return router
}
