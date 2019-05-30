package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)
	})

	engine.GET("/async", func(c *gin.Context) {
		// 复制上下文到协程
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})

	engine.Run()
}
