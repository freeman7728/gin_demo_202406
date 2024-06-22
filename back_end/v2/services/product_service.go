package services

import (
	"database_lesson/dao"
	"database_lesson/dto"
	"database_lesson/models"
	"errors"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

func InsertProductService(list *dto.ProductList) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	var l dto.ProductInsertResponse
	for _, item := range list.ProductList {
		res := dao.InsertProductDao(&item)
		if res.Error != nil {
			resp.Code = http.StatusBadRequest
			item.Msg = res.Error.Error()
			var mysqlErr *mysql.MySQLError
			if errors.As(res.Error, &mysqlErr) && mysqlErr.Number == 1452 {
				resp.Message = "产品插入失败，检查供应商是否存在"
			}
			l.FailList = append(l.FailList, item)
		} else {
			l.SuccessList = append(l.SuccessList, item)
		}
	}
	resp.Data = l
	return
}
func DeleteProductService(product *models.Product) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.DeleteProductDao(product)
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
func UpdateProductService(product *models.Product) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.UpdateProductDao(product)
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
func SelectProductService() (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	var list []models.Product
	res := dao.SelectProductDao(&list)
	if res.Error != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = res.Error.Error()
		return
	}
	var temp dto.ProductList
	temp.ProductList = list
	resp.Data = temp
	return
}
