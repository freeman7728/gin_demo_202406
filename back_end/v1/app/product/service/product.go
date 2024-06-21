package service

import (
	"context"
	"database_lesson/app/product/repository/db/dao"
	"database_lesson/idl/pb"
	"database_lesson/pkg/e"
	"sync"
)

type ProductModel struct {
	// @inject_tag: json:"id" form:"id"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id"`
	// @inject_tag: json:"name" form:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	// @inject_tag: json:"price" form:"price"
	Price float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price" form:"price"`
	// @inject_tag: json:"introduction" form:"introduction"
	Introduction string `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction" form:"introduction"`
	// @inject_tag: json:"note" form:"note"
	Note string `protobuf:"bytes,5,opt,name=note,proto3" json:"note" form:"note"`
	// @inject_tag: json:"producer_id" form:"producer_id"
	ProducerId int32 `protobuf:"varint,6,opt,name=producer_id,json=producerId,proto3" json:"producer_id" form:"producer_id"` //供应商外码
}

type ProductFailToInsert struct {
	// @inject_tag: json:"id" form:"id"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id"`
	// @inject_tag: json:"name" form:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	// @inject_tag: json:"price" form:"price"
	Price float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price" form:"price"`
	// @inject_tag: json:"introduction" form:"introduction"
	Introduction string `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction" form:"introduction"`
	// @inject_tag: json:"note" form:"note"
	Note string `protobuf:"bytes,5,opt,name=note,proto3" json:"note" form:"note"`
	// @inject_tag: json:"producer_id" form:"producer_id"
	ProducerId int32 `protobuf:"varint,6,opt,name=producer_id,json=producerId,proto3" json:"producer_id" form:"producer_id"` //供应商外码
	// @inject_tag: json:"msg" form:"msg"
	Msg string `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg" form:"msg"`
}

type ProductInsertRequest struct {
	// @inject_tag: json:"product" form:"product"
	Product []ProductModel `json:"product" form:"product"`
}

type ProductInsertResponse struct {
	// @inject_tag: json:"success_list" form:"success_list"
	SuccessList []*ProductModel `protobuf:"bytes,1,rep,name=success_list,json=successList,proto3" json:"success_list" form:"success_list"`
	// @inject_tag: json:"fail_list" form:"fail_list"
	FailList []*ProductFailToInsert `protobuf:"bytes,2,rep,name=fail_list,json=failList,proto3" json:"fail_list" form:"fail_list"`
	// @inject_tag: json:"code" form:"code"
	Code int32 `protobuf:"varint,3,opt,name=code,proto3" json:"code" form:"code"`
	// @inject_tag: json:"msg" form:"msg"
	Msg string `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg" form:"msg"`
}

//以上结构体用于swagger扫描生成接口文档，代码中不做实际使用

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

// @Summary Insert products
// @Description Insert multiple products into the database
// @Tags product
// @Accept  json
// @Produce  json
// @Param   product body ProductInsertRequest true "Product Insert Request"
// @Success      200  {object}  ProductInsertResponse
// @Router /product/insert [post]
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
