package gin

import (
	"GinNote/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() { // 1、获取querystring参数
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) { //http://web?username=小王子&address=沙河
		address := c.Query("address") // 返回 url 中key对应的值（如果存在），否则返回空字符串
		//username ,err := c.GetQuery("username")
		username := c.DefaultQuery("username", "小王子") // 指定返回的默认值字符串
		c.JSON(http.StatusOK, gin.H{                  //输出json结果给调用方
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}

func queryForm() { //2、获取表单form参数
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		//username := c.DefaultPostForm("username", "小王子")// DefaultPostForm取不到值时会返回指定的默认值
		username := c.PostForm("username")
		password := c.PostForm("password")
		//username1, ok := c.GetPostForm("username")
		//if !ok {
		//	username1 = "sb"
		//}
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password": password,
		})
	})
	r.Run(":8080")
}

func queryPath() { //3、获取Path参数
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) { // http://user/search/:xiaobai/:hunan
		username := c.Param("username") //获取对应路径下指定key的参数
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.Run(":8080")
}

func parameterBind() { //4、参数绑定
	router := gin.Default()

	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var login model.Login

		if err := c.ShouldBind(&login); err == nil { //将前端参数和结构体内成员 自动选择绑定
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{" loginJSON bind error": err.Error()})
		}
	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	router.POST("/loginForm", func(c *gin.Context) {
		var login model.Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	router.GET("/loginForm", func(c *gin.Context) {
		var login model.Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
