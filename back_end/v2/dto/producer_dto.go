package dto

import "database_lesson/models"

type ProducerInsertDto struct {
	ProducerList *[]models.Producer `json:"list"`
}

type ProducerInsertResponseDto struct {
	FailList    []models.Producer `json:"fail_list"`
	SuccessList []models.Producer `json:"success_list"`
}

type ProducerList struct {
	ProducerList []models.Producer `json:"list"`
}

type Producer struct {
	Producer models.Producer `json:"producer"`
}
