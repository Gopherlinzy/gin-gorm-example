package util

import (
	"fmt"
	"time"

	"github.com/Gopherlinzy/go-gin-example/pkg/setting"
	"github.com/golang-jwt/jwt"
)

//签名密钥
var jwtSecret = []byte(setting.JwtSecret)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	c := MyClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			//什么时间生效
			NotBefore: time.Now().Unix(),
			//什么时间失效 3小时生效
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
			//签发者
			Issuer: "gin-blog",
		},
	}
	//NewWithClaims使用指定的签名方法和声明创建一个新的令牌。
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Println(t)
	// SignedString创建并返回一个完整的有符号的JWT。
	//使用令牌中指定的SigningMethod对令牌进行签名。
	token, err := t.SignedString(jwtSecret)
	return token, err
}

func ParseToken(tokenString string) (*MyClaims, error) {
	//解析JWT
	tokenCliams, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenCliams != nil {
		//使用断言获取MyClaims结构体
		if cliams, ok := tokenCliams.Claims.(*MyClaims); ok && tokenCliams.Valid {
			return cliams, nil
		}
	}
	return nil, err
}
