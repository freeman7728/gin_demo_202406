package dto

import "database_lesson/models"

type ProductList struct {
	ProductList []models.Product `json:"list"`
}

type ProductInsertResponse struct {
	SuccessList []models.Product `json:"success_list"`
	FailList    []models.Product `json:"fail_list"`
}
