package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// 多格式渲染
	// 既然请求可以使用不同的content-type，响应也如此。通常响应会有html，text，plain，json和xml等
	// gin提供了很优雅的渲染方法。到目前为止，我们已经见识了c.String， c.JSON，c.HTML，下面介绍一下c.XML
	engine.GET("/render", func(c *gin.Context) {
		contentType := c.DefaultQuery("content_type", "json")
		if contentType == "json" {
			c.JSON(http.StatusOK, gin.H{
				"user":   "rsj217",
				"passwd": "123",
			})
		} else if contentType == "xml" {
			c.XML(http.StatusOK, gin.H{
				"user":   "rsj217",
				"passwd": "123",
			})
		}
	})

	engine.Run()
}
