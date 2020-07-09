package common

import (
	"ginessential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/*
jwt的原理和session有点相像，其目的是为了解决restful api中无状态性
因为restful接口需要权限校验。但是又不能每个请求都把用户名密码传入，因此常用token的方法
 */

var jwtKey = []byte("a_secret_credential")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 从tokenString中解析出Claims然后返回
func ReleaseToken(user model.User)(string, error){
	// 定义token有效期  7 * 24h = 7 days , 根据自己需求调整
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),  // token发放时间
			Issuer: "fullstackpentest.com",
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)  // use jwtKey generate token

	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func ParseToken(tokenString string)(*jwt.Token, *Claims, error) {
	claims := &Claims{}
	
	token,err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}