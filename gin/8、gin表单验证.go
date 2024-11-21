package gin

import (
	"GinNote/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var login model.Login
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	router.POST("/SignUpJSON", func(c *gin.Context) {
		var SignUpJSON model.SignUp
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&SignUpJSON); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"Msg": "Sign Up successes",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	println("Gin frameWork running localhost:80")
	router.Run()
}
