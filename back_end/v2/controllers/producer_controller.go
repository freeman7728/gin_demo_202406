package controllers

import (
	"database_lesson/dto"
	"database_lesson/models"
	"database_lesson/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
func SelectProducerController(c *gin.Context) {
	var producer models.Producer
	if err := c.ShouldBindJSON(&producer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.SelectProducerService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
