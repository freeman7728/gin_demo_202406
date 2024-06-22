package dao

import (
	"database_lesson/global"
	"database_lesson/models"
	"gorm.io/gorm"
)

func InsertOrderDao(model *models.Order) (tx *gorm.DB) {
	sql := `INSERT INTO list (time,note,employee_id) VALUES (?,?,?)`
	return global.DB.Exec(sql, model.Time, model.Note, model.EmployeeId)
}
func DeleteOrderDao(model *models.Order) (tx *gorm.DB) {
	sql := `DELETE list FROM list WHERE id = ?`
	return global.DB.Exec(sql, model.Id)
}
func UpdateOrderDao(model *models.Order) (tx *gorm.DB) {
	sql := `UPDATE list set quantity=?,total_price=?,time=?,note=?,employee_id=? WHERE id=?`
	return global.DB.Exec(sql, model.Quantity, model.TotalPrice, model.Time, model.Note, model.EmployeeId, model.Id)
}
func SelectOrderDao(list *[]models.Order) (tx *gorm.DB) {
	sql := `SELECT * FROM list`
	return global.DB.Raw(sql).Scan(list)
}
func SelectOrderById(model *models.Order) (tx *gorm.DB) {
	sql := `SELECT * FROM list WHERE id=?`
	return global.DB.Raw(sql, model.Id).Scan(model)
}
