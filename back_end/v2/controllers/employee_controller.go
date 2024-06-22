package controllers

import (
	_ "database_lesson/dao"
	"database_lesson/dto"
	"database_lesson/models"
	"database_lesson/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
接收gin上下文，解析后转发到service处理
*/

// @Summary 批量导入用户
// @Description 批量导入用户，并且返回成功失败列表
// @Tags employee
// @Accept  json
// @Produce  json
// @Param   list body dto.EmployeeList true "Request"
// @Success      200  {object}  dto.EmployeeInsertResult
// @Router /employee/insert [post]
func InsertEmployeeController(c *gin.Context) {
	var employees dto.EmployeeList
	if err := c.ShouldBindJSON(&employees); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.InsertEmployeeService(&employees)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 用户登录
// @Description 验证账号密码
// @Tags employee
// @Accept  json
// @Produce  json
// @Param   employee body models.Employee true "Request"
// @Success      200  {object}  dto.EmployeeLoginResponse
// @Router /employee/login [post]
func LoginEmployeeController(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.LoginEmployeeService(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 用户修改信息
// @Description 验证账号密码，然后把密码改为新密码
// @Tags employee
// @Accept  json
// @Produce  json
// @Param   employee body dto.EmployeeUpdateRequest true "Request"
// @Success      200  msg  "修改成功或失败"
// @Router /employee/update [post]
func UpdateEmployeeController(c *gin.Context) {
	var employee dto.EmployeeUpdateRequest
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.UpdateEmployeeService(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除用户
// @Description 验证账号密码，然后删除用户
// @Tags employee
// @Accept  json
// @Produce  json
// @Param   employee body models.Employee true "Request"
// @Success      200  msg  "删除成功或失败"
// @Router /employee/delete [post]
func DeleteEmployeeController(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.DeleteEmployeeService(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 通过id请求详情
// @Description 通过id请求详情
// @Tags employee
// @Accept  url
// @Produce  json
// @Success      200  string  models.Employee
// @Router /employee/{id} [Get]
func SelectEmployeeByIdController(c *gin.Context) {
	var employee models.Employee
	strId := c.Param("id")
	employee.ID, _ = strconv.Atoi(strId)
	// 转发到service层处理
	err, res := services.SelectEmployeeById(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
