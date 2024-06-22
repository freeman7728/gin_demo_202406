package controllers

import (
	"database_lesson/dto"
	"database_lesson/models"
	"database_lesson/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 批量导入产品
// @Description 批量导入产品，并且返回成功失败列表
// @Tags product
// @Accept  json
// @Produce  json
// @Param   product body dto.ProductList true "Producer Request"
// @Success      200  {object}  dto.ProductInsertResponse
// @Router /product/insert [post]
func InsertProductController(c *gin.Context) {
	var products dto.ProductList
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.InsertProductService(&products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除产品
// @Description 删除产品，返回外键检查结果和删除结果，如果有外键依赖则会提示
// @Tags product
// @Accept  json
// @Produce  json
// @Param   product body models.Product true "Producer Request"
// @Success      200  string  "msg"
// @Router /product/delete [post]
func DeleteProductController(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.DeleteProductService(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 更新产品
// @Description 更新产品，除了id之外都能改，所有不会有外键冲突
// @Tags product
// @Accept  json
// @Produce  json
// @Param   product body models.Product true "Producer Request"
// @Success      200  string  "msg"
// @Router /product/update [post]
func UpdateProductController(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.UpdateProductService(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 请求所有产品
// @Description 删除产品，返回外键检查结果和删除结果，如果有外键依赖则会提示
// @Tags product
// @Accept  json
// @Produce  json
// @Success      200  {object}  dto.ProductList
// @Router /product/select [post]
func SelectProductController(c *gin.Context) {
	// 转发到service层处理
	err, res := services.SelectProductService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
