package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
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
*/

func main() {
	// 初始化引擎
	engine := gin.Default()

	// Any registers a route that matches all the HTTP methods.
	// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
	// 注册一个路由和处理函数
	engine.Any("/", WebRoot)
	// 绑定端口，然后启动应用
	engine.Run()
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 c 中
* 输出响应 hello, world
 */
func WebRoot(c *gin.Context) {
	c.String(http.StatusOK, "hello, world")
}
