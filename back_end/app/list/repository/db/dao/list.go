package dao

import (
	"context"
	"database_lesson/idl/pb"
	"gorm.io/gorm"
)

type ListDao struct {
	*gorm.DB
}

func NewListDao(ctx context.Context) *ListDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ListDao{NewDBClient(ctx)}
}

func (dao *ListDao) InsertList(model *pb.ListModel) (tx *gorm.DB) {
	sql := `INSERT INTO list (quantity,total_price,time,note,employee_id) VALUES (?,?,?,?,?)`
	return dao.DB.Exec(sql, model.Quantity, model.TotalPrice, model.Time, model.Note, model.EmployeeId)
}
func (dao *ListDao) DeleteList(model *pb.ListDeleteRequest) (tx *gorm.DB) {
	sql := `DELETE list FROM list WHERE id = ?`
	return dao.DB.Exec(sql, model.List.Id)
}
func (dao *ListDao) UpdateList(model *pb.ListUpdateRequest) (tx *gorm.DB) {
	sql := `UPDATE list set quantity=?,total_price=?,time=?,note=?,employee_id=? WHERE id=?`
	return dao.DB.Exec(sql, model.List.Quantity, model.List.TotalPrice, model.List.Time, model.List.Note, model.List.EmployeeId, model.List.Id)
}
func (dao *ListDao) SelectList(model *pb.ListSelectResponse) (tx *gorm.DB) {
	sql := `SELECT * FROM list`
	return dao.DB.Raw(sql).Scan(&model.List)
}
