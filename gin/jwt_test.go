package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"testing"
	"time"
)

var jwtSecret = []byte("liulengbo")

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
			token, err := genRegisteredClaims()
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
	r.GET("/protected", authMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "这是受保护的资源"})
	})

	r.Run(":8080")
}

// GenRegisteredClaims 使用默认声明创建jwt
func genRegisteredClaims() (string, error) {
	jwtClaims := &jwt.RegisteredClaims{
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ID:       "liulengbo",
	}
	JwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return JwtToken.SignedString(jwtSecret)
}

func VaildateRegisteredClaims(token string) bool {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return false
	}
	return parse.Valid
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

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
