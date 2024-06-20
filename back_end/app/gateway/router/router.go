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
	return ginRouter
}
