package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义路由
// 1、方式1
/*
func main() {
	engine := gin.Default()
	http.ListenAndServe(":8080", engine)
}
*/

// 方式2
func main() {
	engine := gin.Default()

	s := &http.Server{
		Addr:           ":8000",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}