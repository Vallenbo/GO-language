package gin

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	//r.Use(indexHandler) //全局可以使用中间件
	r.GET("/", StatCost(), indexHandler) //中间按先后顺序执行

}

func StatCost() gin.HandlerFunc { // 1、自定义中间件方式一 ， StatCost 是一个统计耗时请求耗时的中间件
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		c.Next()             // 直接调用该请求的剩余处理程序，如indexHandler中间件，执行完该中间件再向下执行
		// c.Abort() // 不调用该请求的剩余处理程序
		cost := time.Since(start) // 计算耗时
		log.Println(cost)
	}
}

func indexHandler(c *gin.Context) { // 2、自定义中间件方式二，
	log.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

func Gobal() { //3、全局路由注册
	r := gin.New()    // 新建一个没有任何默认中间件的路由
	r.Use(StatCost()) // 注册一个全局中间件

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run()
}
