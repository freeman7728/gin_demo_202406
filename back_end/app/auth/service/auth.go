package service

import (
	"context"
	"database_lesson/app/auth/repository/db/dao"
	"database_lesson/app/auth/repository/db/model"
	"database_lesson/idl/pb"
	"database_lesson/pkg/e"
	"sync"
)

var AuthSrvIns *AuthService
var AuthSrvInsOnce sync.Once

type AuthService struct {
}

// GetAuthSrv 懒汉单例模式
func GetAuthSrv() *AuthService {
	AuthSrvInsOnce.Do(func() {
		AuthSrvIns = &AuthService{}
	})
	return AuthSrvIns
}

// RegisterToken
// /*
// 先查level -> 注入到model
// 再生成token -> 注入到model
// 删除id对应token
// model注入数据库
// */
func (receiver *AuthService) RegisterToken(ctx context.Context, in *pb.TokenRegisterReq, out *pb.TokenRegisterRsp) (err error) {
	out.Code = e.SUCCESS
	authIns := model.BindAuthRegister(in)
	daoIns := dao.NewAuthDao(ctx)
	err = daoIns.GetLevel(authIns).Error
	//通过id查询不到这个人，返回
	if err != nil {
		out.Code = e.ERROR
		return
	}
	err = model.GenToken(authIns)
	if err != nil {
		out.Code = e.ERROR
		return
	}
	err = daoIns.RegisterToken(authIns).Error
	if err != nil {
		out.Code = e.ERROR
		return
	}
	out.Token = authIns.Token
	return
}

// DiscoverToken /*
// 入参id
// */
func (receiver *AuthService) DiscoverToken(ctx context.Context, in *pb.TokenDiscoveryReq, out *pb.TokenDiscoveryRsp) (err error) {
	out.Code = e.SUCCESS
	authIns := model.BindAuthDiscover(in)
	daoIns := dao.NewAuthDao(ctx)
	var oldToken string
	var token string
	err = daoIns.GetToken(authIns, &oldToken).Error
	if err != nil {
		out.Code = e.ERROR
		return
	}
	if oldToken == "" || !model.TokenIsVa(&oldToken) {
		err = daoIns.RegisterToken(authIns).Error
		if err != nil {
			out.Code = e.ERROR
			return
		}
	}
	err = daoIns.GetToken(authIns, &token).Error
	if err != nil {
		out.Code = e.ERROR
		return
	}
	out.Token = token
	return
}
func (receiver *AuthService) AuthorizeToken(ctx context.Context, in *pb.TokenAuthorizationReq, out *pb.TokenAuthorizationRsp) (err error) {
	out.Code = e.SUCCESS
	daoIns := dao.NewAuthDao(ctx)
	authIns := model.BindAuthAuthorize(in)
	var level int
	err = daoIns.GetLevelByToken(authIns, &level).Error
	if err != nil {
		out.Code = e.ERROR
		return
	}
	if level == 0 {
		out.Code = e.ERROR
		return
	}
	out.Level = int32(level)
	return
}
