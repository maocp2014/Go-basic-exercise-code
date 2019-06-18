package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// 注册模板文件路径
	engine.LoadHTMLGlob("src/go_pratice_code/gin/form/templates/*")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	// 注册单个文件上传处理器
	engine.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)
		// gin将net/http包的FormFile函数封装到c.Request
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		// 获取文件名
		filename := header.Filename
		fmt.Println(file, err, filename)
		// 写入文件
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusCreated, "upload successful")
	})

	// 注册多个个文件上传处理器
	engine.POST("/multi/upload", func(c *gin.Context) {
		// 设定最大可用的内存
		err := c.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
		}

		formdata := c.Request.MultipartForm
		files := formdata.File["upload"]
		fmt.Println("filename:", files) // filename: [0xc00021c050 0xc00021c0a0]
		// 循环接收文件
		for i, _ := range files {
			file, err := files[i].Open()
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			out, err := os.Create(files[i].Filename)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
			}

			c.String(http.StatusCreated, "upload successful")
		}
	})

	engine.Run()
}
