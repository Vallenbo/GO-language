package gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"testing"
	"time"
)

var CustomSecret = []byte("liulengbo")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestJwt(t *testing.T) {
	r := gin.Default()

	// 登录路由，生成 token
	r.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 这里可以进行用户认证逻辑，为了示例简单，假设用户名和密码都是"admin"
		if user.Username == "admin" && user.Password == "admin" {
			token, err := GenToken(user.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 token 失败"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "认证失败"})
		}
	})

	// 需要认证的路由
	api := r.Group("/api", authMiddleware())
	api.GET("/protected", func(c *gin.Context) {
		c.BindJSON()
		c.JSON(http.StatusOK, gin.H{"message": "欢迎，" + c.MustGet("username").(string)})
	})

	r.Run(":8080")
}

// 验证 JWT token 的中间件
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供 token"})
			c.Abort()
			return
		}

		CustomClaim, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
			c.Abort()
			return
		}
		if CustomClaim == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无法解析 token"})
			c.Abort()
			return
		}
		//if !token.Valid().Error() {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
		//	c.Abort()
		//	return
		//}
		c.Set("username", CustomClaim.Username)
		c.Next()
	}
}

type CustomClaims struct { // 自定义声明类型 并内嵌jwt.RegisteredClaims
	Username             string `json:"username"` // 可根据需要自行添加字段
	jwt.RegisteredClaims        // 内嵌标准的声明
}

func GenToken(username string) (string, error) {
	claims := CustomClaims{ // 创建一个我们自己的声明Claims
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "liulengbo", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用指定的签名方法创建签名对象
	return token.SignedString(CustomSecret)                    // 使用指定的secret签名并获得完整的编码后的字符串token
}

func ParseToken(tokenString string) (*CustomClaims, error) { // ParseToken 解析JWT
	// 解析token
	// 直接使用标准的Claim则可以直接使用Parse方法
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) { return CustomSecret, nil })
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
