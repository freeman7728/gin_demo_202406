package http

import (
	"database_lesson/app/gateway/rpc"
	"database_lesson/idl/pb"
	"database_lesson/pkg/ctl"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EmployeeSignupHandler(ctx *gin.Context) {
	var req pb.EmployeeRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Employee Bind 绑定参数失败"))
	}
	for _, model := range req.EmployeeList {
		fmt.Println(model.Name)
	}
	resp, err := rpc.EmployeeSignup(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "EmployeeSignup RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
}
