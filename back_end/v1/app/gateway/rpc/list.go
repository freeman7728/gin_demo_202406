package rpc

import (
	"context"
	"database_lesson/idl/pb"
)

func ListInsert(ctx context.Context, req *pb.ListInsertRequest) (resp *pb.ListInsertResponse, err error) {
	resp, err = ListService.InsertList(ctx, req)
	if err != nil {
		return
	}
	return
}
func ListDelete(ctx context.Context, req *pb.ListDeleteRequest) (resp *pb.ListDeleteResponse, err error) {
	resp, err = ListService.DeleteList(ctx, req)
	if err != nil {
		return
	}
	return
}
func ListUpdate(ctx context.Context, req *pb.ListUpdateRequest) (resp *pb.ListUpdateResponse, err error) {
	resp, err = ListService.UpdateList(ctx, req)
	if err != nil {
		return
	}
	return
}
func ListSelect(ctx context.Context, req *pb.ListSelectRequest) (resp *pb.ListSelectResponse, err error) {
	resp, err = ListService.SelectList(ctx, req)
	if err != nil {
		return
	}
	return
}
