package http

import (
	"database_lesson/app/gateway/rpc"
	"database_lesson/idl/pb"
	"database_lesson/pkg/ctl"
	"database_lesson/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EmployeeSignupHandler(ctx *gin.Context) {
	var req pb.EmployeeRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Employee Bind 绑定参数失败"))
	}
	//for _, model := range req.EmployeeList {
	//	fmt.Println(model.Name)
	//}
	resp, err := rpc.EmployeeSignup(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "EmployeeSignup RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, ""))
}
func EmployeeLoginHandler(ctx *gin.Context) {
	var req pb.EmployeeLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "Login Bind 绑定参数失败"))
	}
	resp, err := rpc.EmployeeLogin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "EmployeeLogin RPC 调用失败"))
		return
	}
	if resp.Code == e.WrongPassword {
		ctx.JSON(http.StatusOK, ctl.RespWrong(ctx, e.GetMsg(int(resp.Code))))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp, "登录成功"))
}
