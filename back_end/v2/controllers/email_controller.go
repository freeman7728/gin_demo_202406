package controllers

import (
	"database_lesson/dto"
	"database_lesson/models"
	"database_lesson/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 发送邮箱验证码
// @Description 使用用户id，向对应用户的邮箱发送验证码
// @Tags email
// @Accept  json
// @Produce  json
// @Param   detail body models.Employee true "Employee Request"
// @Success      200  string  "msg"
// @Router /email [post]
func GenCodeController(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.GenCode(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 检查邮箱验证码
// @Description 使用用户id，和验证码检查
// @Tags email
// @Accept  json
// @Produce  json
// @Param   detail body dto.CertifyDto true "certify Request"
// @Success      200  string  "msg"
// @Router /email/auth [post]
func VerifyCodeController(c *gin.Context) {
	var v dto.CertifyDto
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, res := services.VerifyCodeCode(&v)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
