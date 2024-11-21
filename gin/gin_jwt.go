package gin

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var mySigningKey = []byte("liwenzhou.com") // 用于签名的字符串

// GenRegisteredClaims 使用默认声明创建jwt
func GenRegisteredClaims() (string, error) {
	claims := &jwt.RegisteredClaims{ // 创建 Claims
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间
		Issuer:    "qimi",                                             // 签发人
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 生成token对象

	return token.SignedString(mySigningKey) // 生成签名字符串
}

// ParseRegisteredClaims 解析jwt
func ValidateRegisteredClaims(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { // 解析token
		return mySigningKey, nil
	})
	if err != nil { // 解析token失败
		return false
	}
	return token.Valid
}
func main() {

}
