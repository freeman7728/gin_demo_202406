package controllers

import (
	"database_lesson/dto"
	"database_lesson/models"
	"database_lesson/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 批量导入订单
// @Description 批量导入订单，并且返回成功失败列表
// @Tags order
// @Accept  json
// @Produce  json
// @Param   list body dto.OrderList true "Order Request"
// @Success      200  {object}  dto.OrderInsertResponse
// @Router /order/insert [post]
func InsertOrderController(c *gin.Context) {
	var orders dto.OrderList
	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.InsertOrderService(&orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除订单
// @Description 删除订单，返回外键检查结果和删除结果，如果有外键依赖则会提示
// @Tags order
// @Accept  json
// @Produce  json
// @Param   order body models.Order true "Order Request"
// @Success      200  string  "msg"
// @Router /order/delete [post]
func DeleteOrderController(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.DeleteOrderService(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 更新订单
// @Description 更新产品，除了id之外都能改，所有不会有外键冲突
// @Tags order
// @Accept  json
// @Produce  json
// @Param   order body models.Order true "Order Request"
// @Success      200  string  "msg"
// @Router /order/update [post]
func UpdateOrderController(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.UpdateOrderService(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 请求所有订单
// @Description 请求所有订单
// @Tags order
// @Accept  json
// @Produce  json
// @Success      200  {object}  dto.OrderList
// @Router /order/select [post]
func SelectOrderController(c *gin.Context) {
	// 转发到service层处理
	err, res := services.SelectOrderService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 通过id请求详情
// @Description 通过id请求详情
// @Tags order
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success      200  string  models.Order
// @Router /order/{id} [Get]
func SelectOrderByIdController(c *gin.Context) {
	var order models.Order
	strId := c.Param("id")
	order.Id, _ = strconv.Atoi(strId)
	// 转发到service层处理
	err, res := services.SelectOrderById(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
