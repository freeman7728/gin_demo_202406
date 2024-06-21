package dao

import (
	"context"
	"database_lesson/idl/pb"
	"gorm.io/gorm"
)

type DetailDao struct {
	*gorm.DB
}

func NewDetailDao(ctx context.Context) *DetailDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &DetailDao{NewDBClient(ctx)}
}

func (p *DetailDao) InsertDetail(model *pb.DetailModel) (tx *gorm.DB) {
	sql := `INSERT INTO detail (product_id,quantity,note,list_id) VALUES (?,?,?,?)`
	return p.Exec(sql, model.ProductId, model.Quantity, model.Note, model.ListId)
}

func (p *DetailDao) DeleteDetail(model *pb.DetailDeleteRequest) (tx *gorm.DB) {
	sql := `DELETE detail FROM detail WHERE id=?`
	return p.Exec(sql, model.Detail.Id)
}

func (p *DetailDao) UpdateDetail(model *pb.DetailUpdateRequest) (tx *gorm.DB) {
	sql := `UPDATE detail SET quantity=?,note=? WHERE id = ?`
	return p.Exec(sql, model.Detail.Quantity, model.Detail.Note, model.Detail.Id)
}

func (p *DetailDao) SelectDetail(model *pb.DetailSelectResponse) (tx *gorm.DB) {
	sql := `SELECT * FROM detail`
	return p.Raw(sql).Scan(&model.DetailList)
}
