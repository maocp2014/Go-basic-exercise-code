package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerFunc defines the handler used by gin middleware as return value.
// type HandlerFunc func(*Context)
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "client_request")
		c.Next()
		fmt.Println("after middleware")
	}
}

func main() {
	engine := gin.Default()

	// 单个路由中间件
	engine.GET("/before", MiddleWare(), func(c *gin.Context) {
		request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"middile_request": request,
		})
	})

	v1 := engine.Group("/v1", MiddleWare())
	// 或者这样用：
	// v1 := engine.Group("/v1")
	// v1.Use(Middelware())
	{
		v1.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "group middleware")
		})
	}

	// 全局中间件
	engine.Use(MiddleWare())
	{
		engine.GET("/middleware", func(c *gin.Context) {
			request := c.MustGet("request").(string)
			req, _ := c.Get("request")
			c.JSON(http.StatusOK, gin.H{
				"middle_request": request,
				"request":        req,
			})
		})
	}
	engine.Run()
}
