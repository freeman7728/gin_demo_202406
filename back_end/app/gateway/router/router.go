package router

import (
	"database_lesson/app/gateway/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	//ginRouter.Use(middleware.Cors())
	//store := cookie.NewStore([]byte("something-very-secret"))
	//ginRouter.Use(sessions.Sessions("mysession", store))
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	employeeRouter := ginRouter.Group("/employee")
	{
		employeeRouter.POST("/signup", http.EmployeeSignupHandler)
		employeeRouter.POST("/login", http.EmployeeLoginHandler)
		employeeRouter.POST("/update", http.EmployeeUpdateHandler)
		employeeRouter.POST("/delete", http.EmployeeDeleteHandler)
	}
	producerRouter := ginRouter.Group("/producer")
	{
		producerRouter.POST("/insert", http.ProducerInsertHandler)
		producerRouter.POST("/update", http.ProducerUpdateHandler)
		producerRouter.POST("/delete", http.ProducerDeleteHandler)
		producerRouter.POST("/select", http.ProducerSelectHandler)
	}
	productRouter := ginRouter.Group("/product")
	{
		productRouter.POST("/insert", http.ProductInsertHandler)
		productRouter.POST("/update", http.ProductUpdateHandler)
		productRouter.POST("/delete", http.ProductDeleteHandler)
		productRouter.POST("/select", http.ProductSelectHandler)
	}
	listRouter := ginRouter.Group("/list")
	{
		listRouter.POST("/insert", http.ListInsertHandler)
		listRouter.POST("/update", http.ListUpdateHandler)
		listRouter.POST("/delete", http.ListDeleteHandler)
		listRouter.POST("/select", http.ListSelectHandler)
	}
	return ginRouter
}
