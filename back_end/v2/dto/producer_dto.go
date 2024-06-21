package dto

import "database_lesson/models"

type ProducerInsertDto struct {
	ProducerList *[]models.Producer
}

type ProducerInsertResponseDto struct {
	FailList    []models.Producer
	SuccessList []models.Producer
}

type ProducerList struct {
	ProducerList *[]models.Producer
}
