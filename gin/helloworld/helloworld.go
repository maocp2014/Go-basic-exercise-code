package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建包含Recovery和Logger中间件的engine实例
	engine := gin.Default()
	// 绑定路由和路由处理函数
	// gin把request和response都封装到gin.Context的上下文环境
	engine.GET("/", func(c *gin.Context) {
		// String方法将相应的响应状态码和字符串写入response
		c.String(http.StatusOK, "hello world!")
	})
	// 启动engine
	engine.Run(":8080")
}
