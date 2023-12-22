package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Any("/test", func(c *gin.Context) { c.Value("receive!") }) //1、普通路由
	r.NoRoute(func(c *gin.Context) {                             // 1-1、为无匹配路由请求指定响应
		c.HTML(http.StatusNotFound, "statics/404.html", nil)
	})

	userGroup := r.Group("/user") //2、设置路由组
	{
		userGroup.GET("/index", func(c *gin.Context) { c.Value("hello") })
		userGroup.GET("/login", func(c *gin.Context) { c.Value("hello") })
		userGroup.POST("/login", func(c *gin.Context) { c.Value("hello") })
	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) { c.Value("hello") })
		shopGroup.GET("/cart", func(c *gin.Context) { c.Value("hello") })
		shopGroup.POST("/checkout", func(c *gin.Context) { c.Value("hello") })
	}

	shopGroup1 := r.Group("/shop") //2-2、设置嵌套路由组
	{
		shopGroup1.GET("/index", func(c *gin.Context) { c.Value("hello") })
		shopGroup1.GET("/cart", func(c *gin.Context) { c.Value("hello") })
		shopGroup1.POST("/checkout", func(c *gin.Context) { c.Value("hello") })
		xx := shopGroup1.Group("xx") // 嵌套路由组
		xx.GET("/oo", func(c *gin.Context) { c.Value("hello") })
	}
	r.Run()
}
