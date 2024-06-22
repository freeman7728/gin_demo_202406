package services

import (
	"database_lesson/dao"
	"database_lesson/dto"
	"database_lesson/models"
	"errors"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

func InsertOrderService(list *dto.OrderList) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	var l dto.OrderInsertResponse
	for _, item := range list.OrderList {
		res := dao.InsertOrderDao(&item)
		if res.Error != nil {
			resp.Code = http.StatusBadRequest
			item.Msg = res.Error.Error()
			var mysqlErr *mysql.MySQLError
			if errors.As(res.Error, &mysqlErr) && mysqlErr.Number == 1452 {
				item.Msg = "检查员工编号，该员工不存在"
			}
			l.FailList = append(l.FailList, item)
		} else {
			l.SuccessList = append(l.SuccessList, item)
		}
	}
	resp.Data = l
	return
}
func DeleteOrderService(order *models.Order) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.DeleteOrderDao(order)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		var mysqlErr *mysql.MySQLError
		if errors.As(res.Error, &mysqlErr) && mysqlErr.Number == 1451 {
			resp.Message = "产品不可删除，检查订购了该产品的订单"
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
func UpdateOrderService(order *models.Order) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.UpdateOrderDao(order)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		var mysqlErr *mysql.MySQLError
		if errors.As(res.Error, &mysqlErr) && mysqlErr.Number == 1451 {
			resp.Message = "产品不可更改，检查供应商是否存在"
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
func SelectOrderService() (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	var list []models.Order
	res := dao.SelectOrderDao(&list)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	var temp dto.OrderList
	temp.OrderList = list
	resp.Data = temp
	return
}
func SelectOrderById(order *models.Order) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.SelectOrderById(order)
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
	resp.Data = order
	resp.Message = "查询成功"
	return
}
