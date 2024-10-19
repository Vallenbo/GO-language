package log_use

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_ginForZap(t *testing.T) {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello ")
	})
	r.Run()
}
