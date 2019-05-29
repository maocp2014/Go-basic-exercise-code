package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin的路由来自httprouter库，因此httprouter具有的功能，gin也具有，不过gin不支持路由正则表达式

// 1、:name(任意字符串)组成路由参数
/*
func main() {
	engine := gin.Default()
	// :name构成路由参数(name可以为任意字符串)，使用c.Param的方法读取其值
	// name是字符串string，诸如/user/rsj217，和/user/hello都可以匹配，而/user/和/user/rsj217/不会被匹配
	engine.GET("/user/:name", func(c *gin.Context) {
		// 获取name参数
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})

	engine.Run() // 默认监听8080端口
}
*/

// 2、*组成路由参数
func main() {
	engine := gin.Default()
	// *能匹配/字符串等参数
	engine.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	engine.Run()
}
