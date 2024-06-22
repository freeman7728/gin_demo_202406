package dao

import (
	"database_lesson/global"
	"database_lesson/models"
	"gorm.io/gorm"
)

func InsertProducerDao(model *models.Producer) (tx *gorm.DB) {
	sql := `insert into producer (name,short_name,address,tel,email,contact_name,contact_tel,note) VALUES(?,?,?,?,?,?,?,?)`
	return global.DB.Exec(sql, model.Name, model.ShortName, model.Address, model.Tel, model.Email, model.ContactName, model.ContactTel, model.Note)
}
func DeleteProducerDao(model models.Producer) (tx *gorm.DB) {
	sql := `delete from producer where id=?`
	return global.DB.Exec(sql, model.ID)
}
func UpdateProducerDao(model models.Producer) (tx *gorm.DB) {
	sql := `update producer set name = ?,short_name=?,address=?,tel=?,email=?,contact_name=?,contact_tel=?,note=? where id=?`
	return global.DB.Exec(sql, model.Name, model.ShortName, model.Address, model.Tel, model.Email, model.ContactName, model.ContactTel, model.Note, model.ID)
}
func SelectProducerDao(list *[]models.Producer) (tx *gorm.DB) {
	sql := `select * from producer`
	return global.DB.Raw(sql).Scan(list)
}
