package dto

import "database_lesson/models"

type OrderList struct {
	OrderList []models.Order `json:"list"`
}

type OrderInsertResponse struct {
	SuccessList []models.Order `json:"success_list"`
	FailList    []models.Order `json:"fail_list"`
}
