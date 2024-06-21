package rpc

import (
	"context"
	"database_lesson/idl/pb"
)

func ProductInsert(ctx context.Context, req *pb.ProductInsertRequest) (resp *pb.ProductInsertResponse, err error) {
	resp, err = ProductService.InsertProduct(ctx, req)
	if err != nil {
		return
	}
	return
}
func ProductDelete(ctx context.Context, req *pb.ProductDeleteRequest) (resp *pb.ProductDeleteResponse, err error) {
	resp, err = ProductService.DeleteProduct(ctx, req)
	if err != nil {
		return
	}
	return
}
func ProductUpdate(ctx context.Context, req *pb.ProductUpdateRequest) (resp *pb.ProductUpdateResponse, err error) {
	resp, err = ProductService.UpdateProduct(ctx, req)
	if err != nil {
		return
	}
	return
}
func ProductSelect(ctx context.Context, req *pb.ProductSelectRequest) (resp *pb.ProductSelectResponse, err error) {
	resp, err = ProductService.SelectProduct(ctx, req)
	if err != nil {
		return
	}
	return
}
