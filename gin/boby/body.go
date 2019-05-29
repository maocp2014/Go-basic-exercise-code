package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1、单独的请求体body
/*
func main() {
	engine := gin.Default()
	// post
	engine.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		// gin.H封装了生成json的方式，是一个强大的工具。使用golang可以像动态语言一样写字面量的json，对于嵌套json的实现，嵌套gin.H即可。
		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick":    nick,
		})
	})

	engine.Run()
}
*/

// 2、querystring 与 请求体body共同出现

func main() {
	engine := gin.Default()

	engine.PUT("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
		})
	})

	engine.Run()
}

// curl -X PUT "http://127.0.0.1:8080/post?id=1&page=2" -H "Content-Type:application/x-www-form-urlencoded" -d "name=123 &message=hello"

//id: 1; page: 2; name: 123; message: hello
//{"status_code":200}
