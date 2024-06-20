// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: detailService.proto

package pb

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for DetailService service

func NewDetailServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DetailService service

type DetailService interface {
	InsertDetail(ctx context.Context, in *DetailInsertRequest, opts ...client.CallOption) (*DetailInsertResponse, error)
	DeleteDetail(ctx context.Context, in *DetailDeleteRequest, opts ...client.CallOption) (*DetailDeleteResponse, error)
	UpdateDetail(ctx context.Context, in *DetailUpdateRequest, opts ...client.CallOption) (*DetailUpdateResponse, error)
	SelectDetail(ctx context.Context, in *DetailSelectRequest, opts ...client.CallOption) (*DetailSelectResponse, error)
}

type detailService struct {
	c    client.Client
	name string
}

func NewDetailService(name string, c client.Client) DetailService {
	return &detailService{
		c:    c,
		name: name,
	}
}

func (c *detailService) InsertDetail(ctx context.Context, in *DetailInsertRequest, opts ...client.CallOption) (*DetailInsertResponse, error) {
	req := c.c.NewRequest(c.name, "DetailService.InsertDetail", in)
	out := new(DetailInsertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detailService) DeleteDetail(ctx context.Context, in *DetailDeleteRequest, opts ...client.CallOption) (*DetailDeleteResponse, error) {
	req := c.c.NewRequest(c.name, "DetailService.DeleteDetail", in)
	out := new(DetailDeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detailService) UpdateDetail(ctx context.Context, in *DetailUpdateRequest, opts ...client.CallOption) (*DetailUpdateResponse, error) {
	req := c.c.NewRequest(c.name, "DetailService.UpdateDetail", in)
	out := new(DetailUpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detailService) SelectDetail(ctx context.Context, in *DetailSelectRequest, opts ...client.CallOption) (*DetailSelectResponse, error) {
	req := c.c.NewRequest(c.name, "DetailService.SelectDetail", in)
	out := new(DetailSelectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DetailService service

type DetailServiceHandler interface {
	InsertDetail(context.Context, *DetailInsertRequest, *DetailInsertResponse) error
	DeleteDetail(context.Context, *DetailDeleteRequest, *DetailDeleteResponse) error
	UpdateDetail(context.Context, *DetailUpdateRequest, *DetailUpdateResponse) error
	SelectDetail(context.Context, *DetailSelectRequest, *DetailSelectResponse) error
}

func RegisterDetailServiceHandler(s server.Server, hdlr DetailServiceHandler, opts ...server.HandlerOption) error {
	type detailService interface {
		InsertDetail(ctx context.Context, in *DetailInsertRequest, out *DetailInsertResponse) error
		DeleteDetail(ctx context.Context, in *DetailDeleteRequest, out *DetailDeleteResponse) error
		UpdateDetail(ctx context.Context, in *DetailUpdateRequest, out *DetailUpdateResponse) error
		SelectDetail(ctx context.Context, in *DetailSelectRequest, out *DetailSelectResponse) error
	}
	type DetailService struct {
		detailService
	}
	h := &detailServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DetailService{h}, opts...))
}

type detailServiceHandler struct {
	DetailServiceHandler
}

func (h *detailServiceHandler) InsertDetail(ctx context.Context, in *DetailInsertRequest, out *DetailInsertResponse) error {
	return h.DetailServiceHandler.InsertDetail(ctx, in, out)
}

func (h *detailServiceHandler) DeleteDetail(ctx context.Context, in *DetailDeleteRequest, out *DetailDeleteResponse) error {
	return h.DetailServiceHandler.DeleteDetail(ctx, in, out)
}

func (h *detailServiceHandler) UpdateDetail(ctx context.Context, in *DetailUpdateRequest, out *DetailUpdateResponse) error {
	return h.DetailServiceHandler.UpdateDetail(ctx, in, out)
}

func (h *detailServiceHandler) SelectDetail(ctx context.Context, in *DetailSelectRequest, out *DetailSelectResponse) error {
	return h.DetailServiceHandler.SelectDetail(ctx, in, out)
}
