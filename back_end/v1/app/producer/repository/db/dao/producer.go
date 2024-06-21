package dao

import (
	"context"
	"database_lesson/idl/pb"
	"gorm.io/gorm"
)

type ProducerDao struct {
	*gorm.DB
}

func NewProducerDao(ctx context.Context) *ProducerDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ProducerDao{NewDBClient(ctx)}
}

func (dao *ProducerDao) InsertProducer(model *pb.ProducerModel) (tx *gorm.DB) {
	sql := `insert into producer (name,short_name,address,tel,email,contact_name,contact_tel,note) VALUES(?,?,?,?,?,?,?,?)`
	return dao.DB.Exec(sql, model.Name, model.ShortName, model.Address, model.Tel, model.Email, model.ContactName, model.ContactTel, model.Note)
}
func (dao *ProducerDao) UpdateProducer(model *pb.ProducerModel) (tx *gorm.DB) {
	sql := `update producer set name = ?,short_name=?,address=?,tel=?,email=?,contact_name=?,contact_tel=?,note=? where id=?`
	return dao.DB.Exec(sql, model.Name, model.ShortName, model.Address, model.Tel, model.Email, model.ContactName, model.ContactTel, model.Note, model.Id)
}
func (dao *ProducerDao) SelectProducer(response *[]*pb.ProducerModel) (tx *gorm.DB) {
	sql := `select * from producer`
	return dao.DB.Raw(sql).Scan(&response)
}
func (dao *ProducerDao) DeleteProducer(model *pb.ProducerModel) (tx *gorm.DB) {
	sql := `delete from producer where id=?`
	return dao.DB.Exec(sql, model.Id)
}
