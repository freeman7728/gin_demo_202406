package http

import (
	"database_lesson/app/gateway/rpc"
	"database_lesson/idl/pb"
	"database_lesson/pkg/ctl"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProducerInsertHandler(ctx *gin.Context) {
	var req pb.InsertProducerRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Producer Bind 绑定参数失败"))
	}
	resp, err := rpc.ProducerInsert(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "ProducerInsert RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func ProducerDeleteHandler(ctx *gin.Context) {
	var req pb.DeleteProducerRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Producer Bind 绑定参数失败"))
	}
	resp, err := rpc.ProducerDelete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "ProducerDelete RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func ProducerUpdateHandler(ctx *gin.Context) {
	fmt.Println(ctx)
	var req pb.UpdateProducerRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Producer Bind 绑定参数失败"))
	}
	resp, err := rpc.ProducerUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "ProducerUpdate RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func ProducerSelectHandler(ctx *gin.Context) {
	var req pb.SelectProducerRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Producer Bind 绑定参数失败"))
	}
	resp, err := rpc.ProducerSelect(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "ProducerSelect RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
