package http

import (
	"database_lesson/app/gateway/rpc"
	"database_lesson/idl/pb"
	"database_lesson/pkg/ctl"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DetailInsertHandler(ctx *gin.Context) {
	var req pb.DetailInsertRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Detail Bind 绑定参数失败"))
	}
	resp, err := rpc.DetailInsert(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "DetailInsert RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func DetailDeleteHandler(ctx *gin.Context) {
	var req pb.DetailDeleteRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Detail Bind 绑定参数失败"))
	}
	resp, err := rpc.DetailDelete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "DetailDelete RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func DetailUpdateHandler(ctx *gin.Context) {
	var req pb.DetailUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Detail Bind 绑定参数失败"))
	}
	fmt.Println()
	resp, err := rpc.DetailUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "DetailUpdate RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func DetailSelectHandler(ctx *gin.Context) {
	var req pb.DetailSelectRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Detail Bind 绑定参数失败"))
	}
	resp, err := rpc.DetailSelect(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "DetailSelect RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
