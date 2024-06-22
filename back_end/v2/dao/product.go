package dao

import (
	"database_lesson/global"
	"database_lesson/models"
	"gorm.io/gorm"
)

func InsertProductDao(model *models.Product) (tx *gorm.DB) {
	sql := `INSERT INTO product (name,price,introduction,note,producer_id) VALUES (?,?,?,?,?)`
	return global.DB.Exec(sql, model.Name, model.Price, model.Introduction, model.Note, model.ProducerId)
}
func DeleteProductDao(model *models.Product) (tx *gorm.DB) {
	sql := `DELETE product FROM product WHERE id = ?`
	return global.DB.Exec(sql, model.Id)
}
func UpdateProductDao(model *models.Product) (tx *gorm.DB) {
	sql := `UPDATE product SET name=?,price=?,introduction=?,note=?,producer_id=? WHERE id = ?;`
	return global.DB.Exec(sql, model.Name, model.Price, model.Introduction, model.Note, model.ProducerId, model.Id)
}
func SelectProductDao(list *[]models.Product) (tx *gorm.DB) {
	sql := `SELECT * FROM product`
	return global.DB.Raw(sql).Scan(list)
}
func SelectProductById(model *models.Product) (tx *gorm.DB) {
	sql := `SELECT * FROM product WHERE id=?`
	return global.DB.Raw(sql, model.Id).Scan(model)
}
