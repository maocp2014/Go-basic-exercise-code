package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 上传单个文件
/*
func main() {
	engine := gin.Default()

	engine.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)

		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}

		filename := header.Filename
		fmt.Println(file, err, filename)

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

	engine.Run(":8080")
}
*/

// curl -X POST http://127.0.0.1:8080/upload -F "upload=@F:\\桌面\\pic.jpg" -H "Content-Type: multipart/form-data"

// 2、上传多个文件
func main() {
	engine := gin.Default()
	engine.POST("/multi/upload", func(c *gin.Context) {
		// 设定最大可用的内存
		err := c.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
		}

		formdata := c.Request.MultipartForm
		files := formdata.File["upload"]
		fmt.Println("filename:", files)  // filename: [0xc00021c050 0xc00021c0a0]
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
