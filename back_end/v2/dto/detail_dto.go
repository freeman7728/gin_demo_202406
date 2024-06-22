package dto

import "database_lesson/models"

type DetailList struct {
	DetailList []models.Detail `json:"list"`
}

type DetailInsertResponse struct {
	SuccessList []models.Detail `json:"success_list"`
	FailList    []models.Detail `json:"fail_list"`
}

type SumGroupByProductId struct {
	Id  int `json:"id"`
	Sum int `json:"sum"`
}

type SumGroupByProductIdList struct {
	List []SumGroupByProductId `json:"list"`
}
