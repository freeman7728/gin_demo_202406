package service

import (
	"context"
	"database_lesson/app/producer/repository/db/dao"
	"database_lesson/idl/pb"
	"database_lesson/pkg/e"
	"errors"
	"fmt"
	"sync"
)

var ProducerServIns *ProducerServ
var ProducerServOnce sync.Once

type ProducerServ struct{}

func GetProducerServ() *ProducerServ {
	ProducerServOnce.Do(func() {
		ProducerServIns = &ProducerServ{}
	})
	return ProducerServIns
}

func (p *ProducerServ) InsertProducer(ctx context.Context, in *pb.InsertProducerRequest, out *pb.InsertProducerResponse) (err error) {
	daoIns := dao.NewProducerDao(ctx)
	out.Code = e.SUCCESS
	for _, model := range in.ProducerList {
		res := daoIns.InsertProducer(model)
		if res.Error == nil {
			out.SuccessList = append(out.SuccessList, model)
		} else {
			out.Code = e.ERROR
			err = errors.New("插入数据出现错误")
			out.FailList = append(out.FailList, model)
		}
	}
	return
}

func (p *ProducerServ) UpdateProducer(ctx context.Context, in *pb.UpdateProducerRequest, out *pb.UpdateProducerResponse) (err error) {
	daoIns := dao.NewProducerDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.UpdateProducer(in.Producer)
	if res.Error != nil {
		err = res.Error
		out.Code = e.ERROR
		out.Success = false
		return
	}
	if res.RowsAffected == 0 {
		out.Success = false
		return
	}
	out.Success = true
	return
}

func (p *ProducerServ) SelectProducer(ctx context.Context, in *pb.SelectProducerRequest, out *pb.SelectProducerResponse) (err error) {
	daoIns := dao.NewProducerDao(ctx)
	out.Code = e.SUCCESS
	var list []*pb.ProducerModel
	err = daoIns.SelectProducer(&list).Error
	fmt.Println(list)
	if err != nil {
		out.Code = e.ERROR
		return
	}
	out.ProducerList = list
	return
}

func (p *ProducerServ) DeleteProducer(ctx context.Context, in *pb.DeleteProducerRequest, out *pb.DeleteProducerResponse) (err error) {
	daoIns := dao.NewProducerDao(ctx)
	out.Code = e.SUCCESS
	res := daoIns.DeleteProducer(in.Producer)
	if res.Error != nil {
		err = res.Error
		out.Code = e.ERROR
		out.Success = false
	}
	if res.RowsAffected == 0 {
		out.Success = false
		return
	}
	out.Success = true
	return
}
