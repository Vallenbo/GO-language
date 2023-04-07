package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.LoadHTMLFiles("statics/fileUpload.html")
	r.GET("/fileupload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "statics/fileUpload.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) { //接收url：/upload 的 post请求
		file, err := c.FormFile("f1") // 单个文件

		if err != nil { //传输有错误
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("./%s", file.Filename)

		c.SaveUploadedFile(file, dst) // 保存上传文件到指定的目录
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})
	r.Run()
}

func FileUpload() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm() // Multipart form
		files := form.File["file"]   //获取文件对象
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%s_%d", file.Filename, index)
			c.SaveUploadedFile(file, dst) // 上传文件到指定的目录
		}
		c.JSON(http.StatusOK, gin.H{ //返回响应正文
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	router.Run()
}
