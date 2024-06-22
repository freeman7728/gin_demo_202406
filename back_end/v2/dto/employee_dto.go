package dto

import "database_lesson/models"

type EmployeeList struct {
	EmployeeList []models.Employee `json:"list"`
}

type EmployeeInsertResult struct {
	Successful []models.Employee `json:"success_list"`
	Failed     []models.Employee `json:"fail_list"`
}

type EmployeeUpdateRequest struct {
	models.Employee `json:",inline"`
	NewPassword     string `json:"new_password"`
}
