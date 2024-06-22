package middleware

import (
	"database_lesson/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

// 定义token的Claims
type Claims struct {
	UserId int
	jwt.StandardClaims
}

// 登录成功之后就调用这个方法来释放token
func ReleaseToken(employee models.Employee) (string, error) {
	//定义token的过期时间:7天
	//TODO 可以把过期时间抽象到yaml
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: employee.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), //token发放的时间
			Issuer:    "admin",           //发放人
			Subject:   "employee token",  //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用jwt密钥来生成token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
