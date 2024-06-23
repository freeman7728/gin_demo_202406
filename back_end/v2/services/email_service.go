package services

import (
	"database_lesson/dao"
	"database_lesson/dto"
	"database_lesson/global"
	"database_lesson/models"
	"database_lesson/utils"
	"net/http"
	"strconv"
)

func GenCode(employee *models.Employee) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.SelectEmployeeById(employee)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	if res.RowsAffected == 0 {
		resp.Message = "记录不存在"
		resp.Code = http.StatusBadRequest
		return
	}
	//此时employee拿到了邮箱，接下来生成验证码，返回验证码是否生成成功
	code, result := utils.SendEmailValidate([]string{employee.Email})
	if result != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = result.Error()
		return
	}
	//接下来把把员工id作为键，验证码作为值插进去
	var ins dto.CertifyDto
	ins.Code = code
	ins.ID = employee.ID
	err = utils.InsertCodeIntoRedis(&ins)
	resp.Message = "生成成功,查看邮箱"
	return
}

func VerifyCodeCode(certifyDto *dto.CertifyDto) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	codeInRedis := global.RedisClient.Get(strconv.Itoa(certifyDto.ID))
	if codeInRedis == nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "验证码超时"
		return
	}
	if codeInRedis.Val() != certifyDto.Code {
		resp.Code = http.StatusBadRequest
		resp.Message = "验证码错误"
		return
	}
	res := dao.SetEmailAuthById(certifyDto)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	resp.Message = "验证码正确"
	return
}
