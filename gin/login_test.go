package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_login(t *testing.T) {
	engine := gin.Default()
	engine.LoadHTMLFiles("html/login.html")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	engine.POST("/login", func(c *gin.Context) {
		nameMap := c.PostFormMap("name")
		passwordMap := c.PostFormMap("password")

		c.JSON(http.StatusOK, gin.H{
			"name":     nameMap,
			"password": passwordMap,
		})
	})
	engine.GET("/logout", func(c *gin.Context) {
		c.String(http.StatusOK, "退出登录")
	})
	engine.Run()
}
