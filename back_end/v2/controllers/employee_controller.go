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
// @Accept  x-www-form-urlencoded
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

// @Summary 管理员删除其他用户
// @Description 需要操作者id，操作对象id，如果操作者不存在，操作对象不存在，操作权限不足都会返回对应错误
// @Tags employee
// @Accept  json
// @Produce  json
// @Param   employee body dto.DeleteEmployeeByAdminRequest true "Request"
// @Success      200  string  "删除成功"
// @Router /employee/deleteByAdmin [post]
func DeleteEmployeeByAdminController(c *gin.Context) {
	var model dto.DeleteEmployeeByAdminRequest
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.DeleteEmployeeByIdWithoutPassword(&model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 得到所有员工
// @Description 获取所有员工
// @Tags employee
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success      200  {object}  models.Employee
// @Router /employee/getAll [Get]
func SelectAllEmployeeController(c *gin.Context) {
	// 转发到service层处理
	err, res := services.SelectAllEmployee()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 根据员工id修改姓名，电话，工资，备注，级别，邮箱
// @Description 根据员工id修改姓名，电话，工资，备注，级别，邮箱
// @Tags employee
// @Accept  json
// @Produce  json
// @Param   employee body models.Employee true "Request"
// @Success      200  msg  "修改成功或失败"
// @Router /employee/updateById [post]
func UpdateEmployeeByIdController(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 转发到service层处理
	err, res := services.UpdateEmployeeById(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
