package services

import (
	"database_lesson/dao"
	"database_lesson/dto"
	"database_lesson/models"
	"net/http"
)

func InsertProducerService(list *dto.ProducerList) (err error, resp *dto.Response) {
	var res dto.ProducerInsertResponseDto
	resp.Code = http.StatusOK
	for _, item := range *list.ProducerList {
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
func DeleteProducerService(producer *models.Producer) (err error, resp *dto.Response) {
	resp.Code = http.StatusOK
	res := dao.DeleteProducerDao(producer)
	if res.Error != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = res.Error.Error()
		return
	} else if res.RowsAffected == 0 {
		resp.Message = "账号或密码错误"
	} else {
		resp.Message = "删除成功"
	}
	return
}
func UpdateProducerService(producer *models.Producer) (err error, resp *dto.Response) {
	resp.Code = http.StatusOK
	res := dao.DeleteProducerDao(producer)
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
func SelectProducerService() (err error, resp *dto.Response) {
	resp.Code = http.StatusOK
	var res dto.ProducerList
	err = dao.SelectProducerDao(&res).Error
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		return
	}
	resp.Data = res
	return
}
