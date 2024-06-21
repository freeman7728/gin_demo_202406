package http

import (
	"database_lesson/app/gateway/rpc"
	"database_lesson/idl/pb"
	"database_lesson/pkg/ctl"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProductInsertHandler(ctx *gin.Context) {
	var req pb.ProductInsertRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Product Bind 绑定参数失败"))
	}
	resp, err := rpc.ProductInsert(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "ProductInsert RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func ProductDeleteHandler(ctx *gin.Context) {
	var req pb.ProductDeleteRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Product Bind 绑定参数失败"))
	}
	resp, err := rpc.ProductDelete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "ProductDelete RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func ProductUpdateHandler(ctx *gin.Context) {
	var req pb.ProductUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Product Bind 绑定参数失败"))
	}
	fmt.Println()
	resp, err := rpc.ProductUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "ProductUpdate RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func ProductSelectHandler(ctx *gin.Context) {
	var req pb.ProductSelectRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Product Bind 绑定参数失败"))
	}
	resp, err := rpc.ProductSelect(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "ProductSelect RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
