package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 中间件1
func middleware1(c *gin.Context) {
	c.String(http.StatusOK, "exec middleware1\n")

	//你可以写一些逻辑代码

	// 执行该中间件之后的逻辑
	c.Next()

	c.String(http.StatusOK, "quit middleware1\n")
}

// 中间件2
func middleware2(c *gin.Context) {
	c.String(http.StatusOK, "arrive at middleware2\n")
	// 执行该中间件之前，先跳到流程的下一个方法
	c.Next()
	// 流程中的其他逻辑已经执行完了
	c.String(http.StatusOK, "exec middleware2\n")

	//你可以写一些逻辑代码
}

// 路由处理函数
func handler(c *gin.Context) {
	c.String(http.StatusOK, "exec handler\n")
}

func main() {
	engine := gin.Default()

	// 注册一个路由，使用了 middleware1，middleware2 两个中间件
	engine.GET("/mw", middleware1, middleware2, handler)

	// 默认绑定 :8080
	engine.Run()
}

/*
在 middleware2中，执行到 c.Next()时，Gin 会直接跳到流程的下一个方法中，等到这个方法执行完后，
才会回来接着执行 middleware2 剩下的代码。

所以请求上面注册的路由url /mw ，请求先到达middleware1，然后到达 middleware2，
但此时 middleware2调用了 c.Next()，所以 middleware2的代码并没有执行，而是跳到了 handler ，
等 handler执行完成后，跳回到 middleware2，执行 middleware2剩下的代码, 然后返回到middleware1 Next()剩下后的代码
*/

// 输出结果(执行顺序)
/*
curl  http://127.0.0.1:8080/mw
exec middleware1
arrive at middleware2
exec handler
exec middleware2
quit middleware1
*/
