// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: productService.proto

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

// Api Endpoints for ProductService service

func NewProductServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProductService service

type ProductService interface {
	InsertProduct(ctx context.Context, in *ProductInsertRequest, opts ...client.CallOption) (*ProductInsertResponse, error)
	DeleteProduct(ctx context.Context, in *ProductDeleteRequest, opts ...client.CallOption) (*ProductDeleteResponse, error)
	UpdateProduct(ctx context.Context, in *ProductUpdateRequest, opts ...client.CallOption) (*ProductUpdateResponse, error)
	SelectProduct(ctx context.Context, in *ProductSelectRequest, opts ...client.CallOption) (*ProductSelectResponse, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) InsertProduct(ctx context.Context, in *ProductInsertRequest, opts ...client.CallOption) (*ProductInsertResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.InsertProduct", in)
	out := new(ProductInsertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) DeleteProduct(ctx context.Context, in *ProductDeleteRequest, opts ...client.CallOption) (*ProductDeleteResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.DeleteProduct", in)
	out := new(ProductDeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) UpdateProduct(ctx context.Context, in *ProductUpdateRequest, opts ...client.CallOption) (*ProductUpdateResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.UpdateProduct", in)
	out := new(ProductUpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) SelectProduct(ctx context.Context, in *ProductSelectRequest, opts ...client.CallOption) (*ProductSelectResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.SelectProduct", in)
	out := new(ProductSelectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceHandler interface {
	InsertProduct(context.Context, *ProductInsertRequest, *ProductInsertResponse) error
	DeleteProduct(context.Context, *ProductDeleteRequest, *ProductDeleteResponse) error
	UpdateProduct(context.Context, *ProductUpdateRequest, *ProductUpdateResponse) error
	SelectProduct(context.Context, *ProductSelectRequest, *ProductSelectResponse) error
}

func RegisterProductServiceHandler(s server.Server, hdlr ProductServiceHandler, opts ...server.HandlerOption) error {
	type productService interface {
		InsertProduct(ctx context.Context, in *ProductInsertRequest, out *ProductInsertResponse) error
		DeleteProduct(ctx context.Context, in *ProductDeleteRequest, out *ProductDeleteResponse) error
		UpdateProduct(ctx context.Context, in *ProductUpdateRequest, out *ProductUpdateResponse) error
		SelectProduct(ctx context.Context, in *ProductSelectRequest, out *ProductSelectResponse) error
	}
	type ProductService struct {
		productService
	}
	h := &productServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductService{h}, opts...))
}

type productServiceHandler struct {
	ProductServiceHandler
}

func (h *productServiceHandler) InsertProduct(ctx context.Context, in *ProductInsertRequest, out *ProductInsertResponse) error {
	return h.ProductServiceHandler.InsertProduct(ctx, in, out)
}

func (h *productServiceHandler) DeleteProduct(ctx context.Context, in *ProductDeleteRequest, out *ProductDeleteResponse) error {
	return h.ProductServiceHandler.DeleteProduct(ctx, in, out)
}

func (h *productServiceHandler) UpdateProduct(ctx context.Context, in *ProductUpdateRequest, out *ProductUpdateResponse) error {
	return h.ProductServiceHandler.UpdateProduct(ctx, in, out)
}

func (h *productServiceHandler) SelectProduct(ctx context.Context, in *ProductSelectRequest, out *ProductSelectResponse) error {
	return h.ProductServiceHandler.SelectProduct(ctx, in, out)
}
