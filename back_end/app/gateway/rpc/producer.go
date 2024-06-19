package rpc

import (
	"context"
	"database_lesson/idl/pb"
)

func ProducerInsert(ctx context.Context, req *pb.InsertProducerRequest) (resp *pb.InsertProducerResponse, err error) {
	resp, err = ProduceService.InsertProducer(ctx, req)
	if err != nil {
		return
	}
	return
}
func ProducerDelete(ctx context.Context, req *pb.DeleteProducerRequest) (resp *pb.DeleteProducerResponse, err error) {
	resp, err = ProduceService.DeleteProducer(ctx, req)
	if err != nil {
		return
	}
	return
}
func ProducerUpdate(ctx context.Context, req *pb.UpdateProducerRequest) (resp *pb.UpdateProducerResponse, err error) {
	resp, err = ProduceService.UpdateProducer(ctx, req)
	if err != nil {
		return
	}
	return
}
func ProducerSelect(ctx context.Context, req *pb.SelectProducerRequest) (resp *pb.SelectProducerResponse, err error) {
	resp, err = ProduceService.SelectProducer(ctx, req)
	if err != nil {
		return
	}
	return
}
