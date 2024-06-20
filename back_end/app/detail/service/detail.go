package service

import (
	"context"
	"database_lesson/app/detail/repository/db/dao"
	"database_lesson/idl/pb"
	"database_lesson/pkg/e"
	"sync"
)

var DetailServIns *DetailServ
var DetailServOnce sync.Once

type DetailServ struct{}

func GetDetailServ() *DetailServ {
	DetailServOnce.Do(func() {
		DetailServIns = &DetailServ{}
	})
	return DetailServIns
}

func (d *DetailServ) InsertDetail(ctx context.Context, in *pb.DetailInsertRequest, out *pb.DetailInsertResponse) (err error) {
	daoIns := dao.NewDetailDao(ctx)
	out.Code = e.SUCCESS
	for _, model := range in.List {
		res := daoIns.InsertDetail(model)
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

func (d *DetailServ) DeleteDetail(ctx context.Context, in *pb.DetailDeleteRequest, out *pb.DetailDeleteResponse) (err error) {
	daoIns := dao.NewDetailDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.DeleteDetail(in)
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

func (d *DetailServ) UpdateDetail(ctx context.Context, in *pb.DetailUpdateRequest, out *pb.DetailUpdateResponse) (err error) {
	daoIns := dao.NewDetailDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.UpdateDetail(in)
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

func (d *DetailServ) SelectDetail(ctx context.Context, in *pb.DetailSelectRequest, out *pb.DetailSelectResponse) (err error) {
	daoIns := dao.NewDetailDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.SelectDetail(out)
	if res.Error != nil {
		out.Code = e.ERROR
		return
	}
	return
}
