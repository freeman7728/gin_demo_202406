package rpc

import (
	"context"
	"database_lesson/idl/pb"
)

func DetailInsert(ctx context.Context, req *pb.DetailInsertRequest) (resp *pb.DetailInsertResponse, err error) {
	resp, err = DetailService.InsertDetail(ctx, req)
	if err != nil {
		return
	}
	return
}
func DetailDelete(ctx context.Context, req *pb.DetailDeleteRequest) (resp *pb.DetailDeleteResponse, err error) {
	resp, err = DetailService.DeleteDetail(ctx, req)
	if err != nil {
		return
	}
	return
}
func DetailUpdate(ctx context.Context, req *pb.DetailUpdateRequest) (resp *pb.DetailUpdateResponse, err error) {
	resp, err = DetailService.UpdateDetail(ctx, req)
	if err != nil {
		return
	}
	return
}
func DetailSelect(ctx context.Context, req *pb.DetailSelectRequest) (resp *pb.DetailSelectResponse, err error) {
	resp, err = DetailService.SelectDetail(ctx, req)
	if err != nil {
		return
	}
	return
}
