package dao

import (
	"context"
	"database_lesson/idl/pb"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ProductDao{NewDBClient(ctx)}
}

func (p *ProductDao) InsertProduct(model *pb.ProductModel) (tx *gorm.DB) {
	sql := `INSERT INTO product (name,price,introduction,note,producer_id) VALUES (?,?,?,?,?)`
	return p.Exec(sql, model.Name, model.Price, model.Introduction, model.Note, model.ProducerId)
}

func (p *ProductDao) DeleteProduct(model *pb.ProductModel) (tx *gorm.DB) {
	sql := `DELETE product FROM product WHERE id = ?`
	return p.Exec(sql, model.Id)
}

func (p *ProductDao) UpdateProduct(model *pb.ProductModel) (tx *gorm.DB) {
	sql := `UPDATE product SET name=?,price=?,introduction=?,note=?,producer_id=? WHERE id = ?;`
	return p.Exec(sql, model.Name, model.Price, model.Introduction, model.Note, model.ProducerId, model.Id)
}

func (p *ProductDao) SelectProduct(model *pb.ProductSelectResponse) (tx *gorm.DB) {
	sql := `SELECT * FROM product`
	return p.Raw(sql).Scan(&model.Products)
}
