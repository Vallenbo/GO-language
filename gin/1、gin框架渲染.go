package gin

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() { //1、html模板函数
	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLGlob("templates/**")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("templates/posts.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "templates/posts.html", gin.H{ // 返回给前端的数据gin.H (map类型)
			"title": "templates/index",
		})
	})

	r.GET("templates/users.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "templates/users.html",
		})
	})

	r.GET("/GetProtoBuf", func(c *gin.Context) { // protobuf
		//protobufData := &hello_grpc.Test{Msg: "this is the gin for returning data"}
		c.ProtoBuf(http.StatusOK, "this is the gin for returning data")
	})

	r.GET("/GetPureJSON", func(c *gin.Context) { // PureJSON渲染
		c.PureJSON(200, gin.H{
			"html": "<b> Hello, world!</b>",
		})
	})

	println("Gin frameWork running localhost:8080")
	r.Run(":8080")
}

func safeTamplate() { // 2、自定义模板函数
	r := gin.Default()
	r.Static("xxx", "./statics") // 指定静态文件处理路径
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "templates/index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})

	r.Run(":8080")
}

func jsonTamplate() { // 3、json格式数据
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) { // 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"}) // gin.H 是map[string]interface{}的缩写
	})
	r.GET("/moreJSON", func(c *gin.Context) { // 方法二：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	r.Run(":8080")
}

func XMLTamplate() { //4、XML格式数据
	r := gin.Default()

	r.GET("/someXML", func(c *gin.Context) { // gin.H 是map[string]interface{}的缩写
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"}) // 方式一：自己拼接JSON
	})
	r.GET("/moreXML", func(c *gin.Context) { // 方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)

	})
	r.Run(":8080")
}
