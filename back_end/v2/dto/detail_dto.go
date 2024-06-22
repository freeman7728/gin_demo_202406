package dto

import "database_lesson/models"

type DetailList struct {
	DetailList []models.Detail `json:"list"`
}

type DetailInsertResponse struct {
	SuccessList []models.Detail `json:"success_list"`
	FailList    []models.Detail `json:"fail_list"`
}
