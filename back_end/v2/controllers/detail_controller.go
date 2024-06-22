package controllers

import (
	"database_lesson/dto"
	"database_lesson/models"
	"database_lesson/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 批量导入订单详情
// @Description 批量导入订单详情，并且返回成功失败列表
// @Tags detail
// @Accept  json
// @Produce  json
// @Param   list body dto.DetailList true "Detail Request"
// @Success      200  {object}  dto.DetailInsertResponse
// @Router /detail/insert [post]
func InsertDetailController(c *gin.Context) {
	var details dto.DetailList
	if err := c.ShouldBindJSON(&details); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.InsertDetailService(&details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除订单详情
// @Description 删除订单详情，返回外键检查结果和删除结果，如果有外键依赖则会提示
// @Tags detail
// @Accept  json
// @Produce  json
// @Param   detail body models.Detail true "Detail Request"
// @Success      200  string  "msg"
// @Router /detail/delete [post]
func DeleteDetailController(c *gin.Context) {
	var detail models.Detail
	if err := c.ShouldBindJSON(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.DeleteDetailService(&detail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 更新订单详情
// @Description 更新产品，除了id之外都能改，所有不会有外键冲突
// @Tags detail
// @Accept  json
// @Produce  json
// @Param   detail body models.Detail true "Detail Request"
// @Success      200  string  "msg"
// @Router /detail/update [post]
func UpdateDetailController(c *gin.Context) {
	var detail models.Detail
	if err := c.ShouldBindJSON(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.UpdateDetailService(&detail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 请求所有订单详情
// @Description 请求所有订单详情
// @Tags detail
// @Accept  json
// @Produce  json
// @Success      200  {object}  dto.DetailList
// @Router /detail/select [post]
func SelectDetailController(c *gin.Context) {
	// 转发到service层处理
	err, res := services.SelectDetailService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 通过id请求详情
// @Description 通过id请求详情
// @Tags detail
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success      200  string  models.Detail
// @Router /detail/{id} [get]
func SelectDetailByIdController(c *gin.Context) {
	var detail models.Detail
	strId := c.Param("id")
	detail.Id, _ = strconv.Atoi(strId)
	// 转发到service层处理
	err, res := services.SelectDetailById(&detail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 通过订单id请求订单详情
// @Description 通过订单id请求订单详情
// @Tags detail
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success      200  {object}  dto.DetailList
// @Router /detail/orderId/{id} [get]
func SelectDetailByOrderIdController(c *gin.Context) {
	var order models.Order
	strId := c.Param("id")
	order.Id, _ = strconv.Atoi(strId)
	// 转发到service层处理
	err, res := services.SelectDetailByOrderId(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 按照商品，统计每种商品卖出去多少个
// @Description 按照商品，统计每种商品卖出去多少个
// @Tags detail
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success      200  {object}  dto.SumGroupByProductIdList
// @Router /detail/groupByproduct [get]
func SelectSumGroupByProductIdController(c *gin.Context) {
	// 转发到service层处理
	err, res := services.SelectSumGroupByProductId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
