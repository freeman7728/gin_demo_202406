package dao

import (
	"database_lesson/dto"
	"database_lesson/global"
	"database_lesson/models"
	"gorm.io/gorm"
)

func InsertDetailDao(model *models.Detail) (tx *gorm.DB) {
	sql := `INSERT INTO detail (product_id,quantity,note,list_id) VALUES (?,?,?,?)`
	return global.DB.Exec(sql, model.ProductId, model.Quantity, model.Note, model.ListId)
}
func DeleteDetailDao(model *models.Detail) (tx *gorm.DB) {
	sql := `DELETE detail FROM detail WHERE id=?`
	return global.DB.Exec(sql, model.Id)
}
func UpdateDetailDao(model *models.Detail) (tx *gorm.DB) {
	sql := `UPDATE detail SET quantity=?,note=? WHERE id = ?`
	return global.DB.Exec(sql, model.Quantity, model.Note, model.Id)
}
func SelectDetailDao(list *[]models.Detail) (tx *gorm.DB) {
	sql := `SELECT * FROM detail`
	return global.DB.Raw(sql).Scan(list)
}
func SelectDetailById(model *models.Detail) (tx *gorm.DB) {
	sql := `SELECT * FROM detail WHERE id=?`
	return global.DB.Raw(sql, model.Id).Scan(model)
}

func SelectDetailByOrderId(model *models.Order, out *dto.DetailList) (tx *gorm.DB) {
	sql := `SELECT * FROM detail where list_id = ?`
	return global.DB.Raw(sql, model.Id).Scan(&out.DetailList)
}
