package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 校验jwt

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("一闪一闪亮晶晶")

// MyClaims jwt信息结构体
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成jwt
func GenToken(username string) (string, error) {
	// 创建声明
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "jwt-consul-rpc-server",
		},
	}
	// 创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析jwt
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
