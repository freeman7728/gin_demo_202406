package model

import (
	"database_lesson/idl/pb"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义一个jwt加密的密钥
var jwtKey = []byte("a_secret_crect")

// Claims 定义token的Claims
type Claims struct {
	UserId int
	jwt.StandardClaims
}
type Auth struct {
	Id    int
	Token string
	Level int
}

func BindAuthRegister(req *pb.TokenRegisterReq) *Auth {
	auth := &Auth{
		Id: int(req.Id),
	}
	return auth
}

func BindAuthDiscover(req *pb.TokenDiscoveryReq) *Auth {
	auth := &Auth{
		Id: int(req.Id),
	}
	return auth
}

func BindAuthAuthorize(req *pb.TokenAuthorizationReq) *Auth {
	auth := &Auth{
		Token: req.Token,
	}
	return auth
}

func GenToken(auth *Auth) error {
	//定义token的过期时间:7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: auth.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), //token发放的时间
			Issuer:    "admin",           //发放人
			Subject:   "user token",      //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用jwt密钥来生成token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return err
	}
	auth.Token = tokenString
	return nil
}
func TokenIsVa(token *string) bool {
	return true
}
