package services

import (
	"database_lesson/dao"
	"database_lesson/dto"
	"database_lesson/models"
	"errors"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

func InsertDetailService(list *dto.DetailList) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	var l dto.DetailInsertResponse
	for _, item := range list.DetailList {
		res := dao.InsertDetailDao(&item)
		if res.Error != nil {
			resp.Code = http.StatusBadRequest
			item.Msg = res.Error.Error()
			var mysqlErr *mysql.MySQLError
			if errors.As(res.Error, &mysqlErr) && mysqlErr.Number == 1452 {
				item.Msg = "检查订单或者员工编号，该员工或者订单不存在"
			}
			l.FailList = append(l.FailList, item)
		} else {
			l.SuccessList = append(l.SuccessList, item)
		}
	}
	resp.Data = l
	return
}
func DeleteDetailService(order *models.Detail) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.DeleteDetailDao(order)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		var mysqlErr *mysql.MySQLError
		if errors.As(res.Error, &mysqlErr) && mysqlErr.Number == 1451 {
			resp.Message = "外键冲突"
		}
		return
	}
	if res.RowsAffected == 0 {
		resp.Code = http.StatusNotFound
		resp.Message = "记录不存在"
		return
	}
	if res.RowsAffected == 1 {
		resp.Message = "删除成功"
	}
	return
}
func UpdateDetailService(order *models.Detail) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.UpdateDetailDao(order)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		var mysqlErr *mysql.MySQLError
		if errors.As(res.Error, &mysqlErr) && mysqlErr.Number == 1451 {
			resp.Message = "订单详情不可更改，检查订单是否存在"
		}
		return
	}
	if res.RowsAffected == 0 {
		resp.Code = http.StatusNotFound
		resp.Message = "记录不存在"
		return
	}
	if res.RowsAffected == 1 {
		resp.Message = "更改成功"
	}
	return
}
func SelectDetailService() (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	var list []models.Detail
	res := dao.SelectDetailDao(&list)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	var temp dto.DetailList
	temp.DetailList = list
	resp.Data = temp
	return
}
func SelectDetailById(detail *models.Detail) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.SelectDetailById(detail)
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
	resp.Data = detail
	resp.Message = "查询成功"
	return
}
