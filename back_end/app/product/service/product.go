package service

import (
	"context"
	"database_lesson/app/product/repository/db/dao"
	"database_lesson/idl/pb"
	"database_lesson/pkg/e"
	"sync"
)

var ProductSrvIns *ProductSrv
var ProductSrvOnce sync.Once

type ProductSrv struct {
}

func GetProductSrv() *ProductSrv {
	ProductSrvOnce.Do(func() {
		ProductSrvIns = &ProductSrv{}
	})
	return ProductSrvIns
}

func (p *ProductSrv) InsertProduct(ctx context.Context, in *pb.ProductInsertRequest, out *pb.ProductInsertResponse) (err error) {
	daoIns := dao.NewProductDao(ctx)
	out.Code = e.SUCCESS
	for _, model := range in.Product {
		res := daoIns.InsertProduct(model)
		if res.Error != nil {
			out.Code = e.ERROR
			var temp = &pb.ProductFailToInsert{
				Name:         model.Name,
				Id:           model.Id,
				Price:        model.Price,
				Introduction: model.Introduction,
				Note:         model.Note,
				ProducerId:   model.ProducerId,
				Msg:          res.Error.Error(),
			}
			out.FailList = append(out.FailList, temp)
		} else {
			out.SuccessList = append(out.SuccessList, model)
		}
	}
	return
}

func (p *ProductSrv) DeleteProduct(ctx context.Context, in *pb.ProductDeleteRequest, out *pb.ProductDeleteResponse) (err error) {
	daoIns := dao.NewProductDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.DeleteProduct(in.Product)
	if res.Error != nil {
		out.Code = e.ERROR
		out.Success = false
		out.Msg = res.Error.Error()
		return
	}
	if res.RowsAffected == 1 {
		out.Success = true
	}
	return
}

func (p *ProductSrv) UpdateProduct(ctx context.Context, in *pb.ProductUpdateRequest, out *pb.ProductUpdateResponse) (err error) {
	daoIns := dao.NewProductDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.UpdateProduct(in.Product)
	if res.Error != nil {
		out.Code = e.ERROR
		out.Success = false
		out.Msg = res.Error.Error()
		return
	}
	if res.RowsAffected == 1 {
		out.Success = true
	}
	return
}

func (p *ProductSrv) SelectProduct(ctx context.Context, in *pb.ProductSelectRequest, out *pb.ProductSelectResponse) (err error) {
	daoIns := dao.NewProductDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.SelectProduct(out)
	if res.Error != nil {
		out.Code = e.ERROR
		return
	}
	return
}
