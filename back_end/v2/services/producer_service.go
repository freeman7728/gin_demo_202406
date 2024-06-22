package services

import (
	"database_lesson/dao"
	"database_lesson/dto"
	"database_lesson/models"
	"errors"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

func InsertProducerService(list *dto.ProducerList) (err error, resp dto.Response) {
	var res dto.ProducerInsertResponseDto
	resp.Code = http.StatusOK
	for _, item := range list.ProducerList {
		err := dao.InsertProducerDao(&item).Error
		if err != nil {
			resp.Code = http.StatusInternalServerError
			item.Msg = err.Error()
			res.FailList = append(res.FailList, item)
		} else {
			res.SuccessList = append(res.SuccessList, item)
		}
	}
	resp.Data = res
	return
}
func DeleteProducerService(producer *dto.Producer) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.DeleteProducerDao(producer.Producer)
	if res.Error != nil {
		resp.Message = res.Error.Error()
		var mysqlErr *mysql.MySQLError
		if errors.As(res.Error, &mysqlErr) && mysqlErr.Number == 1451 {
			resp.Message = "供应商不可删除，有产品依赖此供应商"
		}
	} else if res.RowsAffected == 0 {
		resp.Message = "记录不存在"
	} else {
		resp.Message = "删除成功"
	}
	return
}
func UpdateProducerService(producer *dto.Producer) (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	res := dao.UpdateProducerDao(producer.Producer)
	if res.Error != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = res.Error.Error()
		return
	} else if res.RowsAffected == 0 {
		resp.Message = "记录不存在"
	} else {
		resp.Message = "更新成功"
	}
	return
}
func SelectProducerService() (err error, resp dto.Response) {
	resp.Code = http.StatusOK
	var res dto.ProducerList
	var list []models.Producer
	err = dao.SelectProducerDao(&list).Error
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		return
	}
	res.ProducerList = list
	resp.Data = res.ProducerList
	return
}
