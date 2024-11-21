package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) { //1、HTTP重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.sogo.com/")
	})

	r.GET("/test1", func(c *gin.Context) { //2、路由重定向
		c.Request.URL.Path = "/test2" // 指定重定向的URL
		r.HandleContext(c)            //继续后续处理
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	r.Run()
}
