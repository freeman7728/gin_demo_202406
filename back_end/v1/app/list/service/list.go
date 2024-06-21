package service

import (
	"context"
	"database_lesson/app/list/repository/db/dao"
	"database_lesson/idl/pb"
	"database_lesson/pkg/e"
	"sync"
)

var ListServIns *ListServ
var ListServInsOnce sync.Once

type ListServ struct{}

func GetListServ() *ListServ {
	ListServInsOnce.Do(func() {
		ListServIns = &ListServ{}
	})
	return ListServIns
}

func (l *ListServ) InsertList(ctx context.Context, in *pb.ListInsertRequest, out *pb.ListInsertResponse) (err error) {
	daoIns := dao.NewListDao(ctx)
	out.Code = e.SUCCESS
	for _, model := range in.List {
		res := daoIns.InsertList(model)
		if res.Error != nil {
			out.Code = e.ERROR
			model.Msg = res.Error.Error()
			out.FailList = append(out.FailList, model)
		} else {
			out.SuccessList = append(out.SuccessList, model)
		}
	}
	return
}

func (l *ListServ) DeleteList(ctx context.Context, in *pb.ListDeleteRequest, out *pb.ListDeleteResponse) (err error) {
	daoIns := dao.NewListDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.DeleteList(in)
	if res.Error != nil {
		out.Code = e.ERROR
		out.Success = false
		return
	}
	if res.RowsAffected == 0 {
		out.Code = e.ERROR
		out.Success = false
		return
	}
	out.Success = true
	return
}

func (l *ListServ) UpdateList(ctx context.Context, in *pb.ListUpdateRequest, out *pb.ListUpdateResponse) (err error) {
	daoIns := dao.NewListDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.UpdateList(in)
	if res.Error != nil {
		out.Code = e.ERROR
		out.Success = false
		out.Msg = res.Error.Error()
		return
	}
	if res.RowsAffected == 0 {
		out.Code = e.ERROR
		out.Success = false
		return
	}
	out.Success = true
	return
}

func (l *ListServ) SelectList(ctx context.Context, in *pb.ListSelectRequest, out *pb.ListSelectResponse) (err error) {
	daoIns := dao.NewListDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.SelectList(out)
	if res.Error != nil {
		out.Code = e.ERROR
		return
	}
	return
}
