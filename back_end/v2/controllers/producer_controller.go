package controllers

import (
	"database_lesson/dto"
	"database_lesson/models"
	"database_lesson/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 批量导入供应商
// @Description 批量导入供应商，并且返回成功失败列表
// @Tags producer
// @Accept  json
// @Produce  json
// @Param   product body dto.ProducerList true "Producer Request"
// @Success      200  {object}  dto.ProducerInsertResponseDto
// @Router /producer/insert [post]
func InsertProducerController(c *gin.Context) {
	var producers dto.ProducerList
	if err := c.ShouldBindJSON(&producers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.InsertProducerService(&producers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除供应商
// @Description 删除供应商，返回外键检查结果和删除结果，如果有外键依赖则会提示
// @Tags producer
// @Accept  json
// @Produce  json
// @Param   product body models.Producer true "Producer Request"
// @Success      200  string  "msg"
// @Router /producer/delete [post]
func DeleteProducerController(c *gin.Context) {
	var producer models.Producer
	if err := c.ShouldBindJSON(&producer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.DeleteProducerService(&producer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 更新供应商
// @Description 更新供应商，除了id之外都能改，所有不会有外键冲突
// @Tags producer
// @Accept  json
// @Produce  json
// @Param   product body models.Producer true "Producer Request"
// @Success      200  string  "msg"
// @Router /producer/update [post]
func UpdateProducerController(c *gin.Context) {
	var producer models.Producer
	if err := c.ShouldBindJSON(&producer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.UpdateProducerService(&producer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 请求所有供应商
// @Description 删除供应商，返回外键检查结果和删除结果，如果有外键依赖则会提示
// @Tags producer
// @Accept  json
// @Produce  json
// @Success      200  {object}  dto.ProducerList
// @Router /producer/select [post]
func SelectProducerController(c *gin.Context) {
	// 转发到service层处理
	err, res := services.SelectProducerService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 通过id请求详情
// @Description 通过id请求详情
// @Tags producer
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success      200  string  models.Producer
// @Router /producer/{id} [Get]
func SelectProducerByIdController(c *gin.Context) {
	var producer models.Producer
	strId := c.Param("id")
	producer.ID, _ = strconv.Atoi(strId)
	// 转发到service层处理
	err, res := services.SelectProducerById(&producer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
